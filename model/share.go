package model

type Address struct {
	ID            int     `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Zip           string  `json:"zip,omitempty"`
	Line1         string  `json:"line_1,omitempty"`
	Line2         string  `json:"line_2,omitempty"`
	Country       string  `json:"country,omitempty"`
	City          string  `json:"city,omitempty"`
	Email         string  `json:"email,omitempty"`
	ContactPerson string  `json:"contact_person,omitempty"`
	ContactNumber string  `json:"contact_number,omitempty"`
	CountryCode   string  `json:"country_code,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
}

type BaseTask struct {
	ID          int      `json:"id"`
	GUID        string   `json:"guid"`
	TimeFrom    string   `json:"time_from"`
	TimeTo      string   `json:"time_to"`
	TimeType    string   `json:"time_type"`
	State       string   `json:"state"`
	Role        string   `json:"role"`
	ServiceTime int      `json:"service_time"`
	Address     *Address `json:"address,omitempty"`
}

type Tag struct {
	Name string `json:"name"`
}

type Measurement struct {
	ID                          int     `json:"id"`
	Quantity                    float64 `json:"quantity"`
	QuantityUnit                string  `json:"quantity_unit"`
	Weight                      float64 `json:"weight"`
	Volume                      float64 `json:"volume"`
	Description                 string  `json:"description"`
	CustomItemID                string  `json:"custom_item_id"`
	CustomItemCheckMethod       string  `json:"custom_item_check_method"`
	CustomItemUnloadCheckMethod string  `json:"custom_item_unload_check_method"`
}

type MeasurementParams struct {
	ID           int     `json:"id,omitempty"`
	Destroy      bool    `json:"_destroy,omitempty"`
	Quantity     float64 `json:"quantity,omitempty"`
	QuantityUnit string  `json:"quantity_unit,omitempty"`
	Weight       float64 `json:"weight,omitempty"`
	Volume       float64 `json:"volume,omitempty"`
	VolumeLength float64 `json:"volume_length,omitempty"`
	VolumeWidth  float64 `json:"volume_width,omitempty"`
	VolumeHeight float64 `json:"volume_height,omitempty"`
	Description  string  `json:"description,omitempty"`
	CustomItemID string  `json:"custom_item_id,omitempty"`
}

type Paginatable interface {
	GetPage() int
	SetPage(int)
	GetPerPage() int
	SetPerPage(int)
}

type ListOptions struct {
	Page    int `url:"page,omitempty" json:"page,omitempty"`
	PerPage int `url:"per_page,omitempty" json:"per_page,omitempty"`
}

func (b *ListOptions) GetPage() int     { return b.Page }
func (b *ListOptions) SetPage(v int)    { b.Page = v }
func (b *ListOptions) GetPerPage() int  { return b.PerPage }
func (b *ListOptions) SetPerPage(v int) { b.PerPage = v }

type CommonListOptions struct {
	ListOptions
	Keyword      *string `url:"keyword,omitempty" json:"keyword,omitempty"`
	State        *string `url:"state,omitempty" json:"state,omitempty"`
	Archived     *bool   `url:"archived,omitempty" json:"archived,omitempty"`
	Date         *string `url:"date,omitempty" json:"date,omitempty"`
	FromDateTime *string `url:"from_datetime,omitempty" json:"from_datetime,omitempty"`
	ToDateTime   *string `url:"to_datetime,omitempty" json:"to_datetime,omitempty"`
	SortBy       *string `url:"sort_by,omitempty" json:"sort_by,omitempty"`
	OrderBy      *string `url:"order_by,omitempty" json:"order_by,omitempty"`
}

type Meta struct {
	TotalPages  int `json:"total"`
	CurrentPage int `json:"page"`
	PerPage     int `json:"per_page"`
}
