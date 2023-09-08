package service

import (
	"sync"
)

var initialized uint32

type sender struct {
	ChBusiness chan string
}

var instance *sender
var once sync.Once

func Instance() *sender {

	once.Do(func() {
		instance = &sender{}
		ch := make(chan string)
		instance.ChBusiness = ch
	})

	return instance
}

func ChBusinessSend(msg string) {
	instance.ChBusiness <- msg
}
