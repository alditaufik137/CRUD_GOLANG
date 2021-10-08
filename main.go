package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123123"
	dbname   = "Tugas_Backend"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlInsert := `
		INSERT INTO final_setting_unor (id, instansi_id, instansi_nama, step_unor_verifikasi_approval)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	id := 0
	err = db.QueryRow(sqlInsert, 1, "22", "aasd", "3").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("done")

	sqlUpdate := `
		UPDATE final_setting_unor
		SET instansi_id = $2, instansi_nama = $3, step_unor_verifikasi_approval = $4,
		WHERE id = $1;`
	_, err = db.Exec(sqlUpdate, 2, "by.u", "32", "1")
	if err != nil {
		panic(err)
	}
	fmt.Println("done")

	sqlDelete := `
		DELETE FROM final_setting_unor WHERE id = $1;`
	_, err = db.Exec(sqlDelete, 1)
	if err != nil {
		panic(err)
	}

	sqlSelect := `
		SELECT FROM final_setting_unor WHERE id = $1;`
	_, err = db.Exec(sqlSelect, 1)
	if err != nil {
		panic(err)
	}
}
