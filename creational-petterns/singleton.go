package creational_petterns

import (
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func fillSingleInstance() {
	lock.Lock()
	defer lock.Unlock()
	if singleInstance == nil {
		singleInstance = &single{}
	}
}

func GetInstance() *single {

	if singleInstance == nil {
		fillSingleInstance()
	}

	return singleInstance
}

// use this design pattern when you want to have one instance of object in your application
