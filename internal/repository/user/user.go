package user

import (
	"Practice/internal/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InRepo(user model.User, db *gorm.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u := UserRepo{Name: user.Name, Email: user.Email, Password: string(hashedPassword)}
	fmt.Println(u)
	result := db.Create(&u)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}

func GetFromEmail(email string, db *gorm.DB) (*model.User, error) {
	var u UserRepo
	result := db.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	mUser := model.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	return &mUser, nil
}
