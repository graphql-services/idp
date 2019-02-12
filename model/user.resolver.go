package model

import (
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"golang.org/x/crypto/bcrypt"
)

// ID ...
func (n User) ID() graphql.ID {
	return graphql.ID(n.IID.String())
}

// Email ...
func (n User) Email() string {
	return n.IEmail
}

// EmailVerified ...
func (n User) EmailVerified() bool {
	return n.IEmailVerified
}

// GivenName ...
func (n User) GivenName() *string {
	return n.IGivenName
}

// FamilyName ...
func (n User) FamilyName() *string {
	return n.IFamilyName
}

// MiddleName ...
func (n User) MiddleName() *string {
	return n.IMiddleName
}

// Updatedat ...
func (n User) Updatedat() graphql.Time {
	return graphql.Time{Time: n.UpdatedAt}
}

// UpdatePassword hash password and assign it to pass word attribute
func (n *User) UpdatePassword(p string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	if err != nil {
		return err
	}
	fmt.Println(p, "=>", string(passwordHash))
	n.IPasswordHash = string(passwordHash)
	return nil
}

// ValidatePassword compare password hashes
func (n User) ValidatePassword(p string) error {
	return bcrypt.CompareHashAndPassword([]byte(n.IPasswordHash), []byte(p))
}
