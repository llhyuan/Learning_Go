package tracker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// A custom type to be used as key
type guidKey int

// the specific private key(1) of type guidKey will be used to store guid in the context for each request.
var key guidKey = 1

// Add guid info into a given context
func contextWithGUID(ctx context.Context, guid string) context.Context {
	return context.WithValue(ctx, key, guid)
}

// extracting the guid info stored in the context
func guidFromContext(ctx context.Context) (string, bool) {
	guid, ok := ctx.Value(key).(string)
	return guid, ok
}

// Transform a handler into our own handler, adding required functionality
// The new handle function:
//
//	1,takes in a new request
//	2, create a new context from the Request
//	3, Reads in the user's GUID, and add it to the context
//	   * if no GUID is found, generate a new GUID for the user, and put it in the context
//	4, Create a new request from the EDITED context
//	5, Fulfil the request with the NEW request.
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		if guid := req.Header.Get("X-GUID"); guid != "" {
			ctx = contextWithGUID(ctx, guid)
		} else {
			ctx = contextWithGUID(ctx, uuid.New().String())
		}

		req = req.WithContext(ctx)
		h.ServeHTTP(rw, req)
	})
}

type Logger struct{}

func (Logger) Log(ctx context.Context, message string) {
	if guid, ok := guidFromContext(ctx); ok {
		message = fmt.Sprintf("GUID: %s - %s", guid, message)
	}
	fmt.Println(message)
}

// Retreive the guid stored in the context
// Put the guid into the Request's header
func Request(req *http.Request) *http.Request {
	ctx := req.Context()
	if guid, ok := guidFromContext(ctx); ok {
		req.Header.Add("X-GUID", guid)
	}
	return req
}
