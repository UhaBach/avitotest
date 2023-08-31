package models

type User struct {
	Id   int
	Name string
	Segments []*Segment
}