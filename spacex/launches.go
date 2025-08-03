package spacex

import (
	"context"
	"fmt"
)

// LaunchesService handles communication with the launch related
// methods of the SpaceX API.
type LaunchesService struct {
	client *Client
}

// Launch represents a SpaceX launch.
type Launch struct {
	FlightNumber       int           `json:"flight_number"`
	Name               string        `json:"name"`
	DateUTC            string        `json:"date_utc"`
	DateUnix           int64         `json:"date_unix"`
	DateLocal          string        `json:"date_local"`
	DatePrecision      string        `json:"date_precision"`
	StaticFireDateUTC  *string       `json:"static_fire_date_utc"`
	StaticFireDateUnix *int64        `json:"static_fire_date_unix"`
	TDB                bool          `json:"tdb"`
	Net                bool          `json:"net"`
	Window             *int          `json:"window"`
	Rocket             *string       `json:"rocket"`
	Success            *bool         `json:"success"`
	Failures           []*Failure    `json:"failures"`
	Upcoming           bool          `json:"upcoming"`
	Details            *string       `json:"details"`
	Fairings           *Fairings     `json:"fairings"`
	Crew               []string      `json:"crew"`
	Ships              []string      `json:"ships"`
	Capsules           []string      `json:"capsules"`
	Payloads           []string      `json:"payloads"`
	Launchpad          *string       `json:"launchpad"`
	Cores              []*CoreLaunch `json:"cores"`
	Links              *LaunchLinks  `json:"links"`
	AutoUpdate         bool          `json:"auto_update"`
	ID                 string        `json:"id"`
}

// Failure represents a launch failure.
type Failure struct {
	Time     int    `json:"time"`
	Altitude int    `json:"altitude"`
	Reason   string `json:"reason"`
}

// Fairings represents launch fairings.
type Fairings struct {
	Reused          *bool    `json:"reused"`
	RecoveryAttempt *bool    `json:"recovery_attempt"`
	Recovered       *bool    `json:"recovered"`
	Ships           []string `json:"ships"`
}

// CoreLaunch represents a core used in a launch.
type CoreLaunch struct {
	Core           *string `json:"core"`
	Flight         *int    `json:"flight"`
	Gridfins       *bool   `json:"gridfins"`
	Legs           *bool   `json:"legs"`
	Reused         *bool   `json:"reused"`
	LandingAttempt *bool   `json:"landing_attempt"`
	LandingSuccess *bool   `json:"landing_success"`
	LandingType    *string `json:"landing_type"`
	Landpad        *string `json:"landpad"`
}

// LaunchLinks represents links related to a launch.
type LaunchLinks struct {
	Patch     *Patch  `json:"patch"`
	Reddit    *Reddit `json:"reddit"`
	Flickr    *Flickr `json:"flickr"`
	Presskit  *string `json:"presskit"`
	Webcast   *string `json:"webcast"`
	YoutubeID *string `json:"youtube_id"`
	Article   *string `json:"article"`
	Wikipedia *string `json:"wikipedia"`
}

// Patch represents launch patch images.
type Patch struct {
	Small *string `json:"small"`
	Large *string `json:"large"`
}

// Reddit represents reddit links.
type Reddit struct {
	Campaign *string `json:"campaign"`
	Launch   *string `json:"launch"`
	Media    *string `json:"media"`
	Recovery *string `json:"recovery"`
}

// Flickr represents flickr images.
type Flickr struct {
	Small    []string `json:"small"`
	Original []string `json:"original"`
}

// LaunchQueryResults represents the result of a launch query.
type LaunchQueryResults struct {
	Docs          []*Launch `json:"docs"`
	TotalDocs     int       `json:"totalDocs"`
	Limit         int       `json:"limit"`
	TotalPages    int       `json:"totalPages"`
	Page          int       `json:"page"`
	PagingCounter int       `json:"pagingCounter"`
	HasPrevPage   bool      `json:"hasPrevPage"`
	HasNextPage   bool      `json:"hasNextPage"`
	PrevPage      *int      `json:"prevPage"`
	NextPage      *int      `json:"nextPage"`
}

// ListAllLaunches lists all launches.
func (s *LaunchesService) ListAllLaunches(ctx context.Context) ([]*Launch, error) {
	u := "launches"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var launches []*Launch
	_, err = s.client.do(ctx, req, &launches)
	if err != nil {
		return nil, err
	}

	return launches, nil
}

// GetLaunch retrieves a specific launch.
func (s *LaunchesService) GetLaunch(ctx context.Context, id string) (*Launch, error) {
	u := fmt.Sprintf("launches/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	launch := new(Launch)
	_, err = s.client.do(ctx, req, launch)
	if err != nil {
		return nil, err
	}

	return launch, nil
}

// GetLatestLaunch retrieves the latest launch.
func (s *LaunchesService) GetLatestLaunch(ctx context.Context) (*Launch, error) {
	u := "launches/latest"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	launch := new(Launch)
	_, err = s.client.do(ctx, req, launch)
	if err != nil {
		return nil, err
	}

	return launch, nil
}

// GetNextLaunch retrieves the next launch.
func (s *LaunchesService) GetNextLaunch(ctx context.Context) (*Launch, error) {
	u := "launches/next"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	launch := new(Launch)
	_, err = s.client.do(ctx, req, launch)
	if err != nil {
		return nil, err
	}

	return launch, nil
}

// ListPastLaunches lists past launches.
func (s *LaunchesService) ListPastLaunches(ctx context.Context) ([]*Launch, error) {
	u := "launches/past"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var launches []*Launch
	_, err = s.client.do(ctx, req, &launches)
	if err != nil {
		return nil, err
	}

	return launches, nil
}

// ListUpcomingLaunches lists upcoming launches.
func (s *LaunchesService) ListUpcomingLaunches(ctx context.Context) ([]*Launch, error) {
	u := "launches/upcoming"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var launches []*Launch
	_, err = s.client.do(ctx, req, &launches)
	if err != nil {
		return nil, err
	}

	return launches, nil
}

// QueryLaunches queries for launches.
func (s *LaunchesService) QueryLaunches(ctx context.Context, query map[string]interface{}) (*LaunchQueryResults, error) {
	u := "launches/query"
	req, err := s.client.newRequest(ctx, "POST", u, query)
	if err != nil {
		return nil, err
	}

	results := new(LaunchQueryResults)
	_, err = s.client.do(ctx, req, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
