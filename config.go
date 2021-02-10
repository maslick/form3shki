package form3shki

type Config struct {
	url string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) BaseUrl() string {
	return c.url
}

func (c *Config) SetBaseUrl(url string) {
	c.url = url
}
