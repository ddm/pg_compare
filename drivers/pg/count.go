package pg

import "fmt"

func (d *PGDriver) Count(table string) (int, error) {
	var count int

	// TODO: FIX UP INJECTION
	err := d.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&count)

	return count, err
}
