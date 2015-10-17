package account

import "sync"

type Account struct {
	balance int64
	closed  bool
	sync.Mutex
}

func Open(initial int64) *Account {
	if initial < 0 {
		return nil
	}
	return &Account{balance: initial, closed: false}
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	} else {
		a.closed = true
		return a.balance, true
	}
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return -1, false
	} else {
		return a.balance, true
	}
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.closed {
		if a.balance+amount < 0 {
			return 0, false
		}
		a.balance += amount
		return a.balance, true
	} else {
		return 0, false
	}
}
