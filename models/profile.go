package models

import (
	"math"
	"time"
)

type Profile struct {
	Base
	Name     string
	Age      float64
	Birthday string
	Parent   []string
}

func NewProfile(Name string, Birthday string, Parent []string) *Profile {
	u := &Profile{
		Name:     Name,
		Age:      GenAge(Birthday),
		Birthday: Birthday,
		Parent:   Parent,
	}

	return u
}

func GenAge(Birthday string) float64 {
	t1, err := time.Parse(time.RFC3339, Birthday)
	if err != nil {

	}

	tnow := time.Now()
	tfinal := math.Floor(tnow.Sub(t1).Hours() / 24 / 365)
	return tfinal
}

/*
type User struct {
  modelImpl
  UserName string
  FullName string
  Email    string
}

func NewUser(userName, fullName, email string) *User {
  u := &User{
    UserName: userName,
    FullName: fullName,
    Email:    email,
  }
  u.SetId(userName)
  return u
}

func (u *User) GetId() string {
  return u.UserName
}
*/
