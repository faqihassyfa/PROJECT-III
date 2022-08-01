package usecase

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/User/data"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type userUserCase struct {
	userData domain.UserData
}

func New(uuc domain.UserData) domain.UserUseCase {
	return &userUserCase{
		userData: uuc,
	}
}

// Register implementasi domain.UserUseCase
func (uuc *userUserCase) RegisterUser(newuser domain.User, IDuser int) int {
	var user = data.FromModel(newuser)

	hashed, hasherr := bcrypt.GenerateFromPassword([]byte(user.Password), IDuser)

	if hasherr != nil {
		log.Println("Duplicate Data User")
		return 500
	}
	user.Password = string(hashed)
	insert := uuc.userData.RegisterData(user.ToModel())

	if insert.ID == 0 {
		log.Println("Empty Data")
		return 404
	}
	return 200
}
