package model

type Vehicle struct {
	ID                    int           `json:"id,omitempty"`
	Guid                  string        `json:"guid,omitempty"`
	PlateNumber           string        `json:"plate_number"`
	Status                string        `json:"status,omitempty"`
	CargoLoad             float64       `json:"cargo_load,omitempty"`
	Model                 string        `json:"model,omitempty"`
	Category              string        `json:"category,omitempty"`
	OwnershipDate         string        `json:"ownership_date,omitempty"`
	RegistrationDate      string        `json:"registration_date,omitempty"`
	InsuranceExpiry       string        `json:"insurance_expiry,omitempty"`
	TaxExpiry             string        `json:"tax_expiry,omitempty"`
	Speed                 float64       `json:"speed,omitempty"`
	CustomFields          []CustomField `json:"custom_fields,omitempty"`
	Skills                []string      `json:"skills,omitempty"`
	VehicleParts          []string      `json:"vehicle_parts,omitempty"`
	CustomFieldAttributes *CustomField  `json:"custom_fields_attributes,omitempty"` // for creation & update
	SkillList             []string      `json:"skill_list,omitempty"`               // for creation & update
}
