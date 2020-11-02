package coder

import (
	"context"
	"fmt"
	"net/http"
)

// DevURL is the parsed json response record for a devURL from cemanager.
type DevURL struct {
	ID     string `json:"id"     table:"ID"`
	URL    string `json:"url"    table:"URL"`
	Port   int    `json:"port"   table:"Port"`
	Access string `json:"access" table:"Access"`
	Name   string `json:"name"   table:"Name"`
}

type delDevURLRequest struct {
	EnvID    string `json:"environment_id"`
	DevURLID string `json:"url_id"`
}

// DelDevURL deletes the specified devurl.
func (c Client) DelDevURL(ctx context.Context, envID, urlID string) error {
	reqURL := fmt.Sprintf("/api/environments/%s/devurls/%s", envID, urlID)

	resp, err := c.request(ctx, http.MethodDelete, reqURL, delDevURLRequest{
		EnvID:    envID,
		DevURLID: urlID,
	})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }() // Best effort. Likely connection drop.

	if resp.StatusCode != http.StatusOK {
		return bodyError(resp)
	}

	return nil
}

type createDevURLRequest struct {
	EnvID  string `json:"environment_id"`
	Port   int    `json:"port"`
	Access string `json:"access"`
	Name   string `json:"name"`
}

// InsertDevURL inserts a new devurl for the authenticated user.
func (c Client) InsertDevURL(ctx context.Context, envID string, port int, name, access string) error {
	reqURL := fmt.Sprintf("/api/environments/%s/devurls", envID)

	resp, err := c.request(ctx, http.MethodPost, reqURL, createDevURLRequest{
		EnvID:  envID,
		Port:   port,
		Access: access,
		Name:   name,
	})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }() // Best effort. Likely connection drop.

	if resp.StatusCode != http.StatusOK {
		return bodyError(resp)
	}

	return nil
}

type updateDevURLRequest createDevURLRequest

// UpdateDevURL updates an existing devurl for the authenticated user.
func (c Client) UpdateDevURL(ctx context.Context, envID, urlID string, port int, name, access string) error {
	reqURL := fmt.Sprintf("/api/environments/%s/devurls/%s", envID, urlID)

	resp, err := c.request(ctx, http.MethodPut, reqURL, updateDevURLRequest{
		EnvID:  envID,
		Port:   port,
		Access: access,
		Name:   name,
	})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }() // Best effort. Likefly connection drop.

	if resp.StatusCode != http.StatusOK {
		return bodyError(resp)
	}

	return nil
}
