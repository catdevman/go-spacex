package spacex

import "context"

// CompanyService handles communication with the company related
// methods of the SpaceX API.
type CompanyService struct {
	client *Client
}

// Company represents SpaceX company information.
type Company struct {
	Name          string        `json:"name"`
	Founder       string        `json:"founder"`
	Founded       int           `json:"founded"`
	Employees     int           `json:"employees"`
	Vehicles      int           `json:"vehicles"`
	LaunchSites   int           `json:"launch_sites"`
	TestSites     int           `json:"test_sites"`
	CEO           string        `json:"ceo"`
	CTO           string        `json:"cto"`
	COO           string        `json:"coo"`
	CTOPropulsion string        `json:"cto_propulsion"`
	Valuation     int64         `json:"valuation"`
	Headquarters  *Headquarters `json:"headquarters"`
	Links         *Links        `json:"links"`
	Summary       string        `json:"summary"`
}

// Headquarters represents the company headquarters.
type Headquarters struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

// Links represents company links.
type Links struct {
	Website     string `json:"website"`
	Flickr      string `json:"flickr"`
	Twitter     string `json:"twitter"`
	ElonTwitter string `json:"elon_twitter"`
}

// GetCompanyInfo retrieves company information.
func (s *CompanyService) GetCompanyInfo(ctx context.Context) (*Company, error) {
	u := "company"
	req, err := s.client.newRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	company := new(Company)
	_, err = s.client.do(ctx, req, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}
