package db

import (
	"fmt"
	"strings"
	"time"
)

type Result struct {
	Row     [][]string
	Columns []string
	Elapsed int64
}

func (c *Client) Execute(query string) (*Result, error) {
	splitedQuery := strings.Split(query, " ")
	switch strings.ToUpper(strings.TrimSpace(splitedQuery[0])) {
	case "SELECT":
		start := time.Now()
		rows, err := c.db.Query(query)
		if err != nil {
			return nil, err
		}

		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		count := len(columns)

		var resultsAny []any

		for rows.Next() {
			argsSlice := make([]any, count)

			for i := range argsSlice {
				var val any
				argsSlice[i] = &val
			}
			if err := rows.Scan(argsSlice...); err != nil {
				return nil, err
			}
			resultsAny = append(resultsAny, argsSlice...)
		}

		matrixAny := make([][]any, len(resultsAny)/count)
		for m := range len(resultsAny) / count {
			matrixAny[m] = make([]any, count)
		}

		for i, result := range resultsAny {
			matrixAny[i/count][i%count] = *result.(*any)
		}

		var finalResult Result

		for _, rr := range matrixAny {
			stringSlice := make([]string, 0, len(rr))
			for _, stringElement := range rr {
				stringSlice = append(stringSlice, fmt.Sprint(stringElement))
			}
			finalResult.Row = append(finalResult.Row, stringSlice)
		}

		finalResult.Columns = append(finalResult.Columns, columns...)

		elapsed := time.Since(start).Milliseconds()
		finalResult.Elapsed = elapsed

		return &finalResult, nil
	case "INSERT":
		return nil, fmt.Errorf("no implemeted yet")
	case "UPDATE":
		return nil, fmt.Errorf("no implemeted yet")
	case "DELETE":
		return nil, fmt.Errorf("no implemeted yet")
	default:
		return nil, fmt.Errorf("invalid operation")
	}
}
