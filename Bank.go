package main

import (
	"errors"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Account struct to hold account information
type Account struct {
	AccountNumber string
	Balance       float64
}

// Bank struct to hold multiple accounts
type Bank struct {
	Accounts map[string]*Account
}

// NewBank creates a new Bank instance
func NewBank() *Bank {
	return &Bank{Accounts: make(map[string]*Account)}
}

// CreateAccount creates a new account with a given account number
func (b *Bank) CreateAccount(accountNumber string) (*Account, error) {
	if _, exists := b.Accounts[accountNumber]; exists {
		return nil, errors.New("account already exists")
	}
	account := &Account{AccountNumber: accountNumber, Balance: 0.0}
	b.Accounts[accountNumber] = account
	return account, nil
}

// Deposit adds money to the specified account
func (b *Bank) Deposit(accountNumber string, amount float64) error {
	account, exists := b.Accounts[accountNumber]
	if !exists {
		return errors.New("account not found")
	}
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	account.Balance += amount
	return nil
}

// Withdraw removes money from the specified account
func (b *Bank) Withdraw(accountNumber string, amount float64) error {
	account, exists := b.Accounts[accountNumber]
	if !exists {
		return errors.New("account not found")
	}
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if amount > account.Balance {
		return errors.New("insufficient funds")
	}
	account.Balance -= amount
	return nil
}

// GetBalance returns the balance of the specified account
func (b *Bank) GetBalance(accountNumber string) (float64, error) {
	account, exists := b.Accounts[accountNumber]
	if !exists {
		return 0.0, errors.New("account not found")
	}
	return account.Balance, nil
}

func main() {
	bank := NewBank()
	a := app.New()
	w := a.NewWindow("Banking Application")

	accountNumberEntry := widget.NewEntry()
	accountNumberEntry.SetPlaceHolder("Account Number")

	messageLabel := widget.NewLabel("")

	createAccountButton := widget.NewButton("Create Account", func() {
		accountNumber := accountNumberEntry.Text
		if accountNumber == "" {
			messageLabel.SetText("Please enter an account number.")
			return
		}
		_, err := bank.CreateAccount(accountNumber)
		if err != nil {
			messageLabel.SetText(err.Error())
		} else {
			messageLabel.SetText(fmt.Sprintf("Account %s created.", accountNumber))
		}
	})

	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("Amount")

	depositButton := widget.NewButton("Deposit", func() {
		accountNumber := accountNumberEntry.Text
		amount, err := strconv.ParseFloat(amountEntry.Text, 64)
		if err != nil {
			messageLabel.SetText("Invalid amount.")
			return
		}
		err = bank.Deposit(accountNumber, amount)
		if err != nil {
			messageLabel.SetText(err.Error())
		} else {
			messageLabel.SetText(fmt.Sprintf("Deposited %.2f to account %s.", amount, accountNumber))
		}
	})

	withdrawButton := widget.NewButton("Withdraw", func() {
		accountNumber := accountNumberEntry.Text
		amount, err := strconv.ParseFloat(amountEntry.Text, 64)
		if err != nil {
			messageLabel.SetText("Invalid amount.")
			return
		}
		err = bank.Withdraw(accountNumber, amount)
		if err != nil {
			messageLabel.SetText(err.Error())
		} else {
			messageLabel.SetText(fmt.Sprintf("Withdrew %.2f from account %s.", amount, accountNumber))
		}
	})

	balanceButton := widget.NewButton("Get Balance", func() {
		accountNumber := accountNumberEntry.Text
		balance, err := bank.GetBalance(accountNumber)
		if err != nil {
			messageLabel.SetText(err.Error())
		} else {
			messageLabel.SetText(fmt.Sprintf("Account %s balance: %.2f", accountNumber, balance))
		}
	})

	w.SetContent(container.NewVBox(
		accountNumberEntry,
		createAccountButton,
		amountEntry,
		depositButton,
		withdrawButton,
		balanceButton,
		messageLabel,
	))

	w.ShowAndRun()
}
