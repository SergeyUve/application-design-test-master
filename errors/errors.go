package errors

import "errors"

var (
	ErrHotelRoomInNotAvailable = errors.New("Hotel room is not available for selected dates")
	ErrInvalidDateRange        = errors.New("invalid date range: 'to' date is before 'from' date")
)
