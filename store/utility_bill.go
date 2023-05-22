package store

import "database/sql"

type UtilityBill struct {
	water       float32
	electricity float32
	createdAt   string
}

func FindUtilityBill() (UtilityBill, error) {
	db, err := sql.Open("mysql", "用户名:密码@tcp(localhost:3306)/数据库名")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("SELECT * FROM electricity ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	return memoOrganizer, nil
}
