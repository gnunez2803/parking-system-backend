package reservations

import "time"

type ParkingSpot struct {
	ID       int     `json:"id"`
	Location string  `json:"location"`
	Price    float32 `json:"price"`
	Occupied bool    `json:"occupied"`
}

type Reservation struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	SpotID    int       `json:"spot_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"` // For simplicity, store hashed password
}
