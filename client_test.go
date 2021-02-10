package form3shki

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client, _ = NewClient()

func TestNewClient(t *testing.T) {
	clt, err := NewClient()
	assert.Nil(t, err, "client should be able to initialize")
	list, _ := clt.List()
	for _, account := range list {
		fmt.Println(account.Id)
	}
}

func TestNewClientWithConfig(t *testing.T) {
	config := NewConfig()
	config.SetBaseUrl("https://helloworld")
	assert.Equal(t, "https://helloworld", config.BaseUrl())

	_, err := NewClientWithConfig(config)
	assert.NotNil(t, err, "this should fail")
	assert.Equal(t, "server not found", err.Error())
	fmt.Println(err.Error())

	config = NewConfig()
	config.SetBaseUrl("https://google.com/gmail")
	_, err = NewClientWithConfig(config)
	assert.NotNil(t, err, "this should fail")
}

func TestCreateAccount(t *testing.T) {
	acc := testAccount()
	result, err := client.Create(*acc)

	assert.Nil(t, err, "the request should not return error")
	assert.Equal(t, acc.Id, result.Id)
	assert.Equal(t, acc.OrganisationId, result.OrganisationId)
	assert.Equal(t, acc.Attributes.BankId, result.Attributes.BankId)
	assert.Equal(t, acc.Attributes.BankIdCode, result.Attributes.BankIdCode)
	assert.Equal(t, acc.Attributes.BaseCurrency, result.Attributes.BaseCurrency)
	assert.Equal(t, acc.Attributes.Bic, result.Attributes.Bic)
	assert.Equal(t, acc.Attributes.Country, result.Attributes.Country)
}

func TestFetchAccount(t *testing.T) {
	acc := testAccount()
	_, err := client.Create(*acc)

	result, err := client.Fetch(acc.Id)
	assert.Nil(t, err, "the request should not return error")
	assert.Equal(t, acc.Id, result.Id)

	_, err = client.Fetch("bad id")
	assert.NotNil(t, err, "account id should be in UUID format")
	assert.Equal(t, `{"error_message":"id is not a valid uuid"}`, err.Error())

	_, err = client.Fetch("00000000-1111-2222-3333-444444555555")
	assert.NotNil(t, err, "this should fail")
	assert.Equal(t, `{"error_message":"record 00000000-1111-2222-3333-444444555555 does not exist"}`, err.Error())
}

func TestListAccounts(t *testing.T) {
	list, err := client.List()
	assert.Nil(t, err, "the request should not return error")
	assert.Greater(t, len(list), 1)

	for _, account := range list {
		fmt.Println(account.Id)
	}
}

func TestDeleteAccount(t *testing.T) {
	acc := testAccount()
	_, err := client.Create(*acc)

	err = client.Delete(acc.Id, 0)
	assert.Nil(t, err, "the request should not return error")

	acc = testAccount()
	_, err = client.Create(*acc)
	err = client.Delete(acc.Id, 1)
	assert.NotNil(t, err, "the request should fail")
	assert.Equal(t, `{"error_message":"invalid version"}`, err.Error())
}

func testAccount() *Account {
	return &Account{
		Type:           "accounts",
		Id:             uuid.New().String(),
		OrganisationId: uuid.New().String(),
		Attributes: Attributes{
			BankId:       "123456",
			BankIdCode:   "GBDSC",
			BaseCurrency: "EUR",
			Bic:          "NWBKGB22",
			Country:      "SI",
		},
	}
}
