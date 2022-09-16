package utils

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/damianino/gradient-graphql/graph/model"
)

func GenerateMissingValues(input *model.NewGradient) *model.Gradient{
	if len(input.Params.Stops) == 0{
		stops := []string{randomHexColor(), randomHexColor()}
		input.Params.Stops = stops
	}
	if len(input.Params.Stops) < 2{
		input.Params.Stops = append(input.Params.Stops, randomHexColor())
	}
	
	return &model.Gradient{
		ID: "",
		Name: input.Name,
		Params: &model.GradientParams{
			X: input.Params.X,
			Y: input.Params.Y,
			Stops: input.Params.Stops,
		},
		Comments: []*model.Comment{},
	}
}

func randomHexColor() string{
	b := make([]byte, 3)
	rand.Read(b)
	return "#" + hex.EncodeToString(b)
}
