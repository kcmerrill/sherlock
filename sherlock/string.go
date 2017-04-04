package sherlock

import (
	"fmt"
	"sync"
	"time"
)

// NewString inits the string struct
func NewString() *String {
	return &String{
		created:  time.Now(),
		modified: time.Now(),
		lock:     &sync.Mutex{},
	}
}

// String property type
type String struct {
	created  time.Time
	modified time.Time
	value    string
	lock     *sync.Mutex
}

// Reset string to ""
func (s *String) Reset() {
	s.lock.Lock()
	s.value = ""
	s.lock.Unlock()
}

// Set the value to be something
func (s *String) Set(something interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value = something.(string)
	s.modified = time.Now()
}

// LastModified returs the last modified time
func (s *String) LastModified() time.Time {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.modified
}

// Created returns the created time
func (s *String) Created() time.Time {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.created
}

// String returns the strings value
func (s *String) String() (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.value, nil
}

// Int not used
func (s *String) Int() (int, error) {
	return 0, fmt.Errorf("string does not implement IntValue")
}

// List converts string to a list
func (s *String) List() ([]string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return []string{s.value}, nil
}

// Add not implemented
func (s *String) Add(something interface{}) {}

// Remove not implemented
func (s *String) Remove(something interface{}) {}
