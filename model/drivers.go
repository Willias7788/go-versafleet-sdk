package model

import "time"

type DriverNote struct {
	ID                       *int   `json:"id"`
	CustomFieldDescriptionID *int   `json:"custom_field_description_id"`
	Value                    string `json:"value"`
}

type Skill struct {
	Name string `json:"name"`
}

type Entity struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ContactNumber string `json:"contact_number"`
}

type Person struct {
	ID               int           `json:"id"`
	Guid             string        `json:"guid"`
	Name             string        `json:"name"`
	Dob              string        `json:"dob"` // Usually "YYYY-MM-DD"
	Nric             string        `json:"nric"`
	ContactNumber    string        `json:"contact_number"`
	License          string        `json:"license"`
	Status           string        `json:"status"`
	Archived         bool          `json:"archived"`
	IsVersadriveUser bool          `json:"is_versadrive_user"`
	Username         string        `json:"username"`
	HasPassword      bool          `json:"has_password"`
	IsAttendant      bool          `json:"is_attendant"`
	LastSeen         *LastSeen     `json:"last_seen"`
	DefaultVehicle   *Vehicle      `json:"default_vehicle"`
	Address          *Address      `json:"address"`
	Attendant        *Person       `json:"attendant"` // null in JSON
	Driver           *Person       `json:"driver"`    // null in JSON
	CustomFields     []CustomField `json:"custom_fields"`
	Skills           []string      `json:"skills"`
}

type LastSeen struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Time      time.Time `json:"time"`
}
