package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creting Single Instance Now")
				singleInstance = &single{}
			}) //creating instance
	} else {
		fmt.Println("Single Instance already created-2") // already created instances
	}
	return singleInstance
}
