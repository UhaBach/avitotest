package models

type Segment struct {
	Id   int
	Name string
	Users []*User
}
