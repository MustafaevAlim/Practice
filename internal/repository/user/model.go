package user

import "gorm.io/gorm"
import "github.com/dgrijalva/jwt-go"

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
