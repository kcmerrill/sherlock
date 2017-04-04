package sherlock

import (
	"strconv"
	"sync"
	"time"
)

// NewInt inits the Int struct
func NewInt() *Int {
	return &Int{
		created:  time.Now(),
		modified: time.Now(),
		lock:     &sync.Mutex{},
	}
}

// Int property type
type Int struct {
	created  time.Time
	modified time.Time
	value    int
	lock     *sync.Mutex
}

// Reset Int to ""
func (i *Int) Reset() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value = 0
}

// Set the value to be something
func (i *Int) Set(something interface{}) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value = something.(int)
	i.modified = time.Now()
}

// LastModified returs the last modified time
func (i *Int) LastModified() time.Time {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.modified
}

// Created returns the created time
func (i *Int) Created() time.Time {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.created
}

// String returns the Ints value
func (i *Int) String() (string, error) {
	i.lock.Lock()
	defer i.lock.Unlock()
	return strconv.Itoa(i.value), nil
}

// Int returns the ints value
func (i *Int) Int() (int, error) {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value, nil
}

// List converts Int to a list
func (i *Int) List() ([]string, error) {
	i.lock.Lock()
	defer i.lock.Unlock()
	return []string{strconv.Itoa(i.value)}, nil
}

// Add increments value by something
func (i *Int) Add(something interface{}) {
	i.lock.Lock()
	defer i.lock.Unlock()
	by := something.(int)
	i.value += by
	i.modified = time.Now()
}

// Remove not implemented
func (i *Int) Remove(something interface{}) {}
