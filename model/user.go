package model

type GPIUser struct {
	ID	int `json:"id,omitempty"`
	Lastname        string  `json:"lname,omitempty"`
	Firstname string  `json:"fname,omitempty"`
	Dob       string `json:"dob,omitempty"`
	Sex       string `json:"gender,omitempty"`
	Email     string  `json:"myemail,omitempty"`
	Address string  `json:"address,omitempty"`
	City       string `json:"city,omitempty"`
	Username       string `json:"username,omitempty"`
	Picture	string `json:"picture,omitempty"`
	Company       string `json:"company,omitempty"`
}

type GPICompanyDetails struct {
	ID	int `json:"id,omitempty"`
	AccountOwner        string  `json:"account_owner,omitempty"`
	AccountID string  `json:"account_id,omitempty"`
	Company       string `json:"company,omitempty"`
	CompanyName       string `json:"companyname,omitempty"`
}