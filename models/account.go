package models

type Account struct {
	Credits  int             `json:"credits"`
	Loans    []PurchasedLoan `json:"loans"`
	Ships    []OwnedShip     `json:"ships"`
	Username string          `json:"username"`
}
