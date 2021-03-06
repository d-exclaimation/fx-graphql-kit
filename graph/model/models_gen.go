// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewThought struct {
	Title    string  `json:"title"`
	Body     string  `json:"body"`
	ImageURL *string `json:"imageURL"`
	UserID   int     `json:"userId"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
