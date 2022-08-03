package usecase

import (
	"PROJECT-III/domain"
	"PROJECT-III/features/User/data"
	"errors"
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
func (uuc *userUserCase) RegisterUser(newuser domain.User) int {
	var user = data.FromModel(newuser)
	validError := uuc.validate.Struct(user)

	if validError != nil {
		log.Println("Validation error : ", validError)
		return 400
	}

	duplicate := uuc.userData.CheckDuplicate(user.ToModel())

	if duplicate {
		log.Println("Duplicate Data")
		return 400
	}

	hashed, hasherr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if hasherr != nil {
		log.Println("Cant encrypt: ", hasherr)
		return 500
	}

	user.Password = string(hashed)
	user.Role = "user"
	insert := uuc.userData.RegisterData(user.ToModel())

	if insert.ID == 0 {
		log.Println("Empty Data")
		return 404
	}
	return 200
}

// DeleteUser implementasi domain.UserUseCase
func (uuc *userUserCase) DeleteUser(userID int) int {
	delete := uuc.userData.DeleteData(userID)

	if !delete {
		log.Println("Data Not Found")
		return 404
	}
	return 200
}

func (uuc *userUserCase) LoginUser(userdata domain.User) (domain.User, error) {
	login := uuc.userData.LoginData(userdata)

	if login.ID == 0 {
		return domain.User{}, errors.New("no data")
	}

	hashpw := uuc.userData.GetPasswordData(userdata.Email)

	err := bcrypt.CompareHashAndPassword([]byte(hashpw), []byte(userdata.Password))

	if err != nil {
		log.Println(bcrypt.ErrMismatchedHashAndPassword, err)
		return domain.User{}, err
	}

	return login, nil
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

func (uuc *userUserCase) UpdateUser(updatedData domain.User, userid int) int {
	var user = data.FromModel(updatedData)

	if userid == 0 {
		log.Println("Data not found")
		return 404
	}

	duplicate := uuc.userData.CheckDuplicate(user.ToModel())

	if duplicate {
		log.Println("Duplicate Data")
		return 400
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("Error encrypt password", err)
		return 500
	}

	updatedData.Password = string(hashed)
	res := uuc.userData.UpdateUserData(updatedData, userid)
	if res.ID == 0 {
		return 500
	}

	return 200
}

func (uuc *userUserCase) AllProduct() ([]domain.Product, int) {
	read := uuc.userData.AllProductData()

	if len(read) == 0 {
		return nil, 404
	}

	return read, 200
}
