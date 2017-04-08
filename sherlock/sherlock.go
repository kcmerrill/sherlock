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
	Events     []string
}

// Event object tracks a specific event
type Event struct {
	Created     time.Time
	Description string
}

// Property will return an entities property
func (e *Entity) Property(name string) Property {
	e.lock.Lock()
	defer e.lock.Unlock()

	// no error checking? YOLO
	return e.properties[name]
}

// Event will create a new event for an entitiy
func (e *Entity) Event(event string) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.Events = append(e.Events, event)
}

// I will create a new int property if it doesn't exist
func (e *Entity) I(name string) Property {
	e.lock.Lock()
	p, exists := e.properties[name]
	// unlock everything
	e.lock.Unlock()
	if exists {
		return p
	}
	return e.NewProperty(name, "int")
}

// S will create a new string property if it doesn't exist
func (e *Entity) S(name string) Property {
	e.lock.Lock()
	p, exists := e.properties[name]
	// unlock everything
	e.lock.Unlock()
	if exists {
		return p
	}
	return e.NewProperty(name, "string")
}

// D will create a new string property if it doesn't exist
func (e *Entity) D(name string) Property {
	e.lock.Lock()
	p, exists := e.properties[name]
	// unlock everything
	e.lock.Unlock()
	if exists {
		return p
	}
	return e.NewProperty(name, "date")
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

// E is shorthand for Entity
func (s *Sherlock) E(id string) *Entity {
	return s.Entity(id)
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
	e.Events = make([]string, 0)
	e.NewProperty("_created", "date").Set(time.Now())
	e.NewProperty("_id", "string").Set(id)
	return e
}
