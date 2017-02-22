package pg

func (d *PGDriver) Views() ([]string, error) {
	rows, err := d.DB.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'VIEW' ORDER BY table_name;")
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

func (d *PGDriver) ViewDefinition(view string) (string, error) {
	var def string
	err := d.DB.QueryRow("SELECT definition FROM pg_views WHERE viewname = $1", view).Scan(&def)

	return def, err
}
