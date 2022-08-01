package usecase

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/User/data"
	"log"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUserCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(uuc domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUserCase{
		userData: uuc,
		validate: v,
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

func (uuc *userUserCase) LoginUser(authData domain.LoginAuth) (data map[string]interface{}, err error) {
	data, err = uuc.userData.LoginData(authData)
	return data, err
}

func (uuc *userUserCase) AccountUser(userid int) (domain.User, []domain.OrderHistory, int) {
	myaccount := uuc.userData.AccountUserData(userid)
	myorder := uuc.userData.HistoryUserData(userid)

	if myaccount.ID == 0 {
		log.Println("Data not found")
		return domain.User{}, nil, 404
	}
	return myaccount, myorder, 200
}
