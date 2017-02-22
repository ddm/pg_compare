package pg

import (
	"database/sql"
	"fmt"
)

func (d *PGDriver) Columns(table string) (map[string]string, error) {
	rows, err := d.DB.Query(`SELECT DISTINCT
		a.attname								AS name,
		format_type(a.atttypid, a.atttypmod)	AS type,
		f.adsrc									AS default,
		NOT a.attnotnull						AS nullable,
		a.attnum								AS attnum
		FROM pg_attribute a
		LEFT JOIN pg_attrdef f ON f.adrelid = a.attrelid  AND f.adnum = a.attnum
		WHERE a.attnum > 0
			AND NOT a.attisdropped
			AND a.attrelid = $1::regclass
		ORDER BY a.attnum
	`, table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := map[string]string{}
	for rows.Next() {
		var name, typ, sdef, attnum string
		var def sql.NullString
		var nullable bool
		if err := rows.Scan(&name, &typ, &def, &nullable, &attnum); err != nil {
			return nil, err
		}

		if def.Valid {
			sdef = def.String
		}

		columns[name] = fmt.Sprintf("%s-%s-%s-%b", typ, sdef, attnum, nullable)
	}

	return columns, nil
}
