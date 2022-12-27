package mock

import (
	"errors"
	"github.com/mrbryside/assessment/internal/pkg/db"
)

type spyStore interface {
	db.Store
	IsWasCalled() bool // additional function
}

// ----------------------------- create fail spy ---------------------------------//
type spyStoreWithCreateFail struct {
	wasCalled bool
}

func newSpyStoreWithCreateFail() *spyStoreWithCreateFail {
	return &spyStoreWithCreateFail{wasCalled: false}
}

func (s *spyStoreWithCreateFail) InitStore() error {
	return nil
}

func (s *spyStoreWithCreateFail) Insert(modelId interface{}, args ...any) error {
	s.wasCalled = true
	// destructuring args
	return errors.New("can't insert")
}

func (s *spyStoreWithCreateFail) FindOne(rowId int, model interface{}, queryLang string) error {
	return nil
}

func (s *spyStoreWithCreateFail) IsWasCalled() bool {
	return s.wasCalled
}

// ----------------------------- create success spy ---------------------------------//
type spyStoreWithCreateSuccess struct {
	wasCalled bool
}

func newSpyStoreWithCreateSuccess() *spyStoreWithCreateSuccess {
	return &spyStoreWithCreateSuccess{wasCalled: false}
}

func (s *spyStoreWithCreateSuccess) InitStore() error {
	return nil
}

func (s *spyStoreWithCreateSuccess) Insert(modelId interface{}, args ...any) error {
	s.wasCalled = true
	p, _ := modelId.(*int)
	*p = 5
	return nil
}

func (s *spyStoreWithCreateSuccess) FindOne(rowId int, model interface{}, queryLang string) error {
	return nil
}

func (s *spyStoreWithCreateSuccess) IsWasCalled() bool {
	return s.wasCalled
}
