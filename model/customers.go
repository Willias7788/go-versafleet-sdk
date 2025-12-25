package model

type Customer struct {
	ID            int    `json:"id"`
	GUID          string `json:"guid"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	ContactPerson string `json:"contact_person"`
	ContactNumber string `json:"contact_number"`
	LogoURL       string `json:"logo_url"`
	Archived      bool   `json:"archived"`
}

type CustomerDetail struct {
	Customer
	BillingAccounts []BillingAccount `json:"billing_accounts"`
}

type BillingAccount struct {
	ID            int       `json:"id"`
	GUID          string    `json:"guid"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	ContactPerson string    `json:"contact_person"`
	ContactNumber string    `json:"contact_number"`
	Archived      bool      `json:"archived"`
	Addresses     []Address `json:"addresses"`
}
