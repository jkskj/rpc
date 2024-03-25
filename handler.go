package main

import (
	"context"
	"log"
)

// TestImpl implements the last service interface defined in the IDL.
type TestImpl struct{}

// Test implements the TestImpl interface.
func (s *TestImpl) Test(ctx context.Context, req string) (resp string, err error) {
	// TODO: Your code here...
	log.Println(req)
	resp = "pong"
	return
}
