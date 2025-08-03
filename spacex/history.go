package spacex

import (
	"context"
	"fmt"
)

// HistoryService handles communication with the history related
// methods of the SpaceX API.
type HistoryService struct {
	client *Client
}

// History represents a SpaceX history event.
type History struct {
	Title         *string `json:"title"`
	EventDateUTC  *string `json:"event_date_utc"`
	EventDateUnix *int    `json:"event_date_unix"`
	Details       *string `json:"details"`
	Links         *struct {
		Article *string `json:"article"`
	} `json:"links"`
}

// ListAllHistory lists all history events.
func (s *HistoryService) ListAllHistory(ctx context.Context) ([]*History, error) {
	u := "history"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var history []*History
	_, err = s.client.do(ctx, req, &history)
	if err != nil {
		return nil, err
	}

	return history, nil
}

// GetHistory retrieves a specific history event.
func (s *HistoryService) GetHistory(ctx context.Context, id string) (*History, error) {
	u := fmt.Sprintf("history/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	history := new(History)
	_, err = s.client.do(ctx, req, history)
	if err != nil {
		return nil, err
	}

	return history, nil
}
