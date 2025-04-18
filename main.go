package main

import (
	"encoding/json"
	"fmt"
	"sync"
)



type Address struct {
	City string
	State string
	Country string
	Pincode json.Number
}

type User struct {
	Name string
	Age json.Number
	Contact string
	Company string
	Address Address
}

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, option *Options) (*Driver, error) {
	//TODO
}

func main() {
	dir := "/"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	
}