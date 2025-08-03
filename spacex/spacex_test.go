package spacex

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// setup sets up a test HTTP server along with a spacex.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the SpaceX client being tested and is configured to use that server.
	client = NewClient(nil)
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func TestCapsulesService_GetCapsule(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/capsules/5e9e2c5bf35918ed873b2664", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"serial":"C101"}`)
	})

	ctx := context.Background()
	capsule, err := client.Capsules.GetCapsule(ctx, "5e9e2c5bf35918ed873b2664")
	if err != nil {
		t.Fatalf("Capsules.GetCapsule returned error: %v", err)
	}

	want := &Capsule{Serial: "C101"}
	if capsule.Serial != want.Serial {
		t.Errorf("Capsules.GetCapsule returned %+v, want %+v", capsule.Serial, want.Serial)
	}
}

func TestCompanyService_GetCompanyInfo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/company", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"name":"SpaceX"}`)
	})

	ctx := context.Background()
	company, err := client.Company.GetCompanyInfo(ctx)
	if err != nil {
		t.Fatalf("Company.GetCompanyInfo returned error: %v", err)
	}

	want := &Company{Name: "SpaceX"}
	if company.Name != want.Name {
		t.Errorf("Company.GetCompanyInfo returned %+v, want %+v", company.Name, want.Name)
	}
}

func TestCoresService_GetCore(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cores/5e9e28a6f35918c0803b265c", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"serial":"B1051"}`)
	})

	ctx := context.Background()
	core, err := client.Cores.GetCore(ctx, "5e9e28a6f35918c0803b265c")
	if err != nil {
		t.Fatalf("Cores.GetCore returned error: %v", err)
	}

	want := &Core{Serial: "B1051"}
	if core.Serial != want.Serial {
		t.Errorf("Cores.GetCore returned %+v, want %+v", core.Serial, want.Serial)
	}
}

func TestCrewService_GetCrew(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	name := "Robert Behnken"
	mux.HandleFunc("/crew/5ebf1a6e23a9a60006e03a7a", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprintf(w, `{"name":"%s"}`, name)
	})

	ctx := context.Background()
	crew, err := client.Crew.GetCrew(ctx, "5ebf1a6e23a9a60006e03a7a")
	if err != nil {
		t.Fatalf("Crew.GetCrew returned error: %v", err)
	}

	want := &Crew{Name: &name}
	if *crew.Name != *want.Name {
		t.Errorf("Crew.GetCrew returned %+v, want %+v", *crew.Name, *want.Name)
	}
}

func TestDragonsService_GetDragon(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/dragons/5e9d058759b1ff74a7ad5f8f", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"id":"5e9d058759b1ff74a7ad5f8f","name":"Dragon 1"}`)
	})

	ctx := context.Background()
	dragon, err := client.Dragons.GetDragon(ctx, "5e9d058759b1ff74a7ad5f8f")
	if err != nil {
		t.Fatalf("Dragons.GetDragon returned error: %v", err)
	}

	want := &Dragon{ID: "5e9d058759b1ff74a7ad5f8f", Name: "Dragon 1"}
	if dragon.Name != want.Name {
		t.Errorf("Dragons.GetDragon returned %+v, want %+v", dragon.Name, want.Name)
	}
}

func TestHistoryService_GetHistory(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	title := "First successfull Dragon launch"
	mux.HandleFunc("/history/5e9d058759b1ff74a7ad5f8f", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprintf(w, `{"title":"%s"}`, title)
	})

	ctx := context.Background()
	history, err := client.History.GetHistory(ctx, "5e9d058759b1ff74a7ad5f8f")
	if err != nil {
		t.Fatalf("History.GetHistory returned error: %v", err)
	}

	want := &History{Title: &title}
	if *history.Title != *want.Title {
		t.Errorf("History.GetHistory returned %+v, want %+v", *history.Title, *want.Title)
	}
}

func TestLandpadsService_GetLandpad(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	name := "LZ-1"
	mux.HandleFunc("/landpads/5e9e3032383ecb267a34e7c7", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprintf(w, `{"name":"%s"}`, name)
	})

	ctx := context.Background()
	landpad, err := client.Landpads.GetLandpad(ctx, "5e9e3032383ecb267a34e7c7")
	if err != nil {
		t.Fatalf("Landpads.GetLandpad returned error: %v", err)
	}

	want := &Landpad{Name: &name}
	if *landpad.Name != *want.Name {
		t.Errorf("Landpads.GetLandpad returned %+v, want %+v", *landpad.Name, *want.Name)
	}
}

