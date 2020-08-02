package entclient

import (
	"time"
)

// User describes a Coder user account
type User struct {
	ID        string    `json:"id" tab:"omit"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" tab:"omit"`
}

// Me gets the details of the authenticated user
func (c Client) Me() (*User, error) {
	var u User
	err := c.requestBody("GET", "/api/users/me", nil, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// SSHKey describes an SSH keypair
type SSHKey struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// SSHKey gets the current SSH kepair of the authenticated user
func (c Client) SSHKey() (*SSHKey, error) {
	var key SSHKey
	err := c.requestBody("GET", "/api/users/me/sshkey", nil, &key)
	if err != nil {
		return nil, err
	}
	return &key, nil
}
