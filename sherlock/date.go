package sherlock

import (
	"sync"
	"time"
)

// NewDate inits the Date struct
func NewDate() *Date {
	return &Date{
		created:  time.Now(),
		modified: time.Now(),
		lock:     &sync.Mutex{},
	}
}

// Date property type
type Date struct {
	created  time.Time
	modified time.Time
	value    time.Time
	lock     *sync.Mutex
}

// Reset Date to now
func (d *Date) Reset() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.value = time.Now()
}

// Set the value to be something
func (d *Date) Set(something interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.value = something.(time.Time)
	d.modified = time.Now()
}

// LastModified returs the last modified time
func (d *Date) LastModified() time.Time {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.modified
}

// Created returns the created time
func (d *Date) Created() time.Time {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.created
}

// String returns the Dates value
func (d *Date) String() (string, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.value.String(), nil
}

// Int returns the Dates value
func (d *Date) Int() (int, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	return int(d.value.Unix()), nil
}

// List converts Date to a list
func (d *Date) List() ([]string, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	return []string{d.value.String()}, nil
}

// Add increments value by something
func (d *Date) Add(something interface{}) {
	d.lock.Lock()
	defer d.lock.Unlock()
	by := something.(string)
	if dur, err := time.ParseDuration(by); err == nil {
		d.value.Add(dur)
		d.modified = time.Now()
	}
}

// Remove not implemented
func (d *Date) Remove(something interface{}) {}
