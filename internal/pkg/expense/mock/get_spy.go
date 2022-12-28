package mock

import (
	"encoding/json"
	"errors"
)

// ----------------------------- get success spy ---------------------------------//
type spyStoreWithGetExpenseSuccess struct {
	wasCalled bool
}

func newSpyStoreWithGetExpenseSuccess() *spyStoreWithGetExpenseSuccess {
	return &spyStoreWithGetExpenseSuccess{wasCalled: false}
}

func (s *spyStoreWithGetExpenseSuccess) InitStore() error {
	return nil
}

func (s *spyStoreWithGetExpenseSuccess) FindOne(rowId int, model interface{}, queryLang string) error {
	s.wasCalled = true

	err := json.Unmarshal([]byte(GetterMock().GetExpenseSuccess().Response), model)
	if err != nil {
		// handle error
	}
	return nil
}

func (s *spyStoreWithGetExpenseSuccess) Insert(modelId interface{}, args ...any) error {
	return nil
}

func (s *spyStoreWithGetExpenseSuccess) IsWasCalled() bool {
	return s.wasCalled
}

// ----------------------------- get success spy ---------------------------------//
type spyStoreWithGetExpenseFail struct {
	wasCalled bool
}

func newSpyStoreWithGetExpenseFail() *spyStoreWithGetExpenseFail {
	return &spyStoreWithGetExpenseFail{wasCalled: false}
}

func (s *spyStoreWithGetExpenseFail) InitStore() error {
	return nil
}

func (s *spyStoreWithGetExpenseFail) FindOne(rowId int, model interface{}, queryLang string) error {
	s.wasCalled = true
	return errors.New("error")
}

func (s *spyStoreWithGetExpenseFail) Insert(modelId interface{}, args ...any) error {
	return nil
}

func (s *spyStoreWithGetExpenseFail) IsWasCalled() bool {
	return s.wasCalled
}
