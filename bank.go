package bank

import (
	"encoding/json"
	"errors"
)

type Customer struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Phone   string `json:"phone,omitempty"`
}
type Account struct {
	StatementInterface `json:"-"` //json忽略
	Customer                      //json单层
	Number             int32      `json:"number,omitempty"`
	Balance            float64    `json:"balance,omitempty"`
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")

	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}
	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}
	a.Balance -= amount
	return nil

}
func (a *Account) Statement() string {
	bts, _ := json.Marshal(a)
	return string(bts)
}

func (a *Account) Transfer(amount float64, receiver *Account) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}
	if err := receiver.Deposit(amount); err != nil {
		return err
	}
	return nil
}

//func Hello() string {
//	return "Hey! I'm working!"
//}
