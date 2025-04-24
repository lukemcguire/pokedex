package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lukemcguire/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient *http.Client
	cache      pokecache.Cache
}

func NewClient(options ...func(*Client)) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		cache: *pokecache.NewCache(5 * time.Minute),
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// Option for customizing httpClient timeout
func WithHTTPTimeout(timeout time.Duration) func(*Client) {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func (c *Client) get(url string) ([]byte, error) {
	// check cache first
	if data, exists := c.cache.Get(url); exists {
		// fmt.Printf("DEBUG: Cache hit for %s\n", url)
		return data, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if res.StatusCode >= 500 {
		return nil, fmt.Errorf("server error: status code %d, body: %s", res.StatusCode, data)
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("client error: status code %d, body: %s", res.StatusCode, data)
	}
	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected error: status code %d, body: %s", res.StatusCode, data)
	}

	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)

	return data, nil
}
