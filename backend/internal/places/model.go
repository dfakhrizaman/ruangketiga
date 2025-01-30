package places

import "time"

type Place struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Address     string    `json:"address"`
	District    string    `json:"district"`
	Subdistrict string    `json:"subdistrict"`
	City        string    `json:"city"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	SizeM2      int       `json:"size_m2"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
