package repositories

import (
	"errors"
	"sync"

	"github.com/dudamesdc/av2-es/src/config"
	model "github.com/dudamesdc/av2-es/src/model"
)

var mu sync.Mutex // Mutex para evitar condições de corrida

// Cria um novo agendamento com ID incremental
func CreatePet(Pet model.Pet) (model.PetResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	// Verificar se o e-mail já existe
	// for _, existingPet := range db.Pets {
	// 	if existingPet.Email == Pet.Email {
	// 		return model.PetResponse{}, errors.New("email already registered")
	// 	}
	// }

	// Incrementar o ID
	var newID int
	if len(db.Pets) > 0 {
		newID = db.Pets[len(db.Pets)-1].ID + 1
	} else {
		newID = 1
	}

	// Criar o PetResponse com o ID
	PetResponse := model.PetResponse{
		ID:        newID,
		Name:      Pet.Name,
		Breed:     Pet.Breed,
		Age:       Pet.Age,
		Species:   Pet.Species,
		ServiceID: Pet.ServiceID,
		OwnerID:   Pet.OwnerID,
	}

	// Adicionar o PetResponse ao banco de dados
	db.Pets = append(db.Pets, PetResponse) // Armazenando PetResponse

	return PetResponse, nil
}

// Atualiza os dados de um usuário existente
func UpdatePet(id int, updatedPet model.Pet) (model.PetResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	for i, Pet := range db.Pets {
		if Pet.ID == id {
			// Atualizar apenas os campos permitidos
			db.Pets[i].Name = updatedPet.Name
			db.Pets[i].Breed = updatedPet.Breed
			db.Pets[i].Age = updatedPet.Age
			db.Pets[i].Species = updatedPet.Species
			db.Pets[i].ServiceID = updatedPet.ServiceID

			return db.Pets[i], nil
		}
	}
	return model.PetResponse{}, errors.New("Pet not found")
}

// Remove um usuário do banco de dados
func DeletePet(id int) error {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	for i, Pet := range db.Pets {
		if Pet.ID == id {
			// Remover usuário pelo índice
			db.Pets = append(db.Pets[:i], db.Pets[i+1:]...)
			return nil
		}
	}
	return errors.New("Pet not found")
}

// Retorna todos os usuários
func GetAllPets() []model.PetResponse {
	db := config.GetDatabase()
	return db.Pets
}

func GetPetByID(id int) (model.PetResponse, error) {
	db := config.GetDatabase()
	for _, Pet := range db.Pets {
		if Pet.ID == id {
			return Pet, nil
		}
	}
	return model.PetResponse{}, errors.New("Pet not found")
}
