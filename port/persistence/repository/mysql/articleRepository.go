package mysql 

import (
	"database/sql"
	"fmt"

    "github.com/satori/go.uuid"

	"admin/domain/model"
)

type ArticleRepository struct {
	Client     *sql.DB
	TableName  string
}

func NewArticleRepository(client *sql.DB, tableName string) *ArticleRepository {
	return &ArticleRepository{
		Client:    client,
		TableName: tableName,
	}
}

func (p *ArticleRepository) Save(article *model.Article) error {
	id := p.NewIdentity()
	article.ID = id
	queryStr := fmt.Sprintf(`INSERT %s VALUES(?,?,?,?,?,?,?,?,?,?,?,?)`, p.TableName)
	_, err := p.Client.Exec(queryStr, 
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
		article.Status,
		article.Deleted,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *ArticleRepository)FindByID(id string) (*model.Article, error) {
	queryStr := fmt.Sprintf(`SELECT * FROM %s WHERE id=?`, p.TableName)
	article := model.Article{}
	err := p.Client.QueryRow(queryStr, id).Scan(
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
		&article.Deleted,
	)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (p *ArticleRepository) UpdateByID(id string, article *model.Article) (*model.Article, error) {
	queryStr := fmt.Sprintf(`UPDATE %s SET title=?,markdowncontent=?,private=?,tags=?,categories=?,type=?,description=?,status=?,deleted=? WHERE id=?`, p.TableName)
	_, err := p.Client.Exec(
		queryStr, 
		article.Title, 
		article.Markdowncontent, 
		article.Private, 
		article.Tags,
		article.Categories,
		article.Type,
		article.Description,
		article.Status,
		article.Deleted,
		id,
	)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (p *ArticleRepository) NewIdentity() string {
	id, _ := uuid.NewV4()
	return id.String()
}

func (p *ArticleRepository) Find(offsetNum, limit int) (total int, articles []*model.Article, err error) {
	queryStr := fmt.Sprintf(`SELECT SQL_CALC_FOUND_ROWS id, title, markdowncontent, private, tags, categories,type,description,createTime, lastUpdateTime, status, deleted  FROM %s LIMIT ?,?`, p.TableName)
	rows, err := p.Client.Query(queryStr, offsetNum, limit)
	if err != nil {
		return 0, nil, err
	}
	queryStr = fmt.Sprintf(`SELECT count(*) FROM %s`, p.TableName)
	err = p.Client.QueryRow(queryStr).Scan(&total)
	fmt.Println(total)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, nil, nil
	}

	for rows.Next() {
		article := new(model.Article)
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
			&article.Deleted,
		)
		articles = append(articles, article)
	}
	
	return total, articles, nil
}