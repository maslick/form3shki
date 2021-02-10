package form3shki

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type IForm3ApiClient interface {
	Create(account Account) (Account, error)
	Fetch(accountId string) (Account, error)
	List() ([]Account, error)
	Delete(accountId string, version int) error
}

type Form3APIClient struct {
	BaseUrl string
}

// Create a new Form3 account client.
// The default Form3 API URL is http://localhost:8080.
// You can override it by setting API_URL environment variable or use Form3APIClient constructor directly.
func New() IForm3ApiClient {
	url := getEnv("API_URL", "http://localhost:8080")
	log.Println("API_URL:", url)
	return &Form3APIClient{BaseUrl: url}
}

func (c *Form3APIClient) Create(account Account) (Account, error) {
	url := c.BaseUrl + "/v1/organisation/accounts"
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(AccountDTO{Account: account})
	if err != nil {
		return Account{}, err
	}

	resp, err := http.Post(url, "application/vnd.api+json", body)
	if err != nil {
		return Account{}, err
	}
	var respObj AccountDTO
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return Account{}, err
	}

	return respObj.Account, nil
}

func (c *Form3APIClient) Fetch(accountId string) (Account, error) {
	url := c.BaseUrl + "/v1/organisation/accounts/" + accountId
	resp, err := http.Get(url)
	if err != nil {
		return Account{}, err
	}

	if resp.Status != "200 OK" {
		text, _ := ioutil.ReadAll(resp.Body)
		return Account{}, errors.New(string(text))
	}

	var respObj AccountDTO
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return Account{}, err
	}

	return respObj.Account, nil
}

func (c *Form3APIClient) List() ([]Account, error) {
	url := c.BaseUrl + "/v1/organisation/accounts"
	resp, err := http.Get(url)
	if err != nil {
		return []Account{}, err
	}

	var respObj AccountsDTO
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Accounts, nil
}

func (c *Form3APIClient) Delete(accountId string, version int) error {
	url := fmt.Sprintf(`%s/v1/organisation/accounts/%s?version=%d`, c.BaseUrl, accountId, version)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.Status != "204 No Content" {
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
