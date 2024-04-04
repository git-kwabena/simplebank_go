package db

import (
	"context"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	// mock an account
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  20,
		Currency: "USD",
	}

	// Pass the mocked account into Create Account func to create account
	account, err := testQueries.CreateAccount(context.Background(), arg)
	//Handle account creation error
	if err != nil {
		t.Fatal(err)
	}
	// Check if account was created else handle error
	if account == nil {
		t.Fatal("account is nil")
	}

	//Check if the entries were entered i.e. Rows were created?
	rows, err := account.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	if rows != 1 {
		t.Errorf("expected 1 row affected, got %d", rows)
	}
	// Get the last inserted rows id to check if account details are correct
	lastInsertedId, err := account.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	// Check if the account details are correct
	// initially createdAt was returning an error so added
	//'parseTime=true' to connection string
	// using go-sql-driver/mysql
	//this asks the driver to scan DATE and DATETIME automatically to time.Time
	createdAccount, err := testQueries.GetAccount(context.Background(), lastInsertedId)
	if err != nil {
		t.Fatal(err)
	}
	if createdAccount.Owner != arg.Owner {
		t.Errorf("expected owner %s, got %s", arg.Owner, createdAccount.Owner)
	}
	if createdAccount.Balance != arg.Balance {
		t.Errorf("expected balance %d, got %d", arg.Balance, createdAccount.Balance)
	}
	if createdAccount.Currency != arg.Currency {
		t.Errorf("expected currency %s, got %s", arg.Currency, createdAccount.Currency)
	}
}
