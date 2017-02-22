package compare

import (
	"fmt"

	"github.com/jelmersnoeck/pg_compare/drivers"
	"github.com/jelmersnoeck/pg_compare/log"
)

func Indexes(source, destination drivers.Driver) error {
	sIdx, err := source.Indexes()
	if err != nil {
		return err
	}
	dIdx, err := destination.Indexes()
	if err != nil {
		return err
	}

	if len(sIdx) != len(dIdx) {
		return fmt.Errorf("Number of indexes for destination (%d) does not match source (%d)", len(dIdx), len(sIdx))
	}

	for table, idxs := range sIdx {
		dTable, ok := dIdx[table]
		if !ok {
			return fmt.Errorf("No destination indexes found for table %s", table)
		}

		for idx, cols := range idxs {
			dcol, ok := dTable[idx]
			if !ok {
				return fmt.Errorf("Destination does not have index `%s`", idx)
			}

			log.Printf("Comparing index `%s` with destination `%s` and source `%s`", idx, dcol, cols)
			if cols != dcol {
				return fmt.Errorf("Destination index `%s` value `%s` does not equal source `%s`", idx, dcol, cols)
			}
		}
	}

	return nil
}
