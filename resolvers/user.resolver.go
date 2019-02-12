package resolvers

import (
	"context"
	"fmt"

	"github.com/graphql-services/idp/model"
	uuid "github.com/satori/go.uuid"
)

func (q *Query) getUserStrict(ctx context.Context, email string) (model.User, error) {
	var user model.User
	query := q.db.Query().Where(&model.User{IEmail: email}).First(&user)
	if query.RecordNotFound() {
		return user, fmt.Errorf("user not found")
	}
	return user, query.Error
}

// GetUser ...
func (q *Query) GetUser(ctx context.Context, p struct{ Email string }) (*model.User, error) {
	var user model.User
	query := q.db.Query().Where(&model.User{IEmail: p.Email}).First(&user)
	return &user, query.Error
}

// Login ...
func (q *Query) Login(ctx context.Context, p struct{ Email, Password string }) (model.User, error) {
	var user model.User
	query := q.db.Query().Where(&model.User{IEmail: p.Email}).First(&user)
	maybeUser, err := q.GetUser(ctx, struct{ Email string }{Email: p.Email})
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

	if err := user.ValidatePassword(p.Password); err != nil {
		return user, fmt.Errorf("Invalid password")
	}

	return user, nil
}

type createUserInput struct {
	Email      string
	Password   string
	GivenName  *string
	FamilyName *string
	MiddleName *string
}

// CreateUser ...
func (q *Query) CreateUser(ctx context.Context, p struct{ User createUserInput }) (model.User, error) {
	ID := uuid.Must(uuid.NewV4())
	user := model.User{IID: ID, IEmail: p.User.Email, IGivenName: p.User.GivenName, IFamilyName: p.User.FamilyName, IMiddleName: p.User.MiddleName}

	if err := user.UpdatePassword(p.User.Password); err != nil {
		return user, err
	}

	res := q.db.Query().Create(&user)

	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}

// DeleteUser ...
func (q *Query) DeleteUser(ctx context.Context, p struct{ ID string }) (model.User, error) {
	var user model.User
	ID := uuid.Must(uuid.FromString(p.ID))
	query := q.db.Query().Where(&model.User{IID: ID}).First(&user)

	if query.RecordNotFound() {
		return user, fmt.Errorf("not found")
	}
	return user, query.Error
}

// VerifyUser ...
func (q *Query) VerifyUser(ctx context.Context, p struct{ Email string }) (model.User, error) {
	user, err := q.getUserStrict(ctx, p.Email)
	if err != nil {
		return user, err
	}

	user.IEmailVerified = true
	res := q.db.Query().Save(&user)

	return user, res.Error
}

// ChangePassword ...
func (q *Query) ChangePassword(ctx context.Context, p struct{ Email, NewPassword string }) (model.User, error) {
	user, err := q.getUserStrict(ctx, p.Email)
	if err != nil {
		return user, err
	}
	user.UpdatePassword(p.NewPassword)
	query := q.db.Query().Save(&user)

	return user, query.Error
}
