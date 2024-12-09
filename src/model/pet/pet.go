package model

type Pet struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Species       string `json:"species"`
	Breed         string `json:"breed"`
	Age           int    `json:"age"`
	VaccinationID int    `json:"vaccination_id"`
	OwnerID       int   `json:"owner_id"` // Relacionamento com o usuário
}
