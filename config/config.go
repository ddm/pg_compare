package config

var Verbose = false

type Config struct {
	// Database options
	Source      string
	Destination string

	// Comparison Options
	Rows        bool
	Columns     bool
	ForeignKeys bool
	Indexes     bool
	Views       bool
}
