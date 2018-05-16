package port

import (
	"admin/domain"
	// "admin/application"
	"database/sql"
	"fmt"
	// "log"
	_ "github.com/go-sql-driver/mysql"
	// "time"
)

func SaveUser(user *domain.User) (*domain.User, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()
	
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}

	stmt, _ := db.Prepare(`INSERT 
		admin.user 
		(
			email, 
			password,
			firstName,
			lastName,
			sex,
			phone,
			isAdmin,
			createTime,
			accessTime,
			lastUpdateTime,
			logoImage
		) VALUES (?,?,?,?,?,?,?,?,?,?)`)
	row, err := stmt.Exec(
		user.Email, 
		user.Password, 
		user.FirstName, 
		user.LastName, 
		user.Sex, 
		user.Phone, 
		user.IsAdmin,
		user.CreateTime,
		user.AccessTime,
		user.LastUpdateTime,
		user.LogoImage,
	)
	if err != nil {
		return nil, err
	}
	id, _ := row.LastInsertId()
	user.ID = string(id)
	return user, nil
}

func Login(email, password string) (*domain.User, error) {

	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()
	
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	stmts, err := db.Prepare(`SELECT id, email, firstName, lastName, sex, phone, isAdmin, logoImage FROM admin.user WHERE email=? and password=?`)
	rows, err := stmts.Query(email, password)
	if err != nil {
		return nil, err
		// fmt.Println(err)
		// panic("用户不存在")
	}
	user := &domain.User{}
	if rows.Next() {
		rows.Scan(
			&user.ID,
			&user.Email, 
			// &user.Password,
			&user.FirstName,
			&user.LastName,
			&user.Sex,
			&user.Phone,
			&user.IsAdmin,
			&user.LogoImage,
		)
	}
	return user, nil
}

func Delete(id int) error {

	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close();

	// if err != nil {
	// 	panic("链接数据库失败")
	// }

	stmts, _ := db.Prepare(`DELETE FROM admin.user where id=?`)
	result, err := stmts.Exec(id)
	num, err := result.RowsAffected();
	fmt.Println(num)

	if err != nil {
		return err
	}

	if num == 1 {
		return nil;
	}

	return err;
}


func Users(offset, end int) ([]*domain.User, error) {
	users := []*domain.User{};

	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	stmts, err := db.Prepare("SELECT * FROM admin.user LIMIT ?, ?")
	rows, err := stmts.Query(offset,end)

	if err != nil {
		return nil , err
		// panic("数据库查询错误")
	}

	for rows.Next() {
		user := domain.User{}
		rows.Scan(
			&user.Email,
			&user.FirstName,
			&user.ID,
			&user.IsAdmin,
			&user.LastName,
			&user.Password,
			&user.Phone,
			&user.Sex,
			&user.CreateTime,
			&user.LastUpdateTime,
			&user.AccessTime,
			&user.LogoImage,
		)
		users = append(users, &user)
	}

	return users, nil;
}

func Update(id int64, user *domain.User) error {

	return nil;
}