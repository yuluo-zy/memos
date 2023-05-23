package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type UtilityBill struct {
	water       float32
	electricity float32
	createdAt   string
}

func FindUtilityBill(mysqlUrl string) (UtilityBill, error) {
	db, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err.Error())
	}
	var utilityBill UtilityBill

	err = db.QueryRow("SELECT remainder, created_at FROM electricity ORDER BY created_at DESC LIMIT 1").Scan(&utilityBill.electricity, &utilityBill.createdAt)

	err = db.QueryRow("SELECT remainder, created_at FROM water ORDER BY created_at DESC LIMIT 1").Scan(&utilityBill.water, &utilityBill.createdAt)

	defer db.Close()

	return utilityBill, nil
}
