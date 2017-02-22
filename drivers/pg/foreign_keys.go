package pg

import "fmt"

func (d *PGDriver) ForeignKeys() (map[string]string, error) {
	rows, err := d.DB.Query(`SELECT DISTINCT
		unique_constraint_name	AS name,
		match_option			AS option,
		update_rule				AS update,
		delete_rule				AS delete
	FROM information_schema.referential_constraints c`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	keys := map[string]string{}
	for rows.Next() {
		var name, option, update, delete string
		if err := rows.Scan(&name, &option, &update, &delete); err != nil {
			return nil, err
		}

		keys[name] = fmt.Sprintf("%s-%s-%s", option, update, delete)
	}

	return keys, nil
}
