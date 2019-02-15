package idp

import (
	"context"
	"fmt"

	"github.com/graphql-services/idp/database"
	uuid "github.com/satori/go.uuid"
)

type Resolver struct {
	DB *database.DB
}

func (r *Resolver) getUserStrict(ctx context.Context, email string) (User, error) {
	var user User
	query := r.DB.Query().Where(&User{Email: email}).First(&user)
	if query.RecordNotFound() {
		return user, fmt.Errorf("user not found")
	}
	return user, query.Error
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input UserInput) (User, error) {
	ID := uuid.Must(uuid.NewV4())
	user := User{ID: ID.String(), Email: input.Email, GivenName: input.GivenName, FamilyName: input.FamilyName, MiddleName: input.MiddleName}

	if err := user.UpdatePassword(input.Password); err != nil {
		return user, err
	}

	res := r.DB.Query().Create(&user)

	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (User, error) {
	var user User
	query := r.DB.Query().Where(&User{ID: id}).First(&user)

	if query.RecordNotFound() {
		return user, fmt.Errorf("not found")
	}
	return user, query.Error
}
func (r *mutationResolver) VerifyUser(ctx context.Context, email string) (User, error) {
	user, err := r.getUserStrict(ctx, email)
	if err != nil {
		return user, err
	}

	user.EmailVerified = true
	res := r.DB.Query().Save(&user)

	return user, res.Error
}
func (r *mutationResolver) ChangePassword(ctx context.Context, email string, newPassword string) (User, error) {
	user, err := r.getUserStrict(ctx, email)
	if err != nil {
		return user, err
	}
	user.UpdatePassword(newPassword)
	query := r.DB.Query().Save(&user)

	return user, query.Error
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUser(ctx context.Context, email string) (*User, error) {
	var user User
	query := r.DB.Query().Where(&User{Email: email}).First(&user)
	return &user, query.Error
}
func (r *queryResolver) Login(ctx context.Context, email string, password string) (User, error) {
	var user User
	query := r.DB.Query().Where(&User{Email: email}).First(&user)
	maybeUser, err := r.GetUser(ctx, email)
	if err != nil {
		return user, err
	}

	if query.Error != nil {
		return user, query.Error
	}
	if maybeUser == nil {
		return user, fmt.Errorf("User not found")
	}
	user = *maybeUser

	if err := user.ValidatePassword(password); err != nil {
		return user, fmt.Errorf("Invalid password")
	}

	return user, nil
}
