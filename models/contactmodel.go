package models

type ContactModel struct {
	Contact_type_id    int    `json:"contact_type_id"`
	Name_1             string `json:"name_1"`
	Name_2             string `json:"name_2"`
	Salutation_id      int    `json:"salutation_id,omitempty"`
	Salutation_form    string `json:"salutation_form,omitempty"`
	Titel_id           int    `json:"title_id,omitempty"`
	Birthday           string `json:"birthday,omitempty"`
	Address            string `json:"address,omitempty"`
	Postcode           string `json:"postcode,omitempty"`
	City               string `json:"city,omitempty"`
	Country_id         int    `json:"country_id,omitempty"`
	Mail               string `json:"mail,omitempty"`
	Mail_second        string `json:"mail_second,omitempty"`
	Phone_fixed        string `json:"phone_fixed,omitempty"`
	Phone_fixed_second string `json:"phone_fixed_second,omitempty"`
	Phone_mobile       string `json:"phone_mobile,omitempty"`
	Fax                string `json:"fax,omitempty"`
	Url                string `json:"url,omitempty"`
	Skype_name         string `json:"skype_name,omitempty"`
	Remarks            string `json:"remarks,omitempty"`
	Language_id        int    `json:"language_id,omitempty"`
	Contact_group_ids  string `json:"contact_group_ids,omitempty"`
	Contact_branch_ids string `json:"contact_branch_ids,omitempty"`
	User_id            int    `json:"user_id"`
	Owner_id           int    `json:"owner_id"`
}
