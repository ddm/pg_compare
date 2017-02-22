package compare

import (
	"errors"
	"fmt"

	"github.com/jelmersnoeck/pg_compare/drivers"
	"github.com/jelmersnoeck/pg_compare/log"
)

type countEntry struct {
	table string
	count int
	err   error
}

func Count(source, destination drivers.Driver) error {

	sourceTables, err := source.Tables()
	if err != nil {
		return err
	}

	destinationTables, err := destination.Tables()
	if err != nil {
		return err
	}

	if len(sourceTables) != len(destinationTables) {
		return fmt.Errorf("The amount of tables in the destination (%d) does not match the source (%d)", len(destinationTables), len(sourceTables))
	}

	sourceCount, err := countPerTable(source, sourceTables)
	if err != nil {
		return err
	}
	destinationCount, err := countPerTable(destination, destinationTables)
	if err != nil {
		return err
	}

	mismatches := map[string]map[string]int{}
	for table, count := range sourceCount {
		dst, ok := destinationCount[table]
		if !ok {
			return fmt.Errorf("Destination does not have source table `%s`", table)
		}

		if dst != count {
			mismatches[table] = map[string]int{
				"source":      count,
				"destination": dst,
			}
		}
	}

	if len(mismatches) != 0 {
		fmt.Println(mismatches)
		return errors.New("mismatches")
	}

	return nil
}

func countPerTable(db drivers.Driver, tables []string) (map[string]int, error) {
	countChan := make(chan *countEntry)

	for _, table := range tables {
		go func(table string, ch chan *countEntry) {
			count, err := db.Count(table)
			ch <- &countEntry{table, count, err}
		}(table, countChan)
	}

	tableCount := map[string]int{}
	for range tables {
		select {
		case entry := <-countChan:
			log.Printf("Getting table count for %s on host %s", entry.table, db.Host())

			tableCount[entry.table] = 0
			if entry.err != nil {
				// do someething with the error
				fmt.Println(entry.err)
				continue
			}

			tableCount[entry.table] = entry.count
		}
	}

	return tableCount, nil
}
