package sqlc

//func TestQueries_CreateUser(t *testing.T) {
//	createRandomUser(t)
//}
//
//func createRandomUser(t *testing.T) Users {
//	hashedPassword, err := utils.HashPassword(utils.RandomString(11))
//	require.NoError(t, err)
//	args := CreateUserParams{
//		Username:       utils.RandOwner(),
//		FullName:       utils.RandOwner(),
//		Email:          utils.RandomEmail(),
//		HashedPassword: hashedPassword,
//	}
//
//	user, err := testQueries.CreateUser(context.Background(), args)
//	require.NoError(t, err)
//	require.NotEmpty(t, user)
//
//	require.Equal(t, user.FullName, args.FullName)
//	require.Equal(t, user.Username, args.Username)
//	require.Equal(t, user.Email, args.Email)
//	require.Equal(t, user.HashedPassword, args.HashedPassword)
//
//	require.NotZero(t, user.CreatedAt)
//	require.True(t, user.PasswordChangedAt.IsZero())
//
//	return user
//}
//
//func TestQueries_GetUser(t *testing.T) {
//	user1 := createRandomUser(t)
//	user2, err := testQueries.GetUser(context.Background(), user1.Username)
//	require.NoError(t, err)
//	require.NotEmpty(t, user2)
//
//	require.Equal(t, user1.Email, user2.Email)
//	require.Equal(t, user1.FullName, user2.FullName)
//	require.Equal(t, user1.Username, user2.Username)
//	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
//}
