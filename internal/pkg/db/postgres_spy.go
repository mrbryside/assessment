package db

type StoreSpy interface {
	Store
	IsWasCalled() bool // additional function
}

type storeSpy struct {
	wasCalled bool
	insert    func(...interface{}) error
	findOne   func(...interface{}) error
	update    func(...interface{}) error
}

func NewStoreSpy(
	insert func(...interface{}) error,
	findOne func(...interface{}) error,
	update func(...interface{}) error,
) *storeSpy {
	return &storeSpy{
		insert:  insert,
		findOne: findOne,
		update:  update,
	}
}

func (f *storeSpy) InitStore() error {
	return nil
}

func (f *storeSpy) Script() script {
	return newScript()
}

func (f *storeSpy) Insert(script string, args ...interface{}) error {
	// initial id from first model arguments
	modelId := args[0]
	f.wasCalled = true
	return f.insert(modelId)
}

func (f *storeSpy) FindOne(rowId int, queryLang string, args ...interface{}) error {
	f.wasCalled = true
	return f.findOne(args...)
}

func (f *storeSpy) Update(script string, args ...interface{}) error {
	f.wasCalled = true
	return f.update(args...)
}

func (f *storeSpy) IsWasCalled() bool {
	return f.wasCalled
}
