package spacex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.spacexdata.com/v4/"
	userAgent      = "go-spacex-api-client"
)

// Client for the SpaceX API.
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	Capsules   *CapsulesService
	Company    *CompanyService
	Cores      *CoresService
	Crew       *CrewService
	Dragons    *DragonsService
	History    *HistoryService
	Landpads   *LandpadsService
	Launches   *LaunchesService
	Launchpads *LaunchpadsService
	Payloads   *PayloadsService
	Roadster   *RoadsterService
	Rockets    *RocketsService
	Ships      *ShipsService
	Starlink   *StarlinkService
}

// NewClient returns a new SpaceX API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		client:    httpClient,
		UserAgent: userAgent,
	}
	if c.BaseURL == nil {
		c.BaseURL, _ = url.Parse(defaultBaseURL)
	}

	c.Capsules = &CapsulesService{client: c}
	c.Company = &CompanyService{client: c}
	c.Cores = &CoresService{client: c}
	c.Crew = &CrewService{client: c}
	c.Dragons = &DragonsService{client: c}
	c.History = &HistoryService{client: c}
	c.Landpads = &LandpadsService{client: c}
	c.Launches = &LaunchesService{client: c}
	c.Launchpads = &LaunchpadsService{client: c}
	c.Payloads = &PayloadsService{client: c}
	c.Roadster = &RoadsterService{client: c}
	c.Rockets = &RocketsService{client: c}
	c.Ships = &ShipsService{client: c}
	c.Starlink = &StarlinkService{client: c}

	return c
}

func (c *Client) newRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("baseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	if c := resp.StatusCode; 200 <= c && c <= 299 {
		if v != nil {
			if w, ok := v.(io.Writer); ok {
				io.Copy(w, resp.Body)
			} else {
				decErr := json.NewDecoder(resp.Body).Decode(v)
				if decErr == io.EOF {
					decErr = nil // ignore EOF errors caused by empty response body
				}
				if decErr != nil {
					err = decErr
				}
			}
		}
		return resp, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf("failed to read error response body: %w", err)
	}

	return resp, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
}
