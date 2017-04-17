package main

import (
	"flag"

	"github.com/kcmerrill/sherlock/sherlock"
)

var (
	udp string
	web string
)

func main() {
	flag.StringVar(&udp, "udp-port", "8081", "The port in which to handle UDP requests")
	flag.StringVar(&web, "web-port", "80", "The port in which to handle WEB requests")
	flag.Parse()
	s := sherlock.New()
	go s.UDP(udp)
	s.Web(web)
}
