package wallet

type Wallet struct {
	balance int
}

func (c *Wallet) Balance() int {
	return c.balance
}

func (c *Wallet) Deposit(amount int) {
	c.balance += amount
}

func (c *Wallet) Withdraw(amount int) {
	c.balance -= amount
}
