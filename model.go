package form3shki

// Account represents Form3 account.
type Account struct {
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Attributes     Attributes `json:"attributes"`
}

// Attributes represents Account's attributes.
type Attributes struct {
	BankID       string `json:"bank_id"`
	BankIDCode   string `json:"bank_id_code"`
	BaseCurrency string `json:"base_currency"`
	Bic          string `json:"bic"`
	Country      string `json:"country"`
}

type accountDTO struct {
	Account Account `json:"data"`
	Links   links   `json:"links"`
}

type accountsDTO struct {
	Accounts []Account `json:"data"`
	Links    links     `json:"links"`
}

type links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
	Self  string `json:"self"`
}
