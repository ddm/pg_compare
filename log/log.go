package log

import (
	"log"

	"github.com/jelmersnoeck/pg_compare/config"
)

func Printf(txt string, args ...interface{}) {
	if config.Verbose {
		log.Printf(txt, args...)
	}
}
