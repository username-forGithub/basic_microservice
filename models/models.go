package models

import "time"

type Order struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	Address       string    `json:"address,omitempty"`
	Meta          string    `json:"meta,omitempty"`
	Status        string    `json:"status,omitempty"`
	Deliviry_time time.Time `json:"deliviry_time"`

	Coordinate *Coordinate `json:"coordinate"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}
