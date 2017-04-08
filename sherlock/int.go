package sherlock

import (
	"strconv"
	"sync"
	"time"
)

// NewInt inits the Int struct
func NewInt() *Int {
	return &Int{
		CreatedDate: time.Now(),
		Modified:    time.Now(),
		lock:        &sync.Mutex{},
	}
}

// Int property type
type Int struct {
	CreatedDate time.Time `json:"created"`
	Modified    time.Time
	Value       int
	lock        *sync.Mutex
}

// Reset Int to ""
func (i *Int) Reset() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Value = 0
}

// Set the value to be something
func (i *Int) Set(something interface{}) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.Value = something.(int)
	i.Modified = time.Now()
}

// LastModified returs the last modified time
func (i *Int) LastModified() time.Time {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.Modified
}

// Created returns the created time
func (i *Int) Created() time.Time {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.CreatedDate
}

// String returns the Ints value
func (i *Int) String() string {
	i.lock.Lock()
	defer i.lock.Unlock()
	return strconv.Itoa(i.Value)
}

// Int returns the ints value
func (i *Int) Int() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.Value
}

// List converts Int to a list
func (i *Int) List() []string {
	i.lock.Lock()
	defer i.lock.Unlock()
	return []string{strconv.Itoa(i.Value)}
}

// Add increments value by something
func (i *Int) Add(something interface{}) {
	i.lock.Lock()
	defer i.lock.Unlock()
	by := something.(int)
	i.Value += by
	i.Modified = time.Now()
}

// Remove not implemented
func (i *Int) Remove(something interface{}) {}
