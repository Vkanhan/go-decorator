package main

import (
	"errors"
	"fmt"
	"time"
)

func withRetry(fn func() error, retries int) func() error {
	return func() error {
		var err error
		for range retries {
			err = fn()
			if err == nil {
				return nil
			}
			time.Sleep(time.Second)
		}
		return err
	}
}

func callAPI() error {
	return errors.New("temporary failure")
}

func main() {
	safeCall := withRetry(callAPI, 3)
	err := safeCall()
	if err != nil {
		fmt.Println("API failed after retry", err)
	}
}
