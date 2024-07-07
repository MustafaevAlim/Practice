package user

import (
	"Practice/internal/model"
	"Practice/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db = repository.InitDB()

type userRepo struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(200)"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func AddUserInRepo(user model.User) {
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

func GetUserEmail(email string) (*model.User, error) {
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
