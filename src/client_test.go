package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client = New()

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
