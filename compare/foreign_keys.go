package compare

import (
	"errors"
	"fmt"

	"github.com/jelmersnoeck/pg_compare/drivers"
	"github.com/jelmersnoeck/pg_compare/log"
)

func ForeignKeys(source, destination drivers.Driver) error {
	sKeys, err := source.ForeignKeys()
	if err != nil {
		return err
	}
	dKeys, err := destination.ForeignKeys()
	if err != nil {
		return err
	}

	if len(dKeys) != len(sKeys) {
		return fmt.Errorf("Destination keys (%d) do not match source keys (%d)", len(dKeys), len(sKeys))
	}

	mismatches := map[string]map[string]string{}
	for key, value := range sKeys {
		dVal, ok := dKeys[key]
		if !ok {
			return fmt.Errorf("Destination does not have foreign key `%s`", key)
		}

		if dVal != value {
			mismatches[key] = map[string]string{
				"source":      value,
				"destination": dVal,
			}
		}
	}

	if len(mismatches) != 0 {
		log.Printf("ForeignKeys mismatches: %s", mismatches)
		return errors.New("mismatches")
	}

	return nil
}
