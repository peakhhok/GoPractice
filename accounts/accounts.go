package accounts

//Account struct
type Account struct {
	owner   string
	balance int
}

//NewAccount create
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account

}
func (a Account) Deposit(amount int) {

}
