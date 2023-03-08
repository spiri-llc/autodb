package autoslice

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type AutoDB struct {
	DB *sql.DB
}

// AutoQuery streams through a channel the results of your sql
// and converts it to a string array. This allows easier querying
// and easy indexing without the hastle of dealing with individual
// data types for every query.
//
// This significantly slows down performance but allows for easier use without the fuss.
func (a *AutoDB) AutoQuery(sql string) <-chan []string {
	ch := make(chan []string)
	go func() {
		results, err := a.DB.Query(sql)
		if err != nil {
			close(ch)
			return
		}

		columns, _ := results.Columns()
		count := len(columns)
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		for results.Next() {
			for i := range columns {
				valuePtrs[i] = &values[i]
			}

			results.Scan(valuePtrs...)

			arr := make([]string, count)
			var v interface{}
			for i, _ := range columns {
				val := values[i]

				b, ok := val.([]byte)
				if ok {
					v = string(b)
				} else {
					v = val
				}

				// arr[i] = fmt.Sprint(v)
				arr[i] = v.(string)
			}
			ch <- arr
		}
		close(ch)
	}()
	return ch
}
