// Package challenge7 contains the solution for Challenge 7: Bank Account with Error Handling.
package challenge7

import (
	"fmt"
	"sync"
	// Add any other necessary imports
)

// BankAccount represents a bank account with balance management and minimum balance requirements.
type BankAccount struct {
	ID         string
	Owner      string
	Balance    float64
	MinBalance float64
	mu         sync.Mutex // For thread safety
}

// Constants for account operations
const (
	MaxTransactionAmount = 10000.0 // Example limit for deposits/withdrawals
)

// Custom error types

// AccountError is a general error type for bank account operations.
type AccountError struct {
	msg string
	// Implement this error type
}

func (e *AccountError) Error() string {
	// Implement error message
	return e.msg
}

// InsufficientFundsError occurs when a withdrawal or transfer would bring the balance below minimum.
type InsufficientFundsError struct {
	id         string
	owner      string
	minBalance float64
	balance    float64
	amount     float64
	amountName string
	// Implement this error type
}

func (e *InsufficientFundsError) Error() string {
	// Implement error message
	return fmt.Sprintf("Operation %s with amount: %f bring balance: %f below min: %f, acc Id: %s, owner: %s",
		e.amountName, e.amount, e.balance, e.minBalance, e.id, e.owner)
}

// NegativeAmountError occurs when an amount for deposit, withdrawal, or transfer is negative.
type NegativeAmountError struct {
	id         string
	owner      string
	amount     float64
	amountName string
	// Implement this error type
}

func (e *NegativeAmountError) Error() string {
	// Implement error message
	return fmt.Sprintf("Negative %s amount: %f for acc ID: %s, Owner: %s\n", e.amountName, e.amount, e.id, e.owner)
}

// ExceedsLimitError occurs when a deposit or withdrawal amount exceeds the defined limit.
type ExceedsLimitError struct {
	limit      float64
	amount     float64
	amountName string
	id         string
	owner      string
	// Implement this error type
}

func (e *ExceedsLimitError) Error() string {
	// Implement error message
	return fmt.Sprintf("Exceeds limit: %f by %s amount: %f for acc Id: %s, Owner: %s",
		e.limit, e.amountName, e.amount, e.id, e.owner)
}

// NewBankAccount creates a new bank account with the given parameters.
// It returns an error if any of the parameters are invalid.
func NewBankAccount(id, owner string, initialBalance, minBalance float64) (*BankAccount, error) {
	if len(id) == 0 {
		return nil, &AccountError{msg: "invalid id"}
	}
	if len(owner) == 0 {
		return nil, &AccountError{msg: "invalid owner"}
	}
	if initialBalance < 0 {
		return nil, &NegativeAmountError{id: id, owner: owner, amountName: "initialBalance", amount: initialBalance}
	}
	if minBalance < 0 {
		return nil, &NegativeAmountError{id: id, owner: owner, amountName: "minBalance", amount: minBalance}
	}
	if initialBalance < minBalance {
		return nil, &InsufficientFundsError{id: id, owner: owner, amountName: "init account", amount: initialBalance,
			balance: initialBalance, minBalance: minBalance}
	}
	account := BankAccount{
		ID:         id,
		Owner:      owner,
		Balance:    initialBalance,
		MinBalance: minBalance,
		mu:         sync.Mutex{}}
	// Implement account creation with validation
	return &account, nil
}

// Deposit adds the specified amount to the account balance.
// It returns an error if the amount is invalid or exceeds the transaction limit.
func (a *BankAccount) Deposit(amount float64) error {
	if amount > MaxTransactionAmount {
		return &ExceedsLimitError{
			id:         a.ID,
			owner:      a.Owner,
			limit:      MaxTransactionAmount,
			amount:     amount,
			amountName: "deposit",
		}
	} else if amount < 0 {
		return &NegativeAmountError{
			id:         a.ID,
			owner:      a.Owner,
			amount:     amount,
			amountName: "deposit",
		}
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	// Implement deposit functionality with proper error handling
	return nil
}

// Withdraw removes the specified amount from the account balance.
// It returns an error if the amount is invalid, exceeds the transaction limit,
// or would bring the balance below the minimum required balance.
func (a *BankAccount) Withdraw(amount float64) error {
	if amount > MaxTransactionAmount {
		return &ExceedsLimitError{
			id:         a.ID,
			owner:      a.Owner,
			limit:      MaxTransactionAmount,
			amount:     amount,
			amountName: "withdraw",
		}
	} else if amount < 0 {
		return &NegativeAmountError{
			id:         a.ID,
			owner:      a.Owner,
			amount:     amount,
			amountName: "withdraw",
		}
	}
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Balance-amount < a.MinBalance {
		return &InsufficientFundsError{
			id:         a.ID,
			owner:      a.Owner,
			minBalance: a.MinBalance,
			balance:    a.Balance,
			amount:     amount,
			amountName: "withdraw",
		}
	}

	a.Balance -= amount
	// Implement withdrawal functionality with proper error handling
	return nil
}

// Transfer moves the specified amount from this account to the target account.
// It returns an error if the amount is invalid, exceeds the transaction limit,
// or would bring the balance below the minimum required balance.
func (a *BankAccount) Transfer(amount float64, target *BankAccount) error {
	if amount > MaxTransactionAmount {
		return &ExceedsLimitError{
			id:         a.ID,
			owner:      a.Owner,
			limit:      MaxTransactionAmount,
			amount:     amount,
			amountName: "transfer",
		}
	} else if amount < 0 {
		return &NegativeAmountError{
			id:         a.ID,
			owner:      a.Owner,
			amount:     amount,
			amountName: "transfer",
		}
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Balance-amount < a.MinBalance {
		return &InsufficientFundsError{
			id:         a.ID,
			owner:      a.Owner,
			minBalance: a.MinBalance,
			balance:    a.Balance,
			amount:     amount,
			amountName: "transfer",
		}
	}

	target.mu.Lock()
	defer target.mu.Unlock()

	a.Balance -= amount
	target.Balance += amount
	// Implement transfer functionality with proper error handling
	return nil
}
