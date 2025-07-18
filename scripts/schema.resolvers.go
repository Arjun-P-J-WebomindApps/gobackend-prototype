package main

// // This file will be automatically regenerated based on the schema, any resolver implementations
// // will be copied through when generating and any unknown code will be moved to the end.
// // Code generated by github.com/99designs/gqlgen version v0.17.76

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/db/models"
// 	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/graph/model"
// 	"github.com/google/uuid"
// )

// // CreateUser is the resolver for the createUser field.
// func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
// 	id := uuid.New()

// 	user, err := r.DB.Queries.CreateUser(ctx, models.CreateUserParams{
// 		ID:        id,
// 		Name:      input.Name,
// 		Username:  input.Username,
// 		Email:     input.Email,
// 		Password:  input.Password,
// 		Mobile:    input.Mobile,
// 		Role:      input.Role,
// 		IsActive:  true,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	})

// 	if err != nil {
// 		log.Println("Couldn't add user " + err.Error())
// 		fmt.Println("error for ", input.Name, input.Username)
// 		return nil, fmt.Errorf("could not create user: %w", err)
// 	}

// 	return &user, nil
// }

// // Users is the resolver for the users field.
// func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
// 	// panic(fmt.Errorf("not implemented: Users - users"))

// 	users, err := r.DB.Queries.GetAllUsers(ctx)

// 	var userPointer []*models.User

// 	for i := range users {
// 		userPointer = append(userPointer, &users[i])
// 	}

// 	if err != nil {
// 		log.Println("Couldn't get all users" + err.Error())
// 		return userPointer, nil
// 	}

// 	return userPointer, nil
// }

// // DeletedAt is the resolver for the deletedAt field.
// func (r *userResolver) DeletedAt(ctx context.Context, obj *models.User) (*time.Time, error) {
// 	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
// }

// // Mutation returns MutationResolver implementation.
// func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// // Query returns QueryResolver implementation.
// func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// // User returns UserResolver implementation.
// func (r *Resolver) User() UserResolver { return &userResolver{r} }

// type mutationResolver struct{ *Resolver }
// type queryResolver struct{ *Resolver }
// type userResolver struct{ *Resolver }
