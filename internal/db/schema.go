package db

type Table struct {
	Name    string `json:"name"`
	Columns []Column
}

type Column struct {
	Name string
	Type string
}

func (c *Client) Tables() ([]Table, error) {
	query := `
	SELECT name FROM sqlite_master WHERE type='table'
	`
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	var tables []Table
	defer rows.Close()

	for rows.Next() {
		var table Table
		if err := rows.Scan(&table.Name); err != nil {
			return nil, err
		}

		tables = append(tables, table)
	}
	return tables, nil
}

func (c *Client) Columns(table string) ([]Column, error) {
	query := `
	 SELECT name, type FROM pragma_table_info(?)
	`
	rows, err := c.db.Query(query, table)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var columns []Column

	for rows.Next() {
		var column Column

		if err := rows.Scan(&column.Name, &column.Type); err != nil {
			return nil, err
		}

		columns = append(columns, column)
	}
	return columns, nil
}
