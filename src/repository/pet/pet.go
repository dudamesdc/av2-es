package repositories

import (
	"errors"
	"sync"

	"github.com/dudamesdc/av2-es/src/config"
	model "github.com/dudamesdc/av2-es/src/model"
)

var mu sync.Mutex // Mutex para evitar condições de corrida

// Cria um novo pet
func CreatePet(pet model.Pet, userID int) (model.PetResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	// Verificar se o usuário que está criando o pet realmente existe
	userExists := false
	for _, user := range db.Users {
		if user.ID == userID {
			userExists = true
			break
		}
	}

	if !userExists {
		return model.PetResponse{}, errors.New("user does not exist")
	}

	// Incrementar o ID do pet
	var newID int
	if len(db.Pets) > 0 {
		newID = db.Pets[len(db.Pets)-1].ID + 1
	} else {
		newID = 1
	}

	// Criar o PetResponse
	petResponse := model.PetResponse{
		ID:            newID,
		Name:          pet.Name,
		Species:       pet.Species,
		Breed:         pet.Breed,
		Age:           pet.Age,
		VaccinationID: pet.VaccinationID,
		OwnerID:       userID,
	}

	// Adicionar o PetResponse ao banco de dados
	db.Pets = append(db.Pets, petResponse)

	return petResponse, nil
}

// Atualiza os dados de um pet existente
func UpdatePet(id int, updatedPet model.Pet, userID int) (model.PetResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	// Verificar se o pet pertence ao usuário
	for i, pet := range db.Pets {
		if pet.ID == id {
			if pet.OwnerID != userID {
				return model.PetResponse{}, errors.New("you can only update your own pet")
			}

			// Atualizar os dados do pet
			db.Pets[i].Name = updatedPet.Name
			db.Pets[i].Species = updatedPet.Species
			db.Pets[i].Breed = updatedPet.Breed
			db.Pets[i].Age = updatedPet.Age
			db.Pets[i].VaccinationID = updatedPet.VaccinationID
			return db.Pets[i], nil
		}
	}
	return model.PetResponse{}, errors.New("pet not found")
}

// Retorna um pet específico
func GetPetByID(id int, userID int) (model.PetResponse, error) {
	db := config.GetDatabase()

	for _, pet := range db.Pets {
		if pet.ID == id {
			// Verifica se o pet pertence ao usuário
			if pet.OwnerID == userID {
				return pet, nil
			}
			return model.PetResponse{}, errors.New("you can only access your own pet")
		}
	}
	return model.PetResponse{}, errors.New("pet not found")
}

// Retorna todos os pets (somente para admin)
func GetAllPets() ([]model.PetResponse, error) {
	db := config.GetDatabase()
	return db.Pets, nil
}

func DeletePet(id int, userID int) error {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	// Encontrar o pet com o ID especificado
	for i, pet := range db.Pets {
		if pet.ID == id {
			// Verifica se o pet pertence ao usuário
			if pet.OwnerID != userID {
				return errors.New("you can only delete your own pet")
			}

			// Deletar o pet removendo-o da lista
			db.Pets = append(db.Pets[:i], db.Pets[i+1:]...)
			return nil
		}
	}

	return errors.New("pet not found")
}
