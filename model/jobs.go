package model

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
}

type JobListOptions struct {

	// Date Filters
	Date         string `url:"date,omitempty" json:"date,omitempty"`
	FromDateTime string `url:"from_datetime,omitempty" json:"from_datetime,omitempty"`
	ToDateTime   string `url:"to_datetime,omitempty" json:"to_datetime,omitempty"`

	// Search & State
	Keyword    string `url:"keyword,omitempty" json:"keyword,omitempty"`
	State      string `url:"state,omitempty" json:"state,omitempty"`
	CustomerID int    `url:"customer_id,omitempty" json:"customer_id,omitempty"`
	Archived   bool   `url:"archived,omitempty" json:"archived,omitempty"`

	// Sorting
	ListOptions        // Handles Page and PerPage
	SortBy      string `url:"sort_by,omitempty" json:"sort_by,omitempty"`
	OrderBy     string `url:"order_by,omitempty" json:"order_by,omitempty"`
}
