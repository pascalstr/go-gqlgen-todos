package graph

import (
	"context"
	"math/rand"

	// "crypto/rand"
	"fmt"

	"github.com/pascalstr/gqlgen-todos/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}

// func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
// 	todo := &model.Todo{
// 		Text: input.Text,
// 		ID:   fmt.Sprintf("T%d", rand.Intn(100)),
// 		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
// 	}
// 	r.todos = append(r.todos, todo)
// 	return todo, nil
// }

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Intn(100)),
		UserID: input.UserID, // fix this line
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
