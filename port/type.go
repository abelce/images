package port

import (
	"admin/domain"
	// "admin/application"
	"database/sql"
	// "fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	// "time"
)

func GetTypeList(_type string) ([]*domain.Type, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	types := []*domain.Type{}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, _ := db.Prepare(`SELECT * FROM admin.types where type=?`)
	rows, err := stmt.Query(_type)
	if err != nil {
		log.Fatal(`查询失败`)
		return types, err
	}

	for rows.Next() {
		newType := domain.Type{}
		rows.Scan(
			&newType.ID,
			&newType.Type,
			&newType.Label,
			&newType.Value,
		)
		types = append(types, &newType)
	}

	return types, nil
}