package repositories

import "application-design-test-master/models"

// Repository interface defining methods for handling room availability and updates.
type Reposytory interface {
	// GetRoomAvailability returns the room availability.
	GetRoomAvailability(models.Order) []models.RoomAvailability

	// UpdateRoomAvailability updates the room availability information for a specific room.
	UpdateRoomAvailability(roomAvailability models.RoomAvailability, indexOfRoom int)
}
