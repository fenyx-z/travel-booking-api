package dto

type CreateBookingRequest struct {
	UserID     uint `json:"user_id"`
	ScheduleID uint `json:"schedule_id"`
	TotalSeats int  `json:"total_seats"`
}