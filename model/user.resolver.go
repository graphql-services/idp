package model

import (
	graphql "github.com/graph-gophers/graphql-go"
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
