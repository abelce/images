package port

import (
	"admin/domain"
	// "admin/application"
	"database/sql"
	"fmt"
	"log"
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
		article.Status,
		article.Categories,
		article.Type,
		article.Description,
		article.CreateTime,
		article.LastUpdateTime,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	id, _ := row.LastInsertId()
	fmt.Println(id)

	return article, nil
}


func UpdateArticle(id string, article *domain.Article) (*domain.Article, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	// log.Fatal("update")
	if err != nil {
		return nil, err
		// log.Fatal(err)
	}
	
	stmt, _ := db.Prepare(`UPDATE admin.article SET 
		title=?,
		markdowncontent=?,
		private=?,
		tags=?,
		status=?,
		categories=?,
		type=?,
		description=?,
		lastUpdateTime=? WHERE id=?
		`)
	res, err := stmt.Exec(
		article.Title,
		article.Markdowncontent,
		article.Private,
		article.Tags,
		article.Status,
		article.Categories,
		article.Type,
		article.Description,
		article.LastUpdateTime,
		id,
	)
	if err != nil {
		return nil, err
	}
	num, err := res.RowsAffected()
	fmt.Println(num)

	if err != nil || num == 0 {
		return nil, err
	}

	return article, nil
}


func GetArticle(id string) (*domain.Article, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()
	
	if err != nil {
		log.Fatal(err)		
		return nil, err
	}
	
	stmt, _ := db.Prepare(`SELECT * FROM admin.article WHERE id=?`)
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(`查询${id}失败`)
		return nil, err
	}

	article := domain.Article{};
	row.Scan(
		&article.ID,
		&article.Title,
		&article.Markdowncontent,
		&article.Private,
		&article.Tags,
		&article.Categories,
		&article.Type,
		&article.Description,
		&article.CreateTime,
		&article.LastUpdateTime,
		&article.Status,
	 )
	return &article, nil
}

func GetArticleList(offset int, end int) ([]*domain.Article, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	articles := []*domain.Article{}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, _ := db.Prepare(`SELECT * FROM admin.article LIMIT ?,?`)
	rows, err := stmt.Query(offset, end)
	if err != nil {
		log.Fatal(`查询失败`)
		return articles, err
	}

	for rows.Next() {
		article := domain.Article{}
		rows.Scan(
			&article.ID,
			&article.Title,
			&article.Markdowncontent,
			&article.Private,
			&article.Tags,
			&article.Categories,
			&article.Type,
			&article.Description,
			&article.CreateTime,
			&article.LastUpdateTime,
			&article.Status,
		)
		articles = append(articles, &article)
	}

	return articles, nil
}


func ArticleTotal() (int, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	var total int;
	stmt, _ := db.Prepare(`SELECT count(*) FROM admin.article`)
	row := stmt.QueryRow()
	if err != nil {
		log.Fatal(`查询${id}失败`)
		return 0, err
	}
	row.Scan(&total);
	return total, nil;
}