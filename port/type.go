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

func GetTypeList() ([]*domain.Type, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	types := []*domain.Type{}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, _ := db.Prepare(`SELECT * FROM admin.type`)
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(`查询失败`)
		return types, err
	}

	for rows.Next() {
		newType := domain.Type{}
		rows.Scan(
			&newType.ID,
			&newType.Name,
		)
		types = append(types, &newType)
	}

	return types, nil
}

// func FindTypeByName(name) (domain.Type, error) {
// 	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
// 	defer db.Close()

// 	stmt, _ := db.Prepare(`Select * from admin.types where name=?`)
// 	row, err := stmt.Query(name);
// 	if (err) {
// 		return nil, err;
// 	}
// 	newType := domain.Type{}

// }


func IsTypeExist(name string) (bool, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	stmt, _ := db.Prepare(`select count(*) from admin.type where name=?`)
	row, err := stmt.Query(name);
	if err != nil {
		return false, err;
	}
	count := 0
	row.Scan(&count,)

	if count > 0 {
		return true, nil;
	}
	return false, nil;
}

func CreateType(_type *domain.Type) (*domain.Type, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	stmt, _ := db.Prepare(`INSERT admin.type (id, name) VALUES (?, ?)`)
	_, err = stmt.Exec(_type.ID, _type.Name, );
	if err != nil {
		return nil, err;
	}
	return _type, nil;
}