package models

import (
	"math"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Base     `bson:",inline"`
	Name     string
	Age      float64
	Birthday string
	Parent   []string
}

func AddProfile(Id bson.ObjectId, Name string, Birthday string, Parent []string) Profile {

	var profile = Profile{}
	profile.Id = Id
	profile.CreatedAt = time.Now().Format(time.RFC3339)
	profile.Birthday = Birthday
	profile.Age = GenAge(Birthday)
	profile.Name = Name
	profile.Parent = Parent

	return profile
}

func GenAge(Birthday string) float64 {
	t1, err := time.Parse(time.RFC3339, Birthday)
	if err != nil {

	}

	tnow := time.Now()
	tfinal := math.Floor(tnow.Sub(t1).Hours() / 24 / 365)
	return tfinal
}
