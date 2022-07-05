package models

type FictionalUserModel struct {
	Salutation_type string `json:"salutation_type"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Title_id        int    `json:"title_id,omitempty"`
	Is_superadmin   int    `json:"is_superadmin,omitempty"`
	Is_accountant   int    `json:"is_accountant,omitempty"`
}
