package pg

func (d *PGDriver) Tables() ([]string, error) {
	rows, err := d.DB.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE' ORDER BY table_name;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		names = append(names, name)
	}

	return names, rows.Err()
}
