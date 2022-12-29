package expense

import (
	"github.com/mrbryside/assessment/internal/pkg/db"
)

type spyStore interface {
	db.Store
	IsWasCalled() bool // additional function
}
