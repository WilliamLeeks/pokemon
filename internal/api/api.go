package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	apiURLBase string = "https://pokeapi.co/api/v2/pokemon/%s"
	userAgent  string = "Pokemon/1.0"
)

// Pokemon represents data from an API response.
type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// String implements the Stringer interface.
func (p *Pokemon) String() string {
	return p.Name
}

// GetPokemon ...
func GetPokemon(num int) *Pokemon {
	URL := fmt.Sprintf(apiURLBase, strconv.Itoa(num))

	resp, err := request(URL)
	if err != nil {
		fmt.Printf("HTTP Request failed: %s", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Request Body Read failed: %s", err)
		os.Exit(1)
	}

	p := Pokemon{}
	json.Unmarshal([]byte(body), &p)

	return &p
}

func request(URL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	// Set User Agent header.
	req.Header.Set("User-Agent", userAgent)

	// Get an HTTP client with timeouts.
	client := getDefaultClient()

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	// Check status code is 2XX.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, fmt.Errorf("Invalid HTTP status code: %d", resp.StatusCode)
	}

	return resp, nil
}
