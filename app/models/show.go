package models

import "gopkg.in/guregu/null.v3/zero"

type Show struct {
	Id                 int         `json:"id"`
	Name               string      `json:"name"`
	Active             string      `json:"active"`
	Category1          int         `json:"category1"`
	Category2          int         `json:"category2"`
	Category3          int         `json:"category3"`
	Category4          int         `json:"category4"`
	Category5          int         `json:"category5"`
	Category6          int         `json:"category6"`
	Category7          int         `json:"category7"`
	Age                string      `json:"age"`
	WeeklyNut          zero.Float  `json:"weeklynut"`
	NumberOfCast       zero.String `json:"numberofcast"`
	NumberOfMusicians  int         `json:"numberofmusicians"`
	NumberOfStageHands int         `json:"numberofstagehands"`
	NumberOfTrucks     zero.Int    `json:"numberoftrucks"`
	Notes              zero.String `json:"notes"`
}

type Shows struct {
	Shows []Show `json:"shows"`
}
