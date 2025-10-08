package users

import (
	"Smradity/internal/users/roles"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Login            string     `json:"login" bson:"login"`
	Email            string     `json:"email" bson:"email"`
	PasswordHash     string     `json:"password" bson:"password"`
	Role             roles.Role `json:"role" bson:"role"`
	CreatedAt        time.Time  `json:"createdAt" bson:"createdAt"`
}

// SetPassword Функция SetPassword хеширует пароль согласно требованиям
func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hash)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}

//Users CRUD

func CreateUser(user *User) error {
	return mgm.Coll(user).Create(user)
}

func GetUserByLogin(login string) (*User, error) {
	user := &User{}
	err := mgm.Coll(user).First(bson.M{"login": login}, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(id string) (*User, error) {
	user := &User{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = mgm.Coll(user).FindByID(objID, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(user *User) error {
	return mgm.Coll(user).Update(user)
}

func DeleteUser(user *User) error {
	return mgm.Coll(user).Delete(user)
}
