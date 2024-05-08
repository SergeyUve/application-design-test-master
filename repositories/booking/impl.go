package repositories

import (
	"application-design-test-master/models"
	"application-design-test-master/utils"
	"sync"
)

// reposytory struct implements the Reposytory interface.
type reposytory struct {
	roomAvailability []models.RoomAvailability
	mu               sync.Mutex
}

// NewReposytory() create a new instance of Reposytory.
func NewReposytory() Reposytory {
	return &reposytory{
		roomAvailability: []models.RoomAvailability{
			{"reddison", "lux", utils.Date(2024, 1, 1), 1},
			{"reddison", "lux", utils.Date(2024, 1, 2), 1},
			{"reddison", "lux", utils.Date(2024, 1, 3), 1},
			{"reddison", "lux", utils.Date(2024, 1, 4), 1},
			{"reddison", "lux", utils.Date(2024, 1, 5), 0},
		},
	}
}

// GetRoomAvailability() get all rooms with availability.
func (r *reposytory) GetRoomAvailability(order models.Order) []models.RoomAvailability {
	res := make([]models.RoomAvailability, 0)

	r.mu.Lock()
	data := r.roomAvailability
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
func (r *reposytory) UpdateRoomAvailability(roomAvailability models.RoomAvailability, indexOfRoom int) {
	r.mu.Lock()

	r.roomAvailability[indexOfRoom] = roomAvailability

	r.mu.Unlock()

}
