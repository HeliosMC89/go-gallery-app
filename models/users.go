package models

import (
	"errors"
	"os"

	"github.com/heliosmc89/gallery-app-with-go/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var (
	// HMACSecretKey is a variable that defines the secret key needed for hmac
	HMACSecretKey = os.Getenv("HMAC_SECRET_KEY")
	// hmac variable to get the hmac we created
	hmac = utils.NewHMAC(HMACSecretKey)

	// ErrorNotFound is returned when a resource cannot be found.
	ErrorNotFound = errors.New("model: resource not found")
	// ErrorInvalidId will be thrown in case of the id is invalid or equal to zero.
	ErrorInvalidId = errors.New("models: ID provided was invalid")
	// ErrorInvalidPassword will be thrown in case of password missmatch
	ErrorInvalidPassword = errors.New("models: incorrect passsword provided")
)

type User struct {
	gorm.Model
	Name          string `gorm:"not null;type:varchar(255)"`
	Email         string `gorm:"not null;type:varchar(100);unique_index"`
	Password      string `gorm:"not null;type:varchar(100)"`
	RememberToken string `gorm:"nullable;unique_index"`
}

type IUserService interface {
	// Create a new user in system
	Create(user *User) error
	// generate a new remember token for a user
	generateRememberToken(user *User) (*User, error)
	// Update user passed through user object
	Update(user *User) error
	// Fetch the user by the provided token
	ByRememberToken(token string) (*User, error)
	// Search user by ID
	ByID(id uint) (*User, error)
	// Delete user by id
	Delete(id uint) error
	// Authenticate is user is authhorize through is credentials
	Authenticate(email, password string) (*User, error)
	// Email will look up the users using the given email.
	ByEmail(email string) (*User, error)
	// first query of provided database
	first(db *gorm.DB, dst interface{}) error
}

type UserService struct {
	db *gorm.DB
}

func NewUserService() IUserService {
	db, err := utils.GetDatabaseConnection()
	if err != nil {
		panic(err)
	}
	return &UserService{
		db: db,
	}
}

// Create function is used to create a users record
func (u *UserService) Create(user *User) error {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashBytes)
	user, err = u.generateRememberToken(user)
	if err != nil {
		return err
	}
	return u.db.Create(user).Error
}

func (u *UserService) generateRememberToken(user *User) (*User, error) {
	if user.RememberToken == "" {
		token, err := utils.RememberToken()
		if err != nil {
			return nil, err
		}
		user.RememberToken = token
	}
	user.RememberToken = hmac.Hash(user.RememberToken)
	return user, nil
}

// Update will update the provided user wth all of the data passed through the user object.
func (u *UserService) Update(user *User) error {
	user, err := u.generateRememberToken(user)
	if err != nil {
		return err
	}
	return u.db.Save(user).Error
}

// ByRememberToken function is used to fetch the user by the provided token.
func (u *UserService) ByRememberToken(token string) (*User, error) {
	var user User
	err := u.first(u.db.Where("remember_token = ?", token), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ByID will look up the users using the given id.
func (u *UserService) ByID(id uint) (*User, error) {
	var user User
	query := u.db.Where("id = ?", id)
	err := u.first(query, &user)
	return &user, err
}

// Delete function will delete the user with the provided id.
func (u *UserService) Delete(id uint) error {
	// gorm will delete all of the records if the id equals to zero.
	if id == 0 {
		return ErrorInvalidId
	}
	// gorm delete needs the primary key with the reference of the object to understand which table we are deleting from.
	user := User{
		Model: gorm.Model{
			ID: id,
		},
	}
	return u.db.Delete(&user).Error
}

// Authenticate function is used to authorize user through his credentials
func (u *UserService) Authenticate(email, password string) (*User, error) {
	user, err := u.ByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrorInvalidPassword
		default:
			return nil, err
		}
	}
	return user, nil
}

// ByEmail function will look up the users using the given email.
func (u *UserService) ByEmail(email string) (*User, error) {
	var user User
	query := u.db.Where(&user, "email = ?", email)
	err := u.first(query, &user)
	return &user, err
}

// first will query the provided database query and it will get the first item returned and place it
// the query, it will return error not found.
func (u *UserService) first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	switch err {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return ErrorNotFound
	default:
		return err
	}
}
