package expense

import (
	"encoding/json"
	"errors"
	"github.com/lib/pq"
	"github.com/mrbryside/assessment/internal/pkg/util"
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

func (s *spyStoreWithGetExpenseSuccess) FindOne(rowId int, queryLang string, args ...any) error {
	s.wasCalled = true
	var model modelExpense
	err := json.Unmarshal([]byte(getterMock().GetExpenseSuccess().Response), &model)
	if err != nil {
		return err
	}
	id, _ := args[0].(*int)
	*id = model.ID

	title, _ := args[1].(*string)
	*title = model.Title

	amount, _ := args[2].(*int)
	*amount = model.Amount

	note, _ := args[3].(*string)
	*note = model.Note

	tags, _ := args[4].(*pq.StringArray)
	*tags = model.Tags

	return nil
}

func (s *spyStoreWithGetExpenseSuccess) Insert(modelId interface{}, args ...any) error {
	return nil
}

func (s *spyStoreWithGetExpenseSuccess) IsWasCalled() bool {
	return s.wasCalled
}

// ----------------------------- get failed spy ---------------------------------//
type spyStoreWithGetExpenseFail struct {
	wasCalled bool
}

func newSpyStoreWithGetExpenseFail() *spyStoreWithGetExpenseFail {
	return &spyStoreWithGetExpenseFail{wasCalled: false}
}

func (s *spyStoreWithGetExpenseFail) InitStore() error {
	return nil
}

func (s *spyStoreWithGetExpenseFail) FindOne(rowId int, queryLang string, args ...any) error {
	s.wasCalled = true
	return errors.New("error")
}

func (s *spyStoreWithGetExpenseFail) Insert(modelId interface{}, args ...any) error {
	return nil
}

func (s *spyStoreWithGetExpenseFail) IsWasCalled() bool {
	return s.wasCalled
}

// ----------------------------- get not found spy ---------------------------------//
type spyStoreWithGetExpenseNotFound struct {
	wasCalled bool
}

func newSpyStoreWithGetExpenseNotFound() *spyStoreWithGetExpenseNotFound {
	return &spyStoreWithGetExpenseNotFound{wasCalled: false}
}

func (s *spyStoreWithGetExpenseNotFound) InitStore() error {
	return nil
}

func (s *spyStoreWithGetExpenseNotFound) FindOne(rowId int, queryLang string, args ...any) error {
	s.wasCalled = true
	return util.Error().DBNotFound
}

func (s *spyStoreWithGetExpenseNotFound) Insert(modelId interface{}, args ...any) error {
	return nil
}

func (s *spyStoreWithGetExpenseNotFound) IsWasCalled() bool {
	return s.wasCalled
}
