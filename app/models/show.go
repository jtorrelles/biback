package models

type Show struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Shows struct {
	Shows []Show `json:"shows"`
}
