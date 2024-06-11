package view

import (
	"strconv"
)

// AuthenticatedUser check if user is authenticated. Return nil user if not
// func AuthenticatedUser(ctx context.Context) types.AuthenticatedUser {
// 	user, ok := ctx.Value(types.UserContextKey).(types.AuthenticatedUser)
// 	if !ok {
// 		return types.AuthenticatedUser{}
// 	}
// 	return user
// }

func String(i int) string {
	return strconv.Itoa(i)
}
