package services

import (
	"application-design-test-master/models"
	b "application-design-test-master/repositories/booking"
	"application-design-test-master/utils"
	"sync"
	"testing"
)

// reposytory struct implements the Reposytory interface.
type mockReposytory struct {
	RoomAvailability []models.RoomAvailability
	mu               sync.Mutex
}

// GetRoomAvailability() get all rooms with availability.
func (r *mockReposytory) GetRoomAvailability(order models.Order) []models.RoomAvailability {
	res := make([]models.RoomAvailability, 0)

	r.mu.Lock()
	data := r.RoomAvailability
	r.mu.Unlock()

	for _, room := range data {
		if room.Quota == 0 || order.HotelID != room.HotelID || order.RoomID != room.RoomID {
			continue
		}

		res = append(res, room)
	}

	return res
}

// UpdateRoomAvailability() update availability on room with index.
func (r *mockReposytory) UpdateRoomAvailability(roomAvailability models.RoomAvailability, indexOfRoom int) {
	r.mu.Lock()

	r.RoomAvailability[indexOfRoom] = roomAvailability

	r.mu.Unlock()

}

func Test_orderService_CreateOrder(t *testing.T) {
	testsTable := []struct {
		name       string
		repository b.Reposytory
		order      models.Order
		wantErr    string
	}{
		{
			name: "case-1. successfull order",
			repository: &mockReposytory{
				RoomAvailability: []models.RoomAvailability{
					{"reddison", "lux", utils.Date(2024, 1, 1), 1},
					{"reddison", "lux", utils.Date(2024, 1, 2), 1},
					{"reddison", "lux", utils.Date(2024, 1, 3), 1},
					{"reddison", "lux", utils.Date(2024, 1, 4), 1},
					{"reddison", "lux", utils.Date(2024, 1, 5), 0}},
			},
			order: models.Order{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "test@gmail.com",
				From:      utils.Date(2024, 1, 1),
				To:        utils.Date(2024, 1, 4),
			},
			wantErr: "",
		},
		{
			name: "case-2. Hotel room is not available for selected dates",
			repository: &mockReposytory{
				RoomAvailability: []models.RoomAvailability{
					{"reddison", "lux", utils.Date(2024, 1, 1), 0},
					{"reddison", "lux", utils.Date(2024, 1, 2), 0},
					{"reddison", "lux", utils.Date(2024, 1, 3), 0},
					{"reddison", "lux", utils.Date(2024, 1, 4), 0},
					{"reddison", "lux", utils.Date(2024, 1, 5), 0}},
			},
			order: models.Order{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "test@gmail.com",
				From:      utils.Date(2024, 1, 1),
				To:        utils.Date(2024, 1, 4),
			},
			wantErr: "Hotel room is not available for selected dates",
		},
		{
			name: "case-3. invalid date range: 'to' date is before 'from' date",
			repository: &mockReposytory{
				RoomAvailability: []models.RoomAvailability{
					{"reddison", "lux", utils.Date(2024, 1, 1), 1},
					{"reddison", "lux", utils.Date(2024, 1, 2), 1},
					{"reddison", "lux", utils.Date(2024, 1, 3), 1},
					{"reddison", "lux", utils.Date(2024, 1, 4), 1},
					{"reddison", "lux", utils.Date(2024, 1, 5), 1}},
			},
			order: models.Order{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "test@gmail.com",
				From:      utils.Date(2024, 1, 2),
				To:        utils.Date(2024, 1, 1),
			},
			wantErr: "invalid date range: 'to' date is before 'from' date",
		},
		{
			name: "case-4. Hotel room is not available for selected dates",
			repository: &mockReposytory{
				RoomAvailability: []models.RoomAvailability{
					{"reddison", "lux", utils.Date(2024, 1, 1), 1},
					{"reddison", "lux", utils.Date(2024, 1, 2), 1},
					{"reddison", "lux", utils.Date(2024, 1, 3), 1},
					{"reddison", "lux", utils.Date(2024, 1, 4), 1},
					{"reddison", "lux", utils.Date(2024, 1, 5), 1}},
			},
			order: models.Order{
				HotelID:   "test",
				RoomID:    "lux",
				UserEmail: "test@gmail.com",
				From:      utils.Date(2024, 1, 1),
				To:        utils.Date(2024, 1, 4),
			},
			wantErr: "Hotel room is not available for selected dates",
		},
		{
			name: "case-5. Hotel room is not available for selected dates",
			repository: &mockReposytory{
				RoomAvailability: []models.RoomAvailability{
					{"reddison", "lux", utils.Date(2024, 1, 1), 1},
					{"reddison", "lux", utils.Date(2024, 1, 2), 1},
					{"reddison", "lux", utils.Date(2024, 1, 3), 1},
					{"reddison", "lux", utils.Date(2024, 1, 4), 1},
					{"reddison", "lux", utils.Date(2024, 1, 5), 1}},
			},
			order: models.Order{
				HotelID:   "reddison",
				RoomID:    "test",
				UserEmail: "test@gmail.com",
				From:      utils.Date(2024, 1, 1),
				To:        utils.Date(2024, 1, 4),
			},
			wantErr: "Hotel room is not available for selected dates",
		},
	}
	for _, tt := range testsTable {
		t.Run(tt.name, func(t *testing.T) {
			os := &orderService{
				repository: tt.repository,
			}

			err := os.CreateOrder(tt.order)
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("%v \n have error : %v, \n want error: %v", tt.name, err, tt.wantErr)
				}
			}
		})
	}
}
