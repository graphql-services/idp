package idp

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// UpdatePassword hash password and assign it to pass word attribute
func (n *User) UpdatePassword(p string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(p), 12)
	if err != nil {
		return err
	}
	fmt.Println(p, "=>", string(passwordHash))
	n.PasswordHash = string(passwordHash)
	return nil
}

// ValidatePassword compare password hashes
func (n User) ValidatePassword(p string) error {
	return bcrypt.CompareHashAndPassword([]byte(n.PasswordHash), []byte(p))
}
