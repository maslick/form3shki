package form3shki

type Account struct {
	Type           string     `json:"type"`
	Id             string     `json:"id"`
	OrganisationId string     `json:"organisation_id"`
	Attributes     Attributes `json:"attributes"`
}

type Attributes struct {
	BankId       string `json:"bank_id"`
	BankIdCode   string `json:"bank_id_code"`
	BaseCurrency string `json:"base_currency"`
	Bic          string `json:"bic"`
	Country      string `json:"country"`
}

type AccountDTO struct {
	Account Account `json:"data"`
	Links   Links   `json:"links"`
}

type AccountsDTO struct {
	Accounts []Account `json:"data"`
	Links    Links     `json:"links"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
	Self  string `json:"self"`
}
