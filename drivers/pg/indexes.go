package pg

func (d *PGDriver) Indexes() (map[string]map[string]string, error) {
	rows, err := d.DB.Query(`SELECT
		t.relname as table_name,
		i.relname as index_name,
		array_to_string(array_agg(a.attname), ', ') as column_names
	FROM
		pg_class t, pg_class i, pg_index ix, pg_attribute a
	WHERE
		t.oid = ix.indrelid
		AND i.oid = ix.indexrelid
		AND a.attrelid = t.oid
		AND a.attnum = ANY(ix.indkey)
		AND t.relkind = 'r'
	GROUP BY
		t.relname, i.relname
	ORDER BY
		t.relname, i.relname`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idx := map[string]map[string]string{}
	for rows.Next() {
		var tableName, indexName, columnNames string
		if err := rows.Scan(&tableName, &indexName, &columnNames); err != nil {
			return nil, err
		}

		if _, ok := idx[tableName]; !ok {
			idx[tableName] = map[string]string{}
		}
		idx[tableName][indexName] = columnNames
	}

	return idx, nil
}
