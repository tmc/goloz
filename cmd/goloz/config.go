package main

// RunConfig describes the configuration under which to run
type RunConfig struct {
	ServerAddr string
	Insecure   bool
	LocalOnly  bool

	WindowIdx int // optionally used to place initial windows differently.
}
