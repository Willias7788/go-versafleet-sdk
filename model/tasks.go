package model

type TaskResponse struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID                       int                  `json:"id"`
	GUID                     string               `json:"guid"`
	AccountCreatedID         *int                 `json:"account_created_id"`
	AccountID                *int                 `json:"account_id"`
	Price                    float64              `json:"price"`
	InvoiceNumber            string               `json:"invoice_number"`
	TrackingID               string               `json:"tracking_id"`
	TimeFrom                 string               `json:"time_from"`
	TimeTo                   string               `json:"time_to"`
	TimeType                 string               `json:"time_type"`
	TimeWindowID             *int                 `json:"time_window_id"`
	ExpectedCOD              float64              `json:"expected_cod"`
	Remarks                  string               `json:"remarks"`
	ServiceTime              int                  `json:"service_time"`
	ActualTime               *string              `json:"actual_time"`
	ActualLatitude           string               `json:"actual_latitude"`
	ActualLongitude          string               `json:"actual_longitude"`
	EPODURL                  string               `json:"epod_url"`
	AllocatedAccountIDs      []int                `json:"allocated_account_ids"`
	Invoiced                 bool                 `json:"invoiced"`
	RecipientName            *string              `json:"recipient_name"`
	ActualCOD                *float64             `json:"actual_cod"`
	State                    string               `json:"state"`
	Role                     string               `json:"role"`
	Archived                 bool                 `json:"archived"`
	StateUpdatedAt           string               `json:"state_updated_at"`
	LastStartedAt            string               `json:"last_started_at"`
	LastSuccessfulAt         string               `json:"last_successful_at"`
	IsPartialSuccess         bool                 `json:"is_partial_success"`
	JobID                    int                  `json:"job_id"`
	LatestFailureReason      string               `json:"latest_failure_reason"`
	FirstIncompletionNotes   string               `json:"first_incompletion_notes"`
	SecondIncompletionNotes  string               `json:"second_incompletion_notes"`
	ThirdIncompletionNotes   string               `json:"third_incompletion_notes"`
	LastIncompletedLatitude  float64              `json:"last_incompleted_latitude"`
	LastIncompletedLongitude float64              `json:"last_incompleted_longitude"`
	LastCompletedLatitude    float64              `json:"last_completed_latitude"`
	LastCompletedLongitude   float64              `json:"last_completed_longitude"`
	DriverCustomNotes        []DriverNote         `json:"driver_custom_notes"`
	LineItemValidation       []LineItemValidation `json:"line_item_validation"`
	Address                  Address              `json:"address"`
	Job                      Job                  `json:"job"`
	Measurements             []Measurement        `json:"measurements"`
	CustomFields             []CustomField        `json:"custom_fields"`
	Tags                     []Tag                `json:"tags"`
	TaskAssignment           *TaskAssignment      `json:"task_assignment"`
	VehicleSkills            []Skill              `json:"vehicle_skills"`
	VehiclePartSkills        []Skill              `json:"vehicle_part_skills"`
	DriverSkills             []Skill              `json:"driver_skills"`
}

type LineItemValidation struct {
	ID                   int    `json:"id"`
	SKU                  string `json:"sku"`
	Description          string `json:"description"`
	CustomItemID         string `json:"custom_item_id"`
	ExpectedQuantity     int    `json:"expected_quantity"`
	ActualQuantity       int    `json:"actual_quantity"`
	ReasonForDiscrepancy string `json:"reason_for_discrepancy"`
}

type TaskAssignment struct {
	EstimatedStartTime string  `json:"estimated_start_time"`
	Driver             Entity  `json:"driver"`
	Vehicle            Vehicle `json:"vehicle"`
	VehiclePart        Vehicle `json:"vehicle_part"`
	Attendant          Entity  `json:"attendant"`
}

type Entity struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ContactNumber string `json:"contact_number"`
}

type CustomField struct {
	ID                       int    `json:"id"`
	CustomFieldDescriptionID int    `json:"custom_field_description_id"`
	Value                    string `json:"value"`
	Subvalue                 string `json:"subvalue,omitempty"`
}

type DriverNote struct {
	ID                       int    `json:"id"`
	CustomFieldDescriptionID int    `json:"custom_field_description_id"`
	Value                    string `json:"value"`
}

type Skill struct {
	Name string `json:"name"`
}

type TaskListOptions struct {
	// Date & Time Filters (Optional)
	Date          *string `url:"date,omitempty" json:"date,omitempty"`
	FromDateTime  *string `url:"from_datetime,omitempty" json:"from_datetime,omitempty"`
	ToDateTime    string  `url:"to_datetime,omitempty" json:"to_datetime,omitempty"`
	FromCreatedAt *string `url:"from_created_at,omitempty" json:"from_created_at,omitempty"`
	ToCreatedAt   *string `url:"to_created_at,omitempty" json:"to_created_at,omitempty"`
	TimeType      *string `url:"time_type,omitempty" json:"time_type,omitempty"`

	// Search & State (Optional)
	Keyword    *string `url:"keyword,omitempty" json:"keyword,omitempty"`
	State      *string `url:"state,omitempty" json:"state,omitempty"`
	Archived   *bool   `url:"archived,omitempty" json:"archived,omitempty"`
	TrackingID *string `url:"tracking_id,omitempty" json:"tracking_id,omitempty"`

	// Identifiers (Optional)
	CustomerID *int `url:"customer_id,omitempty" json:"customer_id,omitempty"`
	JobID      *int `url:"job_id,omitempty" json:"job_id,omitempty"`
	ID         *int `url:"id,omitempty" json:"id,omitempty"`

	// Pagination (Often has defaults)
	ListOptions
	SortBy  *string `url:"sort_by,omitempty" json:"sort_by,omitempty"`
	OrderBy *string `url:"order_by,omitempty" json:"order_by,omitempty"`
}
