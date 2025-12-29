package model

import "time"

type Job struct {
	ID       int      `json:"id"`
	GUID     string   `json:"guid"`
	JobType  string   `json:"job_type"`
	Remarks  string   `json:"remarks"`
	State    string   `json:"state"`
	Archived bool     `json:"archived"`
	Customer Customer `json:"customer"`
	BaseTask BaseTask `json:"base_task"`
	Tags     []Tag    `json:"tags"`
	Tasks    []Task   `json:"tasks,omitempty"`
}

type JobResponse struct {
	Job Job `json:"job"`
}

// JobListOptions handles filtering for job list requests
type JobListOptions struct {
	CommonListOptions
	CustomerID *int `url:"customer_id,omitempty" json:"customer_id,omitempty"`
}

// JobParams is used for Creating and Updating Jobs
type JobParams struct {
	JobType            string          `json:"job_type,omitempty"`
	Remarks            string          `json:"remarks,omitempty"`
	TagList            []string        `json:"tag_list,omitempty"`
	CustomerID         int             `json:"customer_id,omitempty"`
	BaseTaskAttributes *BaseTaskParams `json:"base_task_attributes,omitempty"`
	TasksAttributes    []TaskParams    `json:"tasks_attributes,omitempty"`
}

// BaseTaskParams for the main task of a job
type BaseTaskParams struct {
	TimeFrom          *string   `json:"time_from,omitempty"`
	TimeTo            *string   `json:"time_to,omitempty"`
	TimeType          *TimeType `json:"time_type,omitempty"` // am - Morning time range from 00:00 till 12:00, pm -  Afternoon time range from 12:00 till 23:59, all_day - All day time range ranging from 00:00 till 23:59, custom - Pickup time range is specified by `time_from` and `time_to` attributes, null - If time_window_id is provided
	TimeWindowID      *int      `json:"time_window_id,omitempty"`
	ServiceTime       int       `json:"service_time,omitempty"`
	AddressAttributes *Address  `json:"address_attributes,omitempty"`
	BillingAccountID  *int      `json:"billing_account_id,omitempty"`
}

// TaskParams for sub-tasks of a job
type TaskParams struct {
	ID                   int                 `json:"id,omitempty"`       // For Update
	Destroy              bool                `json:"_destroy,omitempty"` // For Update
	Price                float64             `json:"price,omitempty"`
	InvoiceNumber        string              `json:"invoice_number,omitempty"`
	TrackingID           string              `json:"tracking_id,omitempty"`
	TimeFrom             *time.Time          `json:"time_from,omitempty"`
	TimeTo               *time.Time          `json:"time_to,omitempty"`
	TimeType             string              `json:"time_type,omitempty"`
	TimeWindowID         *int                `json:"time_window_id,omitempty"`
	ExpectedCod          float64             `json:"expected_cod,omitempty"`
	Remarks              string              `json:"remarks,omitempty"`
	ServiceTime          int                 `json:"service_time,omitempty"`
	AddressAttributes    *Address            `json:"address_attributes,omitempty"`
	Measurements         []MeasurementParams `json:"measurements_attributes,omitempty"`
	CustomFieldGroupID   int                 `json:"custom_field_group_id,omitempty"`
	CustomFields         []CustomField       `json:"custom_fields_attributes,omitempty"`
	TagList              []string            `json:"tag_list,omitempty"`
	Photos               []Photo             `json:"photos_attributes,omitempty"`
	VehicleSkillList     []string            `json:"vehicle_skill_list,omitempty"`
	VehiclePartSkillList []string            `json:"vehicle_part_skill_list,omitempty"`
	DriverSkillList      []string            `json:"driver_skill_list,omitempty"`
}

type Photo struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type TimeType string

const (
	TimeTypeAM     TimeType = "am"
	TimeTypePM     TimeType = "pm"
	TimeTypeAllDay TimeType = "all_day"
	TimeTypeCustom TimeType = "custom"
)
