package store

import "http-rest-api/internal/app/model"

// UserRepository is an interface that contains methods for working with users
type UserRepository struct {
	store *Store
}

// Create is a method that creates a new user
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id", u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil

}

// FindByEmail is a method that finds a user by email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}
