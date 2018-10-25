package bank_test

import (
	"fmt"
	"github.com/golang-training/ch09/ex01/bank"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		bank.Withdraw(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		bank.Withdraw(50)
		bank.Deposit(250)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdrawal(t *testing.T) {
	b1 := bank.Balance()
	ok := bank.Withdraw(50)
	if !ok {
		t.Errorf("ok = false, want true. balance = %d", bank.Balance())
	}
	expected := b1 - 50
	if b2 := bank.Balance(); b2 != expected {
		t.Errorf("balance = %d, want %d", b2, expected)
	}
}

func TestWithdrawalFailsIfInsufficientFunds(t *testing.T) {
	b1 := bank.Balance()
	ok := bank.Withdraw(b1 + 1)
	b2 := bank.Balance()
	if ok {
		t.Errorf("ok = true, want false. balance = %d", b2)
	}
	if b2 != b1 {
		t.Errorf("balance = %d, want %d", b2, b1)
	}
}