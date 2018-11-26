package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kniren/gota/dataframe"
)

// DBのカラムを定義
type Store struct {
	ID      int
	Name    string
	Address string
}

func main() {
	db, err := sql.Open("mysql", "<user>:<password>@tcp(host:3306)/db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, address FROM store")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	user := []Store{}

	for rows.Next() {
		var store Store
		err := rows.Scan(&(store.ID), &(store.Name), &(store.Address))
		if err != nil {
			panic(err.Error())
		}
		user = append(user, store)
	}

	// 上のイテレーション内でエラーがあれば表示
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	df := dataframe.LoadStructs(user)

	fmt.Println(df)
}
