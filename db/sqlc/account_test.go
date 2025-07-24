package sqlc

//func createRandomAccount(t *testing.T) Account {
//	user := createRandomUser(t)
//	args := CreateAccountParams{
//		Owner:    user.Username,
//		Balance:  utils.RandBalance(),
//		Currency: utils.RandCurrency(),
//	}
//
//	account, err := testQueries.CreateAccount(context.Background(), args)
//	require.NoError(t, err)
//	require.NotEmpty(t, account)
//
//	require.Equal(t, args.Owner, account.Owner)
//	require.Equal(t, args.Balance, account.Balance)
//	require.Equal(t, args.Currency, account.Currency)
//
//	require.NotZero(t, account.ID)
//	require.NotZero(t, account.CreatedAt)
//
//	return account
//}
//
//func TestQueries_CreateAccount(t *testing.T) {
//	createRandomAccount(t)
//}

//func TestQueries_GetAccount(t *testing.T) {
//	account1 := createRandomAccount(t)
//	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
//	require.NoError(t, err)
//	require.NotEmpty(t, account2)
//
//	require.Equal(t, account2.ID, account1.ID)
//	require.Equal(t, account2.Owner, account1.Owner)
//	require.Equal(t, account2.Balance, account1.Balance)
//	require.Equal(t, account2.Currency, account1.Currency)
//}
//
//func TestQueries_UpdateAccount(t *testing.T) {
//	account1 := createRandomAccount(t)
//
//	arg := UpdateAccountParams{
//		ID:      account1.ID,
//		Balance: utils.RandBalance(),
//	}
//
//	account2, err := testQueries.UpdateAccount(context.Background(), arg)
//	require.NoError(t, err)
//	require.NotEmpty(t, account2)
//
//	require.Equal(t, account2.ID, account1.ID)
//	require.Equal(t, account2.Owner, account1.Owner)
//	require.Equal(t, arg.Balance, account2.Balance)
//	require.Equal(t, account2.Currency, account1.Currency)
//}
//
//func TestQueries_DeleteAccount(t *testing.T) {
//	account1 := createRandomAccount(t)
//
//	err := testQueries.DeleteAccount(context.Background(), account1.ID)
//	require.NoError(t, err)
//
//	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
//	require.Error(t, err)
//	require.EqualError(t, err, sql.ErrNoRows.Error())
//	require.Empty(t, account2)
//
//}
