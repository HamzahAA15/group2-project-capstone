package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"sirclo/project-capstone/entities/userEntities"

	jwt "github.com/golang-jwt/jwt"
)

var ctxKey = &contextKey{"user"}

type contextKey struct {
	data string
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		if !strings.Contains(header, "Bearer") {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(header, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//validate jwt token
		token, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusForbidden)
			return
		}

		userID := payload["id"].(string)

		// create user and check if user exists in db
		user := userEntities.User{ID: userID}
		// var userServicer userService.UserServiceInterface
		// user, err := userServicer.GetUser(userID)
		// if err != nil {
		// 	next.ServeHTTP(w, r)
		// 	return
		// }
		// user.ID = userID
		// put it in context
		ctx := context.WithValue(r.Context(), ctxKey, &user)
		fmt.Println("ini ctx", ctx)

		// and call the next with our new context
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})

}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *userEntities.User {
	raw, _ := ctx.Value(ctxKey).(*userEntities.User)
	return raw
}
