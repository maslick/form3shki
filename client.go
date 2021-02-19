package form3shki

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Client is a Form3API client.
// Form3API client enables you to do all Form3 account CRUD operations.
type Client interface {
	Create(account Account) (Account, error)
	Fetch(accountID string) (Account, error)
	List(page int, size int) ([]Account, error)
	Delete(accountID string, version int) error
}

// Form3APIClient is the default implementation of the Client interface.
type Form3APIClient struct {
	BaseURL string
}

func (c *Form3APIClient) init() error {
	url := c.BaseURL + "/v1/health"
	resp, err := http.Get(url)

	errMsg := fmt.Sprintf("server %s not found", c.BaseURL)
	if err != nil {
		return errors.New(errMsg)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(errMsg)
	}

	text, _ := ioutil.ReadAll(resp.Body)
	if string(text) != `{"status":"up"}` {
		return errors.New(errMsg)
	}

	return nil
}

// NewClient creates a new Form3 account client.
// The default Form3 API URL is http://localhost:8080.
// You can override it by setting API_URL environment variable or use Form3APIClient constructor directly.
func NewClient() (*Form3APIClient, error) {
	url := getEnv("API_URL", "http://localhost:8080")
	client := &Form3APIClient{BaseURL: url}
	err := client.init()
	return client, err
}

// NewClientWithConfig creates a new Form3 account client from configuration Config:
//
// config := NewConfig()
// config.BaseURL("http://hello.world:8080")
// client, _ := form3shki.NewClientWithConfig(config)
func NewClientWithConfig(config *Config) (*Form3APIClient, error) {
	client := &Form3APIClient{BaseURL: config.url}
	err := client.init()
	return client, err
}

// Create will create a new account.
func (c *Form3APIClient) Create(account Account) (Account, error) {
	url := c.BaseURL + "/v1/organisation/accounts"
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(accountDTO{Account: account})
	if err != nil {
		return Account{}, err
	}

	resp, err := http.Post(url, "application/vnd.api+json", body)
	if err != nil {
		return Account{}, err
	}
	var respObj accountDTO
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return Account{}, err
	}

	return respObj.Account, nil
}

// Fetch retrieves account with accountID.
func (c *Form3APIClient) Fetch(accountID string) (Account, error) {
	url := c.BaseURL + "/v1/organisation/accounts/" + accountID
	resp, err := http.Get(url)
	if err != nil {
		return Account{}, err
	}

	if resp.StatusCode != http.StatusOK {
		text, _ := ioutil.ReadAll(resp.Body)
		return Account{}, errors.New(string(text))
	}

	var respObj accountDTO
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return Account{}, err
	}

	return respObj.Account, nil
}

// List retrieves a list of Account's in a paginated way.
// page == is the page number
// size == how many items to fetch
// E.g. client.List(0, 10) - get max. 10 elements from page 0
func (c *Form3APIClient) List(page int, size int) ([]Account, error) {
	url := fmt.Sprintf("%s/v1/organisation/accounts?page[number]=%d&page[size]=%d", c.BaseURL, page, size)
	resp, err := http.Get(url)
	if err != nil {
		return []Account{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var respObj accountsDTO
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Accounts, nil
}

// Delete will remove account with the given account id and version number (defaults to 0)
func (c *Form3APIClient) Delete(accountID string, version int) error {
	url := fmt.Sprintf(`%s/v1/organisation/accounts/%s?version=%d`, c.BaseURL, accountID, version)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		text, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(text))
	}

	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