func TestLaunchesService_GetLaunch(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/launches/5eb87cd9ffd86e000604b32a", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"id":"5eb87cd9ffd86e000604b32a","name":"FalconSat"}`)
	})

	ctx := context.Background()
	launch, err := client.Launches.GetLaunch(ctx, "5eb87cd9ffd86e000604b32a")
	if err != nil {
		t.Fatalf("Launches.GetLaunch returned error: %v", err)
	}

	want := &Launch{ID: "5eb87cd9ffd86e000604b32a", Name: "FalconSat"}
	if launch.Name != want.Name {
		t.Errorf("Launches.GetLaunch returned %+v, want %+v", launch.Name, want.Name)
	}
}

func TestLaunchpadsService_GetLaunchpad(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	name := "Kwajalein Atoll"
	mux.HandleFunc("/launchpads/5e9e4502f5090995de566f86", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprintf(w, `{"name":"%s"}`, name)
	})

	ctx := context.Background()
	launchpad, err := client.Launchpads.GetLaunchpad(ctx, "5e9e4502f5090995de566f86")
	if err != nil {
		t.Fatalf("Launchpads.GetLaunchpad returned error: %v", err)
	}

	want := &Launchpad{Name: &name}
	if *launchpad.Name != *want.Name {
		t.Errorf("Launchpads.GetLaunchpad returned %+v, want %+v", *launchpad.Name, *want.Name)
	}
}

func TestPayloadsService_GetPayload(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	name := "Iridium NEXT"
	mux.HandleFunc("/payloads/5eb0e4c6b6c3bb0006eeb21e", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprintf(w, `{"name":"%s"}`, name)
	})

	ctx := context.Background()
	payload, err := client.Payloads.GetPayload(ctx, "5eb0e4c6b6c3bb0006eeb21e")
	if err != nil {
		t.Fatalf("Payloads.GetPayload returned error: %v", err)
	}

	want := &Payload{Name: &name}
	if *payload.Name != *want.Name {
		t.Errorf("Payloads.GetPayload returned %+v, want %+v", *payload.Name, *want.Name)
	}
}

func TestRoadsterService_GetRoadsterInfo(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/roadster", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"name":"Elon Musk's Tesla Roadster"}`)
	})

	ctx := context.Background()
	roadster, err := client.Roadster.GetRoadsterInfo(ctx)
	if err != nil {
		t.Fatalf("Roadster.GetRoadsterInfo returned error: %v", err)
	}

	want := &Roadster{Name: "Elon Musk's Tesla Roadster"}
	if roadster.Name != want.Name {
		t.Errorf("Roadster.GetRoadsterInfo returned %+v, want %+v", roadster.Name, want.Name)
	}
}

func TestRocketsService_GetRocket(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/rockets/5e9d0d95eda69955f709d1eb", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"id":"5e9d0d95eda69955f709d1eb","name":"Falcon 1"}`)
	})

	ctx := context.Background()
	rocket, err := client.Rockets.GetRocket(ctx, "5e9d0d95eda69955f709d1eb")
	if err != nil {
		t.Fatalf("Rockets.GetRocket returned error: %v", err)
	}

	want := &Rocket{ID: "5e9d0d95eda69955f709d1eb", Name: "Falcon 1"}
	if rocket.Name != want.Name {
		t.Errorf("Rockets.GetRocket returned %+v, want %+v", rocket.Name, want.Name)
	}
}

func TestShipsService_GetShip(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/ships/5ea6ed2e080df4000697c90a", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{"name":"GO Pursuit"}`)
	})

	ctx := context.Background()
	ship, err := client.Ships.GetShip(ctx, "5ea6ed2e080df4000697c90a")
	if err != nil {
		t.Fatalf("Ships.GetShip returned error: %v", err)
	}

	want := &Ship{Name: "GO Pursuit"}
	if ship.Name != want.Name {
		t.Errorf("Ships.GetShip returned %+v, want %+v", ship.Name, want.Name)
	}
}

func TestStarlinkService_GetStarlink(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	version := "v1.0"
	mux.HandleFunc("/starlink/5eed770f096e59000698560d", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method = %v, want %v", r.Method, "GET")
		}
		fmt.Fprintf(w, `{"version":"%s"}`, version)
	})

	ctx := context.Background()
	starlink, err := client.Starlink.GetStarlink(ctx, "5eed770f096e59000698560d")
	if err != nil {
		t.Fatalf("Starlink.GetStarlink returned error: %v", err)
	}

	want := &Starlink{Version: &version}
	if *starlink.Version != *want.Version {
		t.Errorf("Starlink.GetStarlink returned %+v, want %+v", *starlink.Version, *want.Version)
	}
}
