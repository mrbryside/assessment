package db

type StoreSpy interface {
	Store
	IsWasCalled() bool // additional function
}

type storeSpy struct {
	wasCalled bool
	insert    func(...any) error
	findOne   func(...any) error
}

func NewStoreSpy(insert func(...any) error, findOne func(...any) error) *storeSpy {
	return &storeSpy{
		insert:  insert,
		findOne: findOne,
	}
}

func (f *storeSpy) InitStore() error {
	return nil
}

func (f *storeSpy) Script() script {
	return newScript()
}

func (f *storeSpy) Insert(modelId interface{}, args ...any) error {
	f.wasCalled = true
	return f.insert(modelId)
}

func (f *storeSpy) FindOne(rowId int, queryLang string, args ...any) error {
	f.wasCalled = true
	return f.findOne(args...)
}

func (f *storeSpy) IsWasCalled() bool {
	return f.wasCalled
}
