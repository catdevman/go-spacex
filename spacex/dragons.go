package spacex

import (
	"context"
	"fmt"
)

// DragonsService handles communication with the dragon related
// methods of the SpaceX API.
type DragonsService struct {
	client *Client
}

// Dragon represents a SpaceX dragon.
type Dragon struct {
	Name               string              `json:"name"`
	Type               string              `json:"type"`
	Active             bool                `json:"active"`
	CrewCapacity       int                 `json:"crew_capacity"`
	SidewallAngleDeg   int                 `json:"sidewall_angle_deg"`
	OrbitDurationYr    int                 `json:"orbit_duration_yr"`
	DryMassKg          int                 `json:"dry_mass_kg"`
	DryMassLb          int                 `json:"dry_mass_lb"`
	FirstFlight        *string             `json:"first_flight"`
	HeatShield         *HeatShield         `json:"heat_shield"`
	Thrusters          []Thruster          `json:"thrusters"`
	LaunchPayloadMass  *PayloadMass        `json:"launch_payload_mass"`
	LaunchPayloadVol   *PayloadVolume      `json:"launch_payload_vol"`
	ReturnPayloadMass  *PayloadMass        `json:"return_payload_mass"`
	ReturnPayloadVol   *PayloadVolume      `json:"return_payload_vol"`
	PressurizedCapsule *PressurizedCapsule `json:"pressurized_capsule"`
	Trunk              *Trunk              `json:"trunk"`
	HeightWTrunk       *Dimension          `json:"height_w_trunk"`
	Diameter           *Dimension          `json:"diameter"`
	FlickrImages       []string            `json:"flickr_images"`
	Wikipedia          string              `json:"wikipedia"`
	Description        string              `json:"description"`
	ID                 string              `json:"id"`
}

// HeatShield represents the heat shield of a dragon.
type HeatShield struct {
	Material    string  `json:"material"`
	SizeMeters  float64 `json:"size_meters"`
	TempDegrees int     `json:"temp_degrees"`
	DevPartner  string  `json:"dev_partner"`
}

// Thruster represents a thruster on a dragon capsule.
type Thruster struct {
	Type   string  `json:"type"`
	Amount int     `json:"amount"`
	Pods   int     `json:"pods"`
	Fuel1  string  `json:"fuel_1"`
	Fuel2  string  `json:"fuel_2"`
	ISP    int     `json:"isp"`
	Thrust *Thrust `json:"thrust"`
}

// PayloadMass represents the payload mass.
type PayloadMass struct {
	Kg int `json:"kg"`
	Lb int `json:"lb"`
}

// PayloadVolume represents the payload volume.
type PayloadVolume struct {
	CubicMeters int `json:"cubic_meters"`
	CubicFeet   int `json:"cubic_feet"`
}

// PressurizedCapsule represents the pressurized capsule of a dragon.
type PressurizedCapsule struct {
	PayloadVolume *PayloadVolume `json:"payload_volume"`
}

// Trunk represents the trunk of a dragon.
type Trunk struct {
	TrunkVolume *PayloadVolume `json:"trunk_volume"`
	Cargo       *struct {
		SolarArray         int  `json:"solar_array"`
		UnpressurizedCargo bool `json:"unpressurized_cargo"`
	} `json:"cargo"`
}

// DragonQueryResults represents the result of a dragon query.
type DragonQueryResults struct {
	Docs          []*Dragon `json:"docs"`
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

// ListAllDragons lists all dragons.
func (s *DragonsService) ListAllDragons(ctx context.Context) ([]*Dragon, error) {
	u := "dragons"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var dragons []*Dragon
	_, err = s.client.do(ctx, req, &dragons)
	if err != nil {
		return nil, err
	}

	return dragons, nil
}

// GetDragon retrieves a specific dragon.
func (s *DragonsService) GetDragon(ctx context.Context, id string) (*Dragon, error) {
	u := fmt.Sprintf("dragons/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	dragon := new(Dragon)
	_, err = s.client.do(ctx, req, dragon)
	if err != nil {
		return nil, err
	}

	return dragon, nil
}

// QueryDragons queries for dragons.
func (s *DragonsService) QueryDragons(ctx context.Context, query map[string]interface{}) (*DragonQueryResults, error) {
	u := "dragons/query"
	req, err := s.client.newRequest(ctx, "POST", u, query)
	if err != nil {
		return nil, err
	}

	results := new(DragonQueryResults)
	_, err = s.client.do(ctx, req, results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
