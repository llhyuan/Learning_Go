package main

import (
	"context"
	"context_value/identity"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The interface of our business logic
type Logic interface {
	businessLogic(ctx context.Context, user string, data string) (string, error)
}

// Implementing a Controller type and its method
type Controller struct {
	Logic Logic
}

// handleRequest takes in a request and a ResponseWriter
// Create and extract the information and data required by the businessLogic
// invoke the businessLogic.
func (c Controller) handleRequest(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user, ok := identity.UserFromContext(ctx)
	if !ok {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := req.URL.Query().Get("data")
	result, err := c.Logic.businessLogic(ctx, user, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.Write([]byte(result))

}

// Dependency injection with interface
type BusinessLogic struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	Remote           string
}

type Logger interface {
	Log(context.Context, string)
}

type RequestDecorator func(*http.Request) *http.Request

func (bl BusinessLogic) businessLogic(ctx context.Context, user string, data string) (string, error) {
	bl.Logger.Log(ctx, "starting businesslogic for "+user+" with "+data)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, bl.Remote+"?query="+data, nil)
	if err != nil {
		bl.Logger.Log(ctx, "error building remote request: "+err.Error())
		return "", err
	}

	// Read the user's guid from context and
	// add it as a request header
	req = bl.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		bl.Logger.Log(ctx, fmt.Sprintf("error makeing request: %v", err))
		return "", err
	}

	body := resp.Body
	defer body.Close()

	out, err := ioutil.ReadAll(body)
	if err != nil {
		bl.Logger.Log(ctx, fmt.Sprintf("error reading response: %v", err))
		return "", err
	}

	return string(out), nil
}
