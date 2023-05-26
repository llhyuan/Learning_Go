package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Create a default client.
var client = http.Client{}

func callBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
	ctx, cancel := context.WithCancel(ctx)
	// cancel function must be called.
	// it's ok to call it more than once.
	// Every invocation after the first one will be ignored.
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
	fmt.Println("done with both")
}

func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request err: ", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "respond error: ", err)
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(label, "read error: ", err)
		return err
	}

	result := string(data)
	if result != "" {
		fmt.Println(label, "result: ", result)
	}

	if result == "error" {
		fmt.Println("canceling from: ", label)
		return errors.New("error happened")
	}

	return nil
}
