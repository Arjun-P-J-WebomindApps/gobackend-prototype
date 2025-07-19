package resolvers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/db/models"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/graph"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/graph/model"
	"github.com/google/uuid"
)

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	id := uuid.New()

	user, err := r.DB.Queries.CreateUser(ctx, models.CreateUserParams{
		ID:        id,
		Name:      input.Name,
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		Mobile:    input.Mobile,
		Role:      input.Role,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Couldn't add user " + err.Error())
		fmt.Println("error for ", input.Name, input.Username)
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	return &user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	// panic(fmt.Errorf("not implemented: Users - users"))

	users, err := r.DB.Queries.GetAllUsers(ctx)

	var userPointer []*models.User

	for i := range users {
		userPointer = append(userPointer, &users[i])
	}

	if err != nil {
		log.Println("Couldn't get all users" + err.Error())
		return userPointer, nil
	}

	return userPointer, nil
}

// User returns UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

// DeletedAt is the resolver for the deletedAt field.
func (r *userResolver) DeletedAt(ctx context.Context, obj *models.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: DeletedAt - deletedAt"))
}
