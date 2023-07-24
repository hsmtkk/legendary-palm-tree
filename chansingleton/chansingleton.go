package chansingleton

type Singleton interface {
	AddOne()
}

type singletonImpl struct {
	addCh chan bool
}

var instance singletonImpl

func GetInstance() Singleton {
	return &instance
}

func (s *singletonImpl) AddOne() {
	s.addCh <- true
}
