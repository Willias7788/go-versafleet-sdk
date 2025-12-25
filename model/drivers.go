package model

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
