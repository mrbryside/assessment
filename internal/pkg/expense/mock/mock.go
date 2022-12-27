package mock

const (
	called    = true
	notCalled = false
)

type expenseCreator struct{}

func CreationMock() expenseCreator {
	return expenseCreator{}
}

type expenseGetter struct{}

func GetterMock() expenseGetter {
	return expenseGetter{}
}

type GetExpenseMock struct {
	SpyStore spyStore
	Payload  string
	Response string
	Code     int
	Called   bool
}

func newGetExpenseMock(s spyStore, p string, r string, c int, ca bool) GetExpenseMock {
	return GetExpenseMock{
		SpyStore: s,
		Payload:  p,
		Response: r,
		Code:     c,
		Called:   ca,
	}
}

type CreateExpenseMock struct {
	SpyStore spyStore
	Payload  string
	Response string
	Code     int
	Called   bool
}

func newCreateExpenseMock(s spyStore, p string, r string, c int, ca bool) CreateExpenseMock {
	return CreateExpenseMock{
		SpyStore: s,
		Payload:  p,
		Response: r,
		Code:     c,
		Called:   ca,
	}
}
