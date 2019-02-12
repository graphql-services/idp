package resolvers

import "github.com/graphql-services/idp/database"

// Resolver ...
type Query struct {
	db *database.DB
}

// NewQuery ...
func NewQuery(db *database.DB) Query {
	return Query{db}
}
