package model

type Appointment struct {
	ID       int    `json:"id"`
	Date     string `json:"date"`
	Admin_id int    `json:"admin_id"`
	Pet_id   int    `json:"pet_id"`
	OwnerID  int    `json:"owner_id"`
}
