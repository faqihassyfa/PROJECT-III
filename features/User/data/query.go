package data

import (
	"PROJECT-III/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

//RegisterData implement domain.UserData
func (ud *userData) RegisterData(newuser domain.User) domain.User {
	var user = FromModel(newuser)
	err := ud.db.Create(&user).Error

	if user.ID == 0 {
		log.Println("Invalid ID")
		return domain.User{}
	}

	if err != nil {
		log.Println("Can not create user object", err.Error())
		return domain.User{}
	}

	return user.ToModel()
}

// DeleteData implementasi domain.UserData
func (ud *userData) DeleteData(userID int) bool {
	res := ud.db.Where("ID = ?", userID).Delete(&User{})
	if res.Error != nil {
		log.Println("Cannot delete data", res.Error.Error())
		return false
	}
	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return false
	}
	return true
}

// func (ud *userData) LoginData(authData domain.LoginAuth) (data map[string]interface{}, err error) {
// 	userData := User{}
// 	res := ud.db.Where("email = ?", authData.Email).First(&userData)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}
// 	errCrypt := _bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(authData.Password))
// 	if errCrypt != nil {
// 		return nil, errors.New("invalid password")
// 	}

// 	token, _ := middlewares.CreateToken(int(userData.ID), userData.Role)

// 	var dataToken = map[string]interface{}{}
// 	dataToken["id"] = int(userData.ID)
// 	dataToken["token"] = token
// 	dataToken["role"] = userData.Role
// 	return dataToken, nil
// }

func (ud *userData) AccountUserData(userid int) domain.User {
	var tmp User
	err := ud.db.Where("ID = ?", userid).First(&tmp).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
		return domain.User{}
	}
	return tmp.ToModel()

}

func (ud *userData) HistoryUserData(userid int) []domain.OrderHistory {
	var tmp []OrderHistory
	err := ud.db.Model(&domain.Order{}).Select("orders.id, orders.created_at, orders.totalprice").Where("orders.userid = ?", userid).Find(&tmp).Error
	if err != nil {
		log.Println("There is problem with data", err.Error())
	}
	fmt.Println("history", tmp)
	return ParseOrderHistoryToArr(tmp)
}

func (ud *userData) UpdateUserData(updatedData domain.User, userid int) domain.User {
	var user = FromModel(updatedData)
	err := ud.db.Model(&user).Where("ID = ?", userid).Updates(updatedData)

	if err.Error != nil {
		log.Println("Cant update user object", err.Error.Error())
		return domain.User{}
	}

	if err.RowsAffected == 0 {
		log.Println("Data Not Found")
		return domain.User{}
	}
	user.ID = uint(userid)
	return user.ToModel()
}

func (ud *userData) CheckDuplicate(newuser domain.User) bool {
	var user User
	err := ud.db.Find(&user, "email = ?", newuser.Email)

	if err.RowsAffected == 1 {
		log.Println("Duplicated data", err.Error)
		return true
	}

	return false
}
