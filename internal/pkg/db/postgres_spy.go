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

func (s *storeSpy) InitStore() error {
	return nil
}

func (s *storeSpy) Script() script {
	return newScript()
}

func (s *storeSpy) Insert(script string, args ...interface{}) error {
	// initial id from first model arguments
	modelId := args[0]
	s.wasCalled = true
	return s.insert(modelId)
}

func (s *storeSpy) Find(script string, model interface{}, args ...interface{}) ([]interface{}, error) {
	s.wasCalled = true
	return s.find(args...)
}

func (s *storeSpy) FindOne(rowId int, script string, args ...interface{}) error {
	s.wasCalled = true
	return s.findOne(args...)
}

func (s *storeSpy) Update(script string, args ...interface{}) error {
	s.wasCalled = true
	return s.update(args...)
}

func (s *storeSpy) IsWasCalled() bool {
	return s.wasCalled
}
