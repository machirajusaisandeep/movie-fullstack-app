package models

type Actor struct {
	ID        int
	FirstName string
	LastName  string
	ImageURL  *string // optional - nullable
}
