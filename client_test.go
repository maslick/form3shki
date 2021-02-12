package form3shki

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var client, _ = NewClient()

// Create a new client.
func TestNewClient(t *testing.T) {
	clt, err := NewClient()
	assert.NotNil(t, clt)
	assert.Nil(t, err)
}

// Create Form3APIClient with invalid Config (server URL).
func TestNewClientWithInvalidConfig(t *testing.T) {
	config := NewConfig()
	config.SetBaseURL("https://helloworld")
	assert.Equal(t, "https://helloworld", config.BaseURL())

	clt, err := NewClientWithConfig(config)
	assert.NotNil(t, clt)
	assert.NotNil(t, err)
	assert.Equal(t, "server https://helloworld not found", err.Error())

	config = NewConfig()
	config.SetBaseURL("http://accountapi:8080/v1/organisation/accounts")
	clt, err = NewClientWithConfig(config)
	assert.NotNil(t, clt)
	assert.NotNil(t, err)
	assert.Equal(t, "server http://accountapi:8080/v1/organisation/accounts not found", err.Error())
}

// Create Form3APIClient with valid Config (server URL).
func TestNewClientWithValidConfig(t *testing.T) {
	config := NewConfig()
	config.SetBaseURL("http://accountapi:8080")
	clt, err := NewClientWithConfig(config)
	assert.NotNil(t, clt)
	assert.Nil(t, err)
}

// Create a new Account and assert its fields.
func TestCreateAccount(t *testing.T) {
	acc := randomAccount()
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

// Test Creation of a new Account
func TestFetchAccount(t *testing.T) {
	acc := randomAccount()
	_, _ = client.Create(*acc)

	result, err := client.Fetch(acc.ID)
	assert.Nil(t, err)
	assert.Equal(t, acc.ID, result.ID)

	_ = client.Delete(acc.ID, 0)
}

// Test Creation of a new Account with invalid id format (not UUID)
func TestFetchInvalidAccountUUID(t *testing.T) {
	acc := randomAccount()
	_, _ = client.Create(*acc)

	_, err := client.Fetch("bad id")
	assert.NotNil(t, err)
	assert.Equal(t, `{"error_message":"id is not a valid uuid"}`, err.Error())
	_ = client.Delete(acc.ID, 0)
}

// Test Creation of a new Account with invalid id (UUID)
func TestFetchNonExistentAccount(t *testing.T) {
	acc := randomAccount()
	_, _ = client.Create(*acc)

	_, err := client.Fetch("00000000-1111-2222-3333-444444555555")
	assert.NotNil(t, err)
	assert.Equal(t, `{"error_message":"record 00000000-1111-2222-3333-444444555555 does not exist"}`, err.Error())
	_ = client.Delete(acc.ID, 0)
}

// Test List functionality.
func TestListAccounts(t *testing.T) {
	// Create 3 accounts
	acc1, _ := client.Create(*randomAccount())
	acc2, _ := client.Create(*randomAccount())
	acc3, _ := client.Create(*randomAccount())

	list, err := client.List(0, 3)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(list))

	// Clean
	_ = client.Delete(acc1.ID, 0)
	_ = client.Delete(acc2.ID, 0)
	_ = client.Delete(acc3.ID, 0)
}

// Test List functionality (with pagination).
func TestListAccountsWithPagination(t *testing.T) {
	// Create 10 accounts
	acc0, _ := client.Create(*randomAccount())
	acc1, _ := client.Create(*randomAccount())
	acc2, _ := client.Create(*randomAccount())
	acc3, _ := client.Create(*randomAccount())
	acc4, _ := client.Create(*randomAccount())
	acc5, _ := client.Create(*randomAccount())
	acc6, _ := client.Create(*randomAccount())
	acc7, _ := client.Create(*randomAccount())
	acc8, _ := client.Create(*randomAccount())
	acc9, _ := client.Create(*randomAccount())

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

// Create and delete Account.
func TestDeleteAccount(t *testing.T) {
	acc := randomAccount()
	_, _ = client.Create(*acc)

	err := client.Delete(acc.ID, 0)
	assert.Nil(t, err)
}

// Create Account and delete it giving a different version. Delete Account in the end.
func TestDeleteAccountWithInvalidVersion(t *testing.T) {
	acc := randomAccount()
	_, _ = client.Create(*acc)
	err := client.Delete(acc.ID, 1)
	assert.NotNil(t, err)
	assert.Equal(t, `{"error_message":"invalid version"}`, err.Error())
	_ = client.Delete(acc.ID, 0)
}

// Create Account and delete it giving an invalid Account id.
func TestDeleteInvalidAccountId(t *testing.T) {
	acc := randomAccount()
	_, _ = client.Create(*acc)
	err := client.Delete("123-456", 0)
	assert.NotNil(t, err)
	assert.Equal(t, `{"error_message":"id is not a valid uuid"}`, err.Error())

	_ = client.Delete(acc.ID, 0)
}

func randomAccount() *Account {
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
