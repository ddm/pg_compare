package compare

import (
	"errors"
	"fmt"

	"github.com/jelmersnoeck/pg_compare/drivers"
	"github.com/jelmersnoeck/pg_compare/log"
)

func Columns(source, destination drivers.Driver) error {
	sColumns, err := getColumnsSchema(source)
	if err != nil {
		return err
	}
	dColumns, err := getColumnsSchema(destination)
	if err != nil {
		return err
	}

	if len(sColumns) != len(dColumns) {
		return fmt.Errorf("Destination (%d) does not have the same amount of tables as source (%d)", len(dColumns), len(sColumns))
	}

	mismatches := map[string]map[string]map[string]string{}
	for table, columns := range sColumns {
		dstColumns, ok := dColumns[table]
		if !ok {
			return fmt.Errorf("Destination does not have source table `%s`", table)
		}

		mm := map[string]map[string]string{}
		for column, schema := range columns {
			dstColumn, ok := dstColumns[column]
			if !ok {
				return fmt.Errorf("Destination does not have source column `%s` for table `%s`", column, table)
			}

			if dstColumn != schema {
				mm[column] = map[string]string{
					"source":      schema,
					"destination": dstColumn,
				}
			}
		}

		if len(mm) != 0 {
			mismatches[table] = mm
		}
	}

	if len(mismatches) != 0 {
		log.Printf("Column mismatches: %s", mismatches)
		return errors.New("mismatches")
	}

	return nil
}

func getColumnsSchema(d drivers.Driver) (map[string]map[string]string, error) {
	tables, err := d.Tables()
	if err != nil {
		return nil, err
	}

	tableColumns := map[string]map[string]string{}
	for _, table := range tables {
		log.Printf("Getting columns for table `%s` for host `%s`\n", table, d.Host())

		tc, err := d.Columns(table)
		if err != nil {
			return nil, err
		}

		tableColumns[table] = tc
	}

	return tableColumns, nil
}
