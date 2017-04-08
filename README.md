# sherlock

Entity tracker. What you track is up to you. 

[![Build Status](https://travis-ci.org/kcmerrill/sherlock.svg?branch=master)](https://travis-ci.org/kcmerrill/sherlock) [![Join the chat at https://gitter.im/kcmerrill/sherlock](https://badges.gitter.im/kcmerrill/sherlock.svg)](https://gitter.im/kcmerrill/sherlock?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![sherlock](https://raw.githubusercontent.com/kcmerrill/sherlock/master/assets/sherlock.jpg "sherlock")

## What is it

A lot of front end services enable tracking via curl requests to build up entities. For example, a user joins your site with an email address of `kcmerrill@gmail.com`. A new entity is created and off of that you can start to track things. What you deceide to track is completly up to you. You can increment counters, store strings, store lists, unique lists etc etc ... 

For now, sherlock is only in memory(eventually we will store entities in a database). More to come.

## Usage

```golang

s := sherlock.New() // create a new sherlock

// create NewProperties on the entity(date|string|int|bool|*list) *coming later
s.Entity("kcmerrill@gmail.com").NewProperty("username", "string").Set("themayor")
// lets create another entity property but with shorthand string
s.Entity("doesnotexist").S("str_does_not_exist").Set("some_value")
s.Entity("doesnotexist").I("i_does_not_exist").Set(10)

if name := s.Entity("kcmerrill@gmail.com").Property("username").String(); name != "themayor" {
    t.Fatalf("Expected 'themayor', Actual: '%s'", name)
}

// make sure the entity creation time isn't zero
if s.Entity("kcmerrill@gmail.com").Created().IsZero() {
    t.Fatalf("Created should not be a zero time.Time")
}

// lets play with the counter now
e := s.Entity("kcmerrill@gmail.com")
e.NewProperty("counter", "int").Set(1000)

if i := s.Entity("kcmerrill@gmail.com").Property("counter").Int(); i != 1000 {
    t.Fatalf("Was expecting 'counter' to be 1000")
}

// Add to it
e.Property("counter").Add(100)

if i := s.Entity("kcmerrill@gmail.com").Property("counter").Int(); i != 1100 {
    t.Fatalf("Was expecting 'counter' to be 1100")
}

```