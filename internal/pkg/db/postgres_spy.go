package db

type StoreSpy interface {
	Store
	IsWasCalled() bool // additional function
}

type storeSpy struct {
	wasCalled bool
	insert    func(...interface{}) error
	findOne   func(...interface{}) error
	find      func(...interface{}) ([]interface{}, error)
	update    func(...interface{}) error
}

func NewStoreSpy(
	insert func(...interface{}) error,
	findOne func(...interface{}) error,
	find func(...interface{}) ([]interface{}, error),
	update func(...interface{}) error,
) *storeSpy {
	return &storeSpy{
		insert:  insert,
		findOne: findOne,
		find:    find,
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

func (f *storeSpy) Find(script string, model interface{}, args ...interface{}) ([]interface{}, error) {
	f.wasCalled = true
	return f.find(args...)
}

func (f *storeSpy) FindOne(rowId int, script string, args ...interface{}) error {
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
