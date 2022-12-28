package expense

const (
	called    = true
	notCalled = false
)

type mock struct {
	SpyStore spyStore
	Payload  string
	Response string
	Code     int
	Called   bool
}

func newExpenseMock(s spyStore, p string, r string, c int, ca bool) mock {
	return mock{
		SpyStore: s,
		Payload:  p,
		Response: r,
		Code:     c,
		Called:   ca,
	}
}
