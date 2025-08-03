package spacex

import (
	"context"
	"fmt"
)

// CrewService handles communication with the crew related
// methods of the SpaceX API.
type CrewService struct {
	client *Client
}

// Crew represents a SpaceX crew member.
type Crew struct {
	Name      *string  `json:"name"`
	Status    string   `json:"status"`
	Agency    *string  `json:"agency"`
	Image     *string  `json:"image"`
	Wikipedia *string  `json:"wikipedia"`
	Launches  []string `json:"launches"`
}

// ListAllCrew lists all crew members.
func (s *CrewService) ListAllCrew(ctx context.Context) ([]*Crew, error) {
	u := "crew"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var crew []*Crew
	_, err = s.client.do(ctx, req, &crew)
	if err != nil {
		return nil, err
	}

	return crew, nil
}

// GetCrew retrieves a specific crew member.
func (s *CrewService) GetCrew(ctx context.Context, id string) (*Crew, error) {
	u := fmt.Sprintf("crew/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	crew := new(Crew)
	_, err = s.client.do(ctx, req, crew)
	if err != nil {
		return nil, err
	}

	return crew, nil
}
