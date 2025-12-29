package model

type Customer struct {
	ID                        int             `json:"id"`
	GUID                      string          `json:"guid"`
	Name                      string          `json:"name"`
	Email                     string          `json:"email"`
	ContactPerson             string          `json:"contact_person"`
	ContactNumber             string          `json:"contact_number"`
	LogoURL                   string          `json:"logo_url"`
	Archived                  bool            `json:"archived"`
	BillingAccountsAttributes *BillingAccount `json:"billing_accounts_attributes,omitempty"`
}

type CustomerDetail struct {
	Customer
	BillingAccounts []BillingAccount `json:"billing_accounts"`
}

type BillingAccount struct {
	ID                  int       `json:"id,omitempty"`
	GUID                string    `json:"guid,omitempty"`
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	ContactPerson       string    `json:"contact_person"`
	ContactNumber       string    `json:"contact_number"`
	Archived            bool      `json:"archived"`
	Addresses           []Address `json:"addresses,omitempty"`
	AddressesAttributes *Address  `json:"addresses_attributes,omitempty"`
}

type CustomerListOptions struct {
	CommonListOptions
}
