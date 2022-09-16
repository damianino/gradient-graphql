package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/damianino/gradient-graphql/database"
	"github.com/damianino/gradient-graphql/graph/generated"
	"github.com/damianino/gradient-graphql/graph/model"
	"github.com/damianino/gradient-graphql/graph/validators"
)

var v = validators.Validator()

// CreateGradient is the resolver for the createGradient field.
func (r *mutationResolver) CreateGradient(ctx context.Context, input *model.NewGradient) (*model.Gradient, error) {
	if err := v.Struct(input); err != nil{
		return nil, err
	}
	return db.Save(input)
}

// Comment is the resolver for the comment field.
func (r *mutationResolver) Comment(ctx context.Context, input *model.NewComment) (*model.Gradient, error) {
	return db.Comment(input)
}

// Gradient is the resolver for the gradient field.
func (r *queryResolver) Gradient(ctx context.Context, id string) (*model.Gradient, error) {
	return db.FindById(id)
}

// Gradients is the resolver for the gradients field.
func (r *queryResolver) Gradients(ctx context.Context) ([]*model.Gradient, error) {
	return db.All()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
