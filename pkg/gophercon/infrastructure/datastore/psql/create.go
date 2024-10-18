package psql

import (
	"context"
)

// CreateUser create a new user
func (d *PGInstance) CreateUser(ctx context.Context, user *User) (*User, error) {
	if err := d.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
