package mock

import (
	"errors"
	"github.com/mrbryside/assessment/internal/pkg/db"
)

type spyStore interface {
	db.Store
	CreateWasCalled() bool // additional function
}

// ----------------------------- create fail spy ---------------------------------//
type spyStoreWithCreateFail struct {
	createWasCalled bool
}

func newSpyStoreWithCreateFail() *spyStoreWithCreateFail {
	return &spyStoreWithCreateFail{createWasCalled: false}
}

func (s *spyStoreWithCreateFail) InitStore() error {
	return nil
}

func (s *spyStoreWithCreateFail) Insert(modelId interface{}, args ...any) error {
	s.createWasCalled = true
	// destructuring args
	return errors.New("can't insert")
}

func (s *spyStoreWithCreateFail) CreateWasCalled() bool {
	return s.createWasCalled
}

// ----------------------------- create success spy ---------------------------------//
type spyStoreWithCreateSuccess struct {
	createWasCalled bool
}

func newSpyStoreWithCreateSuccess() *spyStoreWithCreateSuccess {
	return &spyStoreWithCreateSuccess{createWasCalled: false}
}

func (s *spyStoreWithCreateSuccess) InitStore() error {
	return nil
}

func (s *spyStoreWithCreateSuccess) Insert(modelId interface{}, args ...any) error {
	s.createWasCalled = true
	p, _ := modelId.(*int)
	*p = 5
	return nil
}

func (s *spyStoreWithCreateSuccess) CreateWasCalled() bool {
	return s.createWasCalled
}
