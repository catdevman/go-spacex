package spacex

import (
	"context"
	"fmt"
)

// StarlinkService handles communication with the starlink related
// methods of the SpaceX API.
type StarlinkService struct {
	client *Client
}

// Starlink represents a SpaceX starlink satellite.
type Starlink struct {
	Version     *string     `json:"version"`
	Launch      *string     `json:"launch"`
	Longitude   *float64    `json:"longitude"`
	Latitude    *float64    `json:"latitude"`
	HeightKm    *float64    `json:"height_km"`
	VelocityKms *float64    `json:"velocity_kms"`
	SpaceTrack  *SpaceTrack `json:"spaceTrack"`
}

// SpaceTrack represents space track data for a starlink satellite.
type SpaceTrack struct {
	CCSDSOMMVERS       *string  `json:"CCSDS_OMM_VERS"`
	COMMENT            *string  `json:"COMMENT"`
	CREATIONDATE       *string  `json:"CREATION_DATE"`
	ORIGINATOR         *string  `json:"ORIGINATOR"`
	OBJECTNAME         *string  `json:"OBJECT_NAME"`
	OBJECTID           *string  `json:"OBJECT_ID"`
	CENTERNAME         *string  `json:"CENTER_NAME"`
	REFFRAME           *string  `json:"REF_FRAME"`
	TIMESYSTEM         *string  `json:"TIME_SYSTEM"`
	MEANELEMENTTHEORY  *string  `json:"MEAN_ELEMENT_THEORY"`
	EPOCH              *string  `json:"EPOCH"`
	MEANMOTION         *float64 `json:"MEAN_MOTION"`
	ECCENTRICITY       *float64 `json:"ECCENTRICITY"`
	INCLINATION        *float64 `json:"INCLINATION"`
	RAOFASCNODE        *float64 `json:"RA_OF_ASC_NODE"`
	ARGOFPERICENTER    *float64 `json:"ARG_OF_PERICENTER"`
	MEANANOMALY        *float64 `json:"MEAN_ANOMALY"`
	EPHEMERISTYPE      *int     `json:"EPHEMERIS_TYPE"`
	CLASSIFICATIONTYPE *string  `json:"CLASSIFICATION_TYPE"`
	NORADCATID         *int     `json:"NORAD_CAT_ID"`
	ELEMENTSETNO       *int     `json:"ELEMENT_SET_NO"`
	REVATEPOCH         *int     `json:"REV_AT_EPOCH"`
	BSTAR              *float64 `json:"BSTAR"`
	MEANMOTIONDOT      *float64 `json:"MEAN_MOTION_DOT"`
	MEANMOTIONDDOT     *float64 `json:"MEAN_MOTION_DDOT"`
	SEMIMAJORAXIS      *float64 `json:"SEMIMAJOR_AXIS"`
	PERIOD             *float64 `json:"PERIOD"`
	APOAPSIS           *float64 `json:"APOAPSIS"`
	PERIAPSIS          *float64 `json:"PERIAPSIS"`
	OBJECTTYPE         *string  `json:"OBJECT_TYPE"`
	RCSSIZE            *string  `json:"RCS_SIZE"`
	COUNTRYCODE        *string  `json:"COUNTRY_CODE"`
	LAUNCHDATE         *string  `json:"LAUNCH_DATE"`
	SITE               *string  `json:"SITE"`
	DECAYDATE          *string  `json:"DECAY_DATE"`
	DECAYED            *int     `json:"DECAYED"`
	FILE               *int     `json:"FILE"`
	GPID               *int     `json:"GP_ID"`
	TLELINE0           *string  `json:"TLE_LINE0"`
	TLELINE1           *string  `json:"TLE_LINE1"`
	TLELINE2           *string  `json:"TLE_LINE2"`
}

// ListAllStarlink lists all starlink satellites.
func (s *StarlinkService) ListAllStarlink(ctx context.Context) ([]*Starlink, error) {
	u := "starlink"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	var starlink []*Starlink
	_, err = s.client.do(ctx, req, &starlink)
	if err != nil {
		return nil, err
	}

	return starlink, nil
}

// GetStarlink retrieves a specific starlink satellite.
func (s *StarlinkService) GetStarlink(ctx context.Context, id string) (*Starlink, error) {
	u := fmt.Sprintf("starlink/%s", id)
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	starlink := new(Starlink)
	_, err = s.client.do(ctx, req, starlink)
	if err != nil {
		return nil, err
	}

	return starlink, nil
}
