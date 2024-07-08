package user

import (
	"Practice/internal/model"
	"Practice/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var db = repository.InitDB()

func InRepo(user model.User) {
	var err error
	err = db.AutoMigrate(&userRepo{})
	if err != nil {
		panic(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u := userRepo{Name: user.Name, Email: user.Email, Password: string(hashedPassword)}
	result := db.Create(&u)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}

func GetFromEmail(email string) (*model.User, error) {
	var u userRepo
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
