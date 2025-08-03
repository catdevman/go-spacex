package spacex

import (
	"context"
	"fmt"
)

// LaunchpadsService handles communication with the launchpad related
// methods of the SpaceX API.
type LaunchpadsService struct {
	client *Client
}

// Launchpad represents a SpaceX launchpad.
type Launchpad struct {
	Name            *string  `json:"name"`
	FullName        *string  `json:"full_name"`
	Status          string   `json:"status"`
	Locality        *string  `json:"locality"`
	Region          *string  `json:"region"`
	Timezone        *string  `json:"timezone"`
	Latitude        *float64 `json:"latitude"`
	Longitude       *float64 `json:"longitude"`
	LaunchAttempts  int      `json:"launch_attempts"`
	LaunchSuccesses int      `json:"launch_successes"`
	Rockets         []string `json:"rockets"`
	Launches        []string `json:"launches"`
}

// ListAllLaunchpads lists all launchpads.
func (s *LaunchpadsService) ListAllLaunchpads(ctx context.Context) ([]*Launchpad, error) {
	u := "launchpads"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var launchpads []*Launchpad
	_, err = s.client.do(ctx, req, &launchpads)
	if err != nil {
		return nil, err
	}

	return launchpads, nil
}

// GetLaunchpad retrieves a specific launchpad.
func (s *LaunchpadsService) GetLaunchpad(ctx context.Context, id string) (*Launchpad, error) {
	u := fmt.Sprintf("launchpads/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	launchpad := new(Launchpad)
	_, err = s.client.do(ctx, req, launchpad)
	if err != nil {
		return nil, err
	}

	return launchpad, nil
}
