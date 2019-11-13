package models

import (
	"strings"
	"errors"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

)

// User user model
type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(100);unique_index" json:"email"`
	FirstName    string `gorm:"size:100;not null"              json:"firstname"`
	LastName     string `gorm:"size:100;not null"              json:"lastname"`
	Password     string `gorm:"size:100;not null"              json:"password"`      
	ProfileImage string `gorm:"size:255"                       json:"profileimage"`
}

// HashPassword password hashing
func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash enables checking password hash and match password passed
func CheckPasswordHash(password, hash string) error{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}

// BeforeSave will enables password hashing
func (u *User) BeforeSave() error{
	password := strings.TrimSpace(u.Password)
	hashedpassword, err := HashPassword(password)
	if err != nil{
		return err
	}
	u.Password = string(hashedpassword)
	return nil
}

// Prepare will enable cleaning of data
func (u *User) Prepare(){
	u.Email = strings.TrimSpace(u.Email)
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.ProfileImage = strings.TrimSpace(u.ProfileImage)
}

// Validate validated data for login, profile update, picture upload, create users
func (u *User) Validate(action string) error{
	switch strings.ToLower(action) {
	case "login":
		if u.Email == ""{
			return errors.New("Email is required")
		}
		if u.Password == ""{
			return errors.New("Password is required")
		}
		return nil
	default:  // this is the create were all fields are required
		if u.FirstName == ""{
			return errors.New("FirstName is required")
		}
		if u.LastName == ""{
			return errors.New("LastName is required")
		}
		if u.Email == ""{
			return errors.New("Email is required")
		}
		if u.Password == ""{
			return errors.New("Password is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// SaveUser saves user to the database
func (u *User) SaveUser()(*User, error){
	var err error

	// Debug a single operation, show detailed log for this operation
	err = GetDb().Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// GetUser checks if user exists already
func (u *User) GetUser() (*User, error){
	account := &User{}
	if err := GetDb().Debug().Table("users").Where("email = ?", u.Email).First(account).Error; err != nil{
		return nil, err
	}
	return account, nil
}

// GetAllUsers returns all the user
func (u *User) GetAllUsers() (*[]User, error){
	var err error
	users := []User{}
	err = GetDb().Debug().Model(&u).Find(&users).Error
	if err != nil{
		return &[]User{}, err 
	}
	return &users, err
}

