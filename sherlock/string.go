package sherlock

import (
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
func (s *String) String() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.value
}

// Int not used
func (s *String) Int() int {
	return 0
}

// List converts string to a list
func (s *String) List() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	return []string{s.value}
}

// Add not implemented
func (s *String) Add(something interface{}) {}

// Remove not implemented
func (s *String) Remove(something interface{}) {}
