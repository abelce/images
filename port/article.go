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

func CreateArticle(article *domain.Article) (*domain.Article, error){
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()
	
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}

	stmt, _ := db.Prepare(`INSERT 
		admin.article
		(
			id,
			title,
			markdowncontent,
			private,
			tags,
			status,
			categories,
			type,
			description,
			createTime,
			lastUpdateTime
		) VALUES (?,?,?,?,?,?,?,?,?,?,?)`)
    row, err := stmt.Exec(
		article.ID, 
		article.Title,
		article.Markdowncontent,
		article.Private,
		article.Tags,
		article.Categories,
		article.Type,
		article.Description,
		article.CreateTime,
		article.LastUpdateTime,
	)
	if err != nil {
		return nil, err
	}
	id, _ := row.LastInsertId()
	fmt.Println(id)

	return article, nil
}

// func UpdateArticle(article *domain.Article) (*domain.Article, error){
// 	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
// 	defer db.Close()
	
// 	if err != nil {
// 		return nil, err
// 		// log.Fatal(err)
// 	}

// 	stmt, _ := db.Prepare(`UPDATE 
// 		admin.article
// 		(
// 			id,
// 			title,
// 			markdowncontent,
// 			private,
// 			tags,
// 			status,
// 			categories,
// 			type,
// 			description
// 		) VALUES (?,?,?,?,?,?,?,?,?,?)`)
// 	_, err := stmt.Exec(
// 		article.ID, 
// 		article.Title,
// 		article.Markdowncontent,
// 		article.Private,
// 		article.Tags,
// 		article.Categories,
// 		article.Type,
// 		article.Description
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return article, nil
// }