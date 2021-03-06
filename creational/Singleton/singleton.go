package Singleton

/**
  $ Solution
*/
type Singleton interface {
    AddOne() int
}

type singleton struct {
    count int
}

func (s *singleton) AddOne() int {
    s.count++
    return s.count
}

var instance *singleton

func GetInstance() Singleton {
    if instance != nil {
        return instance
    }
    instance = &singleton{
        1,
    }
    return instance
}
