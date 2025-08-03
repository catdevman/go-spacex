package spacex

import (
	"context"
	"fmt"
)

// RocketsService handles communication with the rocket related
// methods of the SpaceX API.
type RocketsService struct {
	client *Client
}

// Rocket represents a SpaceX rocket.
type Rocket struct {
	Name           string          `json:"name"`
	Type           string          `json:"type"`
	Active         bool            `json:"active"`
	Stages         int             `json:"stages"`
	Boosters       int             `json:"boosters"`
	CostPerLaunch  int             `json:"cost_per_launch"`
	SuccessRatePct int             `json:"success_rate_pct"`
	FirstFlight    string          `json:"first_flight"`
	Country        string          `json:"country"`
	Company        string          `json:"company"`
	Height         *Dimension      `json:"height"`
	Diameter       *Dimension      `json:"diameter"`
	Mass           *Mass           `json:"mass"`
	PayloadWeights []PayloadWeight `json:"payload_weights"`
	FirstStage     *FirstStage     `json:"first_stage"`
	SecondStage    *SecondStage    `json:"second_stage"`
	Engines        *Engines        `json:"engines"`
	LandingLegs    *LandingLegs    `json:"landing_legs"`
	FlickrImages   []string        `json:"flickr_images"`
	Wikipedia      string          `json:"wikipedia"`
	Description    string          `json:"description"`
	ID             string          `json:"id"`
}

// Dimension represents height and diameter measurements.
type Dimension struct {
	Meters *float64 `json:"meters"`
	Feet   *float64 `json:"feet"`
}

// Mass represents mass measurements.
type Mass struct {
	Kg *int `json:"kg"`
	Lb *int `json:"lb"`
}

// Thrust represents thrust measurements.
type Thrust struct {
	KN  *int `json:"kN"`
	Lbf *int `json:"lbf"`
}

// PayloadWeight represents payload weight information.
type PayloadWeight struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Kg   int    `json:"kg"`
	Lb   int    `json:"lb"`
}

// FirstStage represents the first stage of a rocket.
type FirstStage struct {
	Reusable       bool    `json:"reusable"`
	Engines        int     `json:"engines"`
	FuelAmountTons float64 `json:"fuel_amount_tons"`
	BurnTimeSec    *int    `json:"burn_time_sec"`
	ThrustSeaLevel *Thrust `json:"thrust_sea_level"`
	ThrustVacuum   *Thrust `json:"thrust_vacuum"`
}

// SecondStage represents the second stage of a rocket.
type SecondStage struct {
	Reusable       bool    `json:"reusable"`
	Engines        int     `json:"engines"`
	FuelAmountTons float64 `json:"fuel_amount_tons"`
	BurnTimeSec    *int    `json:"burn_time_sec"`
	Thrust         *Thrust `json:"thrust"`
	Payloads       *struct {
		Option1          string     `json:"option_1"`
		CompositeFairing *Dimension `json:"composite_fairing"`
	} `json:"payloads"`
}

// Engines represents the rocket engines.
type Engines struct {
	Number  int     `json:"number"`
	Type    string  `json:"type"`
	Version string  `json:"version"`
	Layout  *string `json:"layout"`
	ISP     *struct {
		SeaLevel int `json:"sea_level"`
		Vacuum   int `json:"vacuum"`
	} `json:"isp"`
	EngineLossMax  *int     `json:"engine_loss_max"`
	Propellant1    string   `json:"propellant_1"`
	Propellant2    string   `json:"propellant_2"`
	ThrustSeaLevel *Thrust  `json:"thrust_sea_level"`
	ThrustVacuum   *Thrust  `json:"thrust_vacuum"`
	ThrustToWeight *float64 `json:"thrust_to_weight"`
}

// LandingLegs represents the landing legs of a rocket.
type LandingLegs struct {
	Number   int     `json:"number"`
	Material *string `json:"material"`
}

// RocketQueryResults represents the result of a rocket query.
type RocketQueryResults struct {
	Docs          []*Rocket `json:"docs"`
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

// ListAllRockets lists all rockets.
func (s *RocketsService) ListAllRockets(ctx context.Context) ([]*Rocket, error) {
	u := "rockets"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var rockets []*Rocket
	_, err = s.client.do(ctx, req, &rockets)
	if err != nil {
		return nil, err
	}

	return rockets, nil
}

// GetRocket retrieves a specific rocket.
func (s *RocketsService) GetRocket(ctx context.Context, id string) (*Rocket, error) {
	u := fmt.Sprintf("rockets/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	rocket := new(Rocket)
	_, err = s.client.do(ctx, req, rocket)
	if err != nil {
		return nil, err
	}

	return rocket, nil
}

// QueryRockets queries for rockets.
func (s *RocketsService) QueryRockets(ctx context.Context, query map[string]interface{}) (*RocketQueryResults, error) {
	u := "rockets/query"
	req, err := s.client.newRequest(ctx, "POST", u, query)
	if err != nil {
		return nil, err
	}

	results := new(RocketQueryResults)
	_, err = s.client.do(ctx, req, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
