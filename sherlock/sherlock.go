package sherlock

import (
	"sync"
	"time"
)

// Entity holds our entity information
type Entity struct {
	id         string
	lock       *sync.Mutex
	properties map[string]Property
}

// Property will return an entities property
func (e *Entity) Property(name string) Property {
	e.lock.Lock()
	defer e.lock.Unlock()

	// no error checking? YOLO
	return e.properties[name]
}

// NewProperty will create and return a new property
func (e *Entity) NewProperty(name, param string) Property {
	e.lock.Lock()
	defer e.lock.Unlock()

	var p Property
	switch param {
	case "int":
		p = NewInt()
		break
	case "date":
		p = NewDate()
	case "string":
		fallthrough
	default:
		p = NewString()
	}

	e.properties[name] = p
	return e.properties[name]
}

// Created returns the entity creation date(aka the _created param)
func (e *Entity) Created() time.Time {
	created, _ := e.Property("_created").Int()
	return time.Unix(int64(created), 0)
}

// Property can be multiple things ...
type Property interface {
	Reset()
	Add(something interface{})
	Remove(something interface{})
	Set(something interface{})
	String() (string, error)
	Int() (int, error)
	List() ([]string, error)
	LastModified() time.Time
	Created() time.Time
}

// Sherlock struct
type Sherlock struct {
	lock     *sync.Mutex
	entities map[string]*Entity
}

// Entity returns a string entity, if none exist, creates one
func (s *Sherlock) Entity(id string) *Entity {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, exists := s.entities[id]; !exists {
		// we need to create a blank entity
		s.entities[id] = NewEntity(id)
	}

	return s.entities[id]
}

// New returns a newly initialized sherlock
func New() *Sherlock {
	s := &Sherlock{}
	s.lock = &sync.Mutex{}
	s.entities = make(map[string]*Entity)
	return s
}

// NewEntity returns a new entity
func NewEntity(id string) *Entity {
	e := &Entity{id: id}
	e.properties = make(map[string]Property)
	e.lock = &sync.Mutex{}
	e.NewProperty("_created", "date").Set(time.Now())
	e.NewProperty("_id", "string").Set(id)
	return e
}
