package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/creativelikeadog/taxiapp.com.br/app/forms"
	"time"
)

type Driver struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name      string        `json:"name"`
	CarPlate  string        `json:"carPlate"`
	Location  [2]float64    `json:"location"`
	Available bool          `json:"available"`
	Created   time.Time     `json:"created_at"`
	Updated   time.Time     `json:"updated_at"`
}

type DriverQuery struct {
	IDs       []bson.ObjectId
	Names     []string
	CarPlates []string
	Area      *forms.Area
	Offset    int
	Limit     int
}

type DriverStatusVO struct {
	ID        bson.ObjectId `json:"driverId"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Available bool          `json:"driverAvailable"`
}
