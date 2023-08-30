package hungryload

type IInstance interface {
	Work()
}

type singleton struct{}

func (s *singleton) Work() {}

var instance *singleton

func init() {
	if instance == nil {
		instance = &singleton{}
	}
}

func GetInstance() IInstance {
	return instance
}
