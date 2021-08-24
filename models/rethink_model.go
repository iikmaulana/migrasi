package models

import "time"

type RethinkModel struct {
	Address       string    `json:"address"`
	AppID         string    `json:"app_id"`
	Code          string    `json:"code"`
	CreatedAt     time.Time `json:"created_at"`
	DrivingStatus string    `json:"driving_status"`
	JobStatus     string    `json:"job_status"`
	Name          string    `json:"name"`
	OwnerID       string    `json:"owner_id"`
	Phone         string    `json:"phone"`
	Photo         string    `json:"photo"`
	SimExpired    string    `json:"sim_expired"`
	SimNumber     string    `json:"sim_number"`
	UserID        string    `json:"user_id"`
	VehicleID     string    `json:"vehicle_id"`
}
