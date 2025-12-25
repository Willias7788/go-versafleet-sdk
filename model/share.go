package model

type Address struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Zip           string  `json:"zip"`
	Line1         string  `json:"line_1"`
	Line2         string  `json:"line_2"`
	Country       string  `json:"country"`
	City          string  `json:"city"`
	Email         string  `json:"email"`
	ContactPerson string  `json:"contact_person"`
	ContactNumber string  `json:"contact_number"`
	CountryCode   string  `json:"country_code"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
}

type BaseTask struct {
	ID          int     `json:"id"`
	GUID        string  `json:"guid"`
	TimeFrom    string  `json:"time_from"`
	TimeTo      string  `json:"time_to"`
	TimeType    string  `json:"time_type"`
	State       string  `json:"state"`
	Role        string  `json:"role"`
	ServiceTime int     `json:"service_time"`
	Address     Address `json:"address"`
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

type Meta struct {
	TotalPages  int `json:"total"`
	CurrentPage int `json:"page"`
	PerPage     int `json:"per_page"`
}
