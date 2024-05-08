package services

import (
	e "application-design-test-master/errors"
	"application-design-test-master/models"
	b "application-design-test-master/repositories/booking"
	"application-design-test-master/utils"
	"time"
)

// orderService struct implements the OrderService interface.
type orderService struct {
	repository b.Reposytory
}

// NewOrderService creates a new instance of OrderService.
func NewOrderService(repository b.Reposytory) OrderService {
	return &orderService{
		repository: repository,
	}
}

// CreateOrder create a new order.
func (os *orderService) CreateOrder(order models.Order) error {
	daysToBook, err := utils.DaysBetween(order.From, order.To)
	if err != nil {
		return err
	}

	unavailableDays := make(map[time.Time]struct{})

	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	roomAvailability := os.repository.GetRoomAvailability(order)

	for _, dayToBook := range daysToBook {
		for i, availability := range roomAvailability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}

			availability.Quota -= 1

			os.repository.UpdateRoomAvailability(availability, i)

			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		return e.ErrHotelRoomInNotAvailable
	}

	return nil
}
