package main

// RunConfig describes the configuration under which to run
type RunConfig struct {
	ServerAddr   string
	UserIdentity string
	Insecure     bool
	LocalOnly    bool
}