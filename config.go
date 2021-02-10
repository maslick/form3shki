package form3shki

// Config is a configuration object.
type Config struct {
	url string
}

// NewConfig will create a default Config
func NewConfig() *Config {
	return &Config{}
}

// BaseURL is a getter
func (c *Config) BaseURL() string {
	return c.url
}

// SetBaseURL is a setter
func (c *Config) SetBaseURL(url string) {
	c.url = url
}
