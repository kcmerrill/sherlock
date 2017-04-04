package sherlock

import "testing"

func TestSherlockEntity(t *testing.T) {
	s := New()
	s.Entity("kcmerrill@gmail.com").NewProperty("username", "string").Set("themayor")

	if name, _ := s.Entity("kcmerrill@gmail.com").Property("username").String(); name != "themayor" {
		t.Fatalf("Expected 'themayor', Actual: '%s'", name)
	}

	// make sure the entity creation time isn't zero
	if s.Entity("kcmerrill@gmail.com").Created().IsZero() {
		t.Fatalf("Created sould not be a zero time.Time")
	}

	// lets play with the counter now
	e := s.Entity("kcmerrill@gmail.com")
	e.NewProperty("counter", "int").Set(1000)

	if i, _ := s.Entity("kcmerrill@gmail.com").Property("counter").Int(); i != 1000 {
		t.Fatalf("Was expecting 'counter' to be 1000")
	}

	// Add to it
	e.Property("counter").Add(100)

	if i, _ := s.Entity("kcmerrill@gmail.com").Property("counter").Int(); i != 1100 {
		t.Fatalf("Was expecting 'counter' to be 1100")
	}
}
