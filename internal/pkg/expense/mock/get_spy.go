package mock

import (
	"encoding/json"
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
