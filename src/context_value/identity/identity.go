package identity

import (
	"context"
	"net/http"
)

type userKey int

const key userKey = 1

func extractUser(req *http.Request) (string, error) {
	userCookie, err := req.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		user, err := extractUser(req)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := req.Context()
		ctx = ContextWithUser(ctx, user)
		req = req.WithContext(ctx)

		h.ServeHTTP(rw, req)
	})
}

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)

}

func UserFromContext(ctx context.Context) (string, bool) {
	// .(string) is a type assertion
	//  it's a programer's promise to the compiler that
	//  the type of the data is the provided one.
	//  Unmatched type assertion will cause  a runtime panic.
	user, ok := ctx.Value(key).(string)
	return user, ok
}
