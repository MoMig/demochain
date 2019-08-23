package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(	) *sql.DB{
	db,err := sql.Open("mysql","root:13971334478@tcp(127.0.0.1:3306)/blockdb")
	if err != nil {
		fmt.Println(err)
	}
	return db
	}

func Query(db sql.DB,sql string) *sql.Rows{
	rows,err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func QueryWithArgs(db sql.DB,sql string,args ...interface{}) sql.Rows{
	stmt,err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	rows,err := stmt.Query(args ...)

	if err != nil {
		fmt.Println(err)
	}
	return *rows
}

func Insert(db sql.DB,sql string,args ...interface{}) int64 {
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	res,err := stmt.Exec(args...)
	if err != nil {
		fmt.Println(err)
	}
	cnt,_ := res.RowsAffected()
	return cnt
}

func HandleRls(rows *sql.Rows) map[int]map[string]string {
	results := make(map[int]map[string]string) //最后得到的map
	columns, _ := rows.Columns()
	cur := 0
	for rows.Next() {
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}
		err := rows.Scan(scanArgs...)
		//将数据保存到 record 字典
		if err != err {
			fmt.Println(err)
		}

		row := make(map[string]string) //每行数据
		for k, v := range values {
			if v != nil {
				key := columns[k]
				row[key] = string(v.([]byte))
			}
		}

		results[cur] = row
		cur++
	}

	defer  rows.Close()

	return results
}

/*func check(err error) bool{
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}*/