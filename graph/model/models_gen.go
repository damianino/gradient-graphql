// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	Body string `json:"body"`
}

type Gradient struct {
	ID       string          `json:"_id" bson:"_id"`
	Name     string          `json:"name"`
	Params   *GradientParams `json:"params"`
	Comments []*Comment      `json:"comments"`
}

type GradientParams struct {
	X     int      `json:"x"`
	Y     int      `json:"y"`
	Stops []string `json:"stops"`
}

type NewComment struct {
	ID      string `json:"_id" validate:"id"`
	Comment string `json:"comment" validate:"comment"`
}

type NewGradient struct {
	Name   string     `json:"name" validate:"name"`
	Params *NewParams `json:"params"`
}

type NewParams struct {
	X     int      `json:"x" validate:"imgsize"`
	Y     int      `json:"y" validate:"imgsize"`
	Stops []string `json:"stops" validate:"dive,hexcolor"`
}
