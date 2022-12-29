package expense

import (
	"github.com/mrbryside/assessment/internal/pkg/db"
	"log"
	"net/http"
	"testing"
)

var updateTests = []struct {
	name     string
	code     int
	spy      db.StoreSpy
	payload  string
	response string
	called   bool
}{
	{
		name:     "should return response expense json",
		code:     http.StatusOK,
		spy:      newSpyGetSuccess(),
		payload:  "5",
		response: getResponse,
		called:   true,
	},
}

func TestUpdateExpense(t *testing.T) {
	t.Parallel()
	for _, ut := range updateTests {
		t.Run(ut.name, func(t *testing.T) {
			log.Println("test!!")

		})
	}
}
