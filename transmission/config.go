package transmission

import (
	"net/http"
)

type config struct {
	Username string
	Password string

	UserAgent string

	HTTPClient *http.Client
}

// Option customizes client behaviour
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithAuth sets username and password for authentication.
func WithAuth(user, pass string) Option {
	return optionFunc(func(c *config) {
		c.Username = user
		c.Password = pass
	})
}

// WithHTTPClient sets HTTP client to use.
func WithHTTPClient(client *http.Client) Option {
	return optionFunc(func(c *config) {
		c.HTTPClient = client
	})
}

// WithUserAgent sets User-Agent value.
func WithUserAgent(ua string) Option {
	return optionFunc(func(c *config) {
		c.UserAgent = ua
	})
}
