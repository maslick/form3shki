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
	list, _ := clt.List(0, 10)
	for _, account := range list {
		fmt.Println(account.ID)
	}
}

func TestNewClientWithConfig(t *testing.T) {
	config := NewConfig()
	config.SetBaseURL("https://helloworld")
	assert.Equal(t, "https://helloworld", config.BaseURL())

	_, err := NewClientWithConfig(config)
	assert.NotNil(t, err)
	assert.Equal(t, "server not found", err.Error())
	fmt.Println(err.Error())

	config = NewConfig()
	config.SetBaseURL("https://google.com/gmail")
	_, err = NewClientWithConfig(config)
	assert.NotNil(t, err)
}

func TestCreateAccount(t *testing.T) {
	acc := testAccount()
	result, err := client.Create(*acc)

	assert.Nil(t, err)
	assert.Equal(t, acc.ID, result.ID)
	assert.Equal(t, acc.OrganisationID, result.OrganisationID)
	assert.Equal(t, acc.Attributes.BankID, result.Attributes.BankID)
	assert.Equal(t, acc.Attributes.BankIDCode, result.Attributes.BankIDCode)
	assert.Equal(t, acc.Attributes.BaseCurrency, result.Attributes.BaseCurrency)
	assert.Equal(t, acc.Attributes.Bic, result.Attributes.Bic)
	assert.Equal(t, acc.Attributes.Country, result.Attributes.Country)

	_ = client.Delete(acc.ID, 0)
}

func TestFetchAccount(t *testing.T) {
	acc := testAccount()
	_, err := client.Create(*acc)

	result, err := client.Fetch(acc.ID)
	assert.Nil(t, err)
	assert.Equal(t, acc.ID, result.ID)

	_, err = client.Fetch("bad id")
	assert.NotNil(t, err)
	assert.Equal(t, `{"error_message":"id is not a valid uuid"}`, err.Error())

	_, err = client.Fetch("00000000-1111-2222-3333-444444555555")
	assert.NotNil(t, err)
	assert.Equal(t, `{"error_message":"record 00000000-1111-2222-3333-444444555555 does not exist"}`, err.Error())

	_ = client.Delete(acc.ID, 0)
}

func TestListAccounts(t *testing.T) {
	// Create 3 accounts
	acc1, _ := client.Create(*testAccount())
	acc2, _ := client.Create(*testAccount())
	acc3, _ := client.Create(*testAccount())

	list, err := client.List(0, 3)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(list))

	// Clean
	_ = client.Delete(acc1.ID, 0)
	_ = client.Delete(acc2.ID, 0)
	_ = client.Delete(acc3.ID, 0)
}

func TestListAccountsWithPagination(t *testing.T) {
	// Create 10 accounts
	acc0, _ := client.Create(*testAccount())
	acc1, _ := client.Create(*testAccount())
	acc2, _ := client.Create(*testAccount())
	acc3, _ := client.Create(*testAccount())
	acc4, _ := client.Create(*testAccount())
	acc5, _ := client.Create(*testAccount())
	acc6, _ := client.Create(*testAccount())
	acc7, _ := client.Create(*testAccount())
	acc8, _ := client.Create(*testAccount())
	acc9, _ := client.Create(*testAccount())

	list, err := client.List(0, 3)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(list))

	list, err = client.List(1, 3)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(list))

	list, err = client.List(1, 5)
	assert.Nil(t, err)
	assert.Equal(t, 5, len(list))

	list, err = client.List(2, 5)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(list))

	// Clean
	_ = client.Delete(acc0.ID, 0)
	_ = client.Delete(acc1.ID, 0)
	_ = client.Delete(acc2.ID, 0)
	_ = client.Delete(acc3.ID, 0)
	_ = client.Delete(acc4.ID, 0)
	_ = client.Delete(acc5.ID, 0)
	_ = client.Delete(acc6.ID, 0)
	_ = client.Delete(acc7.ID, 0)
	_ = client.Delete(acc8.ID, 0)
	_ = client.Delete(acc9.ID, 0)
}

func TestDeleteAccount(t *testing.T) {
	acc := testAccount()
	_, err := client.Create(*acc)

	err = client.Delete(acc.ID, 0)
	assert.Nil(t, err)

	acc = testAccount()
	_, err = client.Create(*acc)
	err = client.Delete(acc.ID, 1)
	assert.NotNil(t, err)
	assert.Equal(t, `{"error_message":"invalid version"}`, err.Error())

	err = client.Delete(acc.ID, 0)
}

func testAccount() *Account {
	return &Account{
		Type:           "accounts",
		ID:             uuid.New().String(),
		OrganisationID: uuid.New().String(),
		Attributes: Attributes{
			BankID:       "123456",
			BankIDCode:   "GBDSC",
			BaseCurrency: "EUR",
			Bic:          "NWBKGB22",
			Country:      "SI",
		},
	}
}
