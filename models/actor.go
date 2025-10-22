package models

type Actor struct {
	ID       int
	Name     string
	ImageURL *string // optional - nullable
}
