package tutorial

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/require"
	"simpleBank/utils"
	"testing"
)

func createMockAccount(t *testing.T) (Account, error) {
	arg := CreateAccountParams{
		Owner:    utils.RandOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account, err
}

func TestCreateAccount(t *testing.T) {

	account, err := createMockAccount(t)
	if err != nil {
		fmt.Println("Something is wrong please try again")
	}

	fmt.Println("Account created successfully", account)
}

func TestGetAccount(t *testing.T) {
	account, _ := createMockAccount(t)
	account1, err1 := testQueries.GetAccount(context.Background(), account.ID)
	if err1 != nil {
		fmt.Println("Something is wrong please try again")
	}

	require.NotEmpty(t, account1)
	require.Equal(t, account1.ID, account.ID)
	require.Equal(t, account1.Owner, account.Owner)
	require.Equal(t, account1.Balance, account.Balance)
	require.Equal(t, account1.Currency, account.Currency)
	require.Equal(t, account1.CreatedAt, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	account, _ := createMockAccount(t)

	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: 13000,
	}
	_, err := testQueries.UpdateAccount(context.Background(), args)

	if err != nil {
		panic(err)
	}
	account1, _ := testQueries.GetAccount(context.Background(), account.ID)

	require.Equal(t, args.Balance, account1.Balance)
	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, account.Currency, account1.Currency)
	require.Equal(t, account.ID, account1.ID)
	require.NotEmpty(t, account1)
}

func TestDeleteAccount(t *testing.T) {
	account, _ := createMockAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account1, err1 := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err1)
	require.EqualError(t, err1, sql.ErrNoRows.Error())
	require.Empty(t, account1)
}

func TestListAccount(t *testing.T) {

}
