package resolvers

import (
	"context"
	"fmt"
	"time"

	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/db/models"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/graph"
	"github.com/Arjun-P-J-WebomindApps/gobackend-prototype/internal/graph/model"
)

type userSessionResolver struct{ *Resolver }

// UserSession returns UserSessionResolver implementation.
func (r *Resolver) UserSession() graph.UserSessionResolver { return &userSessionResolver{r} }

// CreateUserSession is the resolver for the createUserSession field.
func (r *mutationResolver) CreateUserSession(ctx context.Context, input model.CreateUserSessionInput) (*models.UserSession, error) {
	panic(fmt.Errorf("not implemented: CreateUserSession - createUserSession"))
}

// CreatedAt is the resolver for the createdAt field.
func (r *userSessionResolver) CreatedAt(ctx context.Context, obj *models.UserSession) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}
