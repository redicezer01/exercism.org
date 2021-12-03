package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	closed  bool
	sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount}
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	return int64(a.balance), true
}

func (a *Account) Deposit(amount int64) (balance int64, ok bool) {
	a.Lock()
	balance, ok = a.Balance()
	if !ok {
		return
	}
	if balance+amount < 0 {
		ok = false
	} else {
		balance += amount
		a.balance = balance
	}
	a.Unlock()
	return
}

func (a *Account) Close() (balance int64, ok bool) {
	a.Lock()
	balance, ok = a.Balance()
	if ok {
		a.closed = true
	}
	a.Unlock()
	return
}
