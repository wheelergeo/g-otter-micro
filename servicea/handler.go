package main

import (
	"context"
	servicea "github.com/wheelergeo/g-otter-gen/servicea"
)

// MyServiceImpl implements the last service interface defined in the IDL.
type MyServiceImpl struct{}

// Hello implements the MyServiceImpl interface.
func (s *MyServiceImpl) Hello(ctx context.Context, req *servicea.MyReq) (resp string, err error) {
	// TODO: Your code here...
	return
}
