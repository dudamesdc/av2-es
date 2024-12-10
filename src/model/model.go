package model

type AppointmentsResponse struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	Admin_id  int    `json:"admin_id"`
	Pet_id    int    `json:"pet_id"`
	OwnerID   int    `json:"owner_id"`
	ServiceID int    `json:"service_id"`
}

type Appointments struct {
	Date      string `json:"date"`
	Admin_id  int    `json:"admin_id"`
	Pet_id    int    `json:"pet_id"`
	OwnerID   int    `json:"owner_id"`
	ServiceID int    `json:"service_id"`
}

type PetResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Species   string `json:"species"`
	Breed     string `json:"breed"`
	Age       int    `json:"age"`
	ServiceID int    `json:"service_id"`
	OwnerID   int    `json:"owner_id"` // Relacionamento com o usuário
}
type Pet struct {
	Name      string `json:"name"`
	Species   string `json:"species"`
	Breed     string `json:"breed"`
	Age       int    `json:"age"`
	ServiceID int    `json:"vaccination_id"`
	OwnerID   int    `json:"owner_id"` // Relacionamento com o usuário
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"type_user"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"type_user"`
}

type ServiceResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	Name string `json:"name"`
}
