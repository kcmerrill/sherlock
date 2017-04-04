# sherlock

Track everything

[![Build Status](https://travis-ci.org/kcmerrill/sherlock.svg?branch=master)](https://travis-ci.org/kcmerrill/sherlock) [![Join the chat at https://gitter.im/kcmerrill/sherlock](https://badges.gitter.im/kcmerrill/sherlock.svg)](https://gitter.im/kcmerrill/sherlock?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![sherlock](https://raw.githubusercontent.com/kcmerrill/sherlock/master/assets/sherlock.jpg "sherlock")

## What is it

A lot of front end services enable tracking via curl requests to build up entities. For example, a user joins your site with an email address of `kcmerrill@gmail.com`. A new entity is created and off of that you can start to track things. What you deciede to track is completly up to you. You can increment counters, store strings, store lists, unique lists etc etc ... 

For now, sherlock is only in memory(eventually we will store entities in a database). More to come.

In time we can add a webservice component but for now, in memory and backend go native services only.

## Binaries || Installation

[![MacOSX](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/apple_logo.png "Mac OSX")] (http://go-dist.kcmerrill.com/kcmerrill/sherlock/mac/amd64) [![Linux](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/linux_logo.png "Linux")] (http://go-dist.kcmerrill.com/kcmerrill/sherlock/linux/amd64)

via go:

`$ go get -u github.com/kcmerrill/sherlock`

## Usage

All sherlock is, at it's core, is a setter and getter of sorts. 

```golang

s := New() // create a new sherlock

// create NewProperties on the entity(date|string|int|*list) *coming later
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

```