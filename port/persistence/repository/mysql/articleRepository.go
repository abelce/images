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

func (p *ArticleRepository) Save(article *model.Article) (*model.Article, error) {
	id := p.NewIdentity()
	article.ID = id

	queryStr := fmt.Sprintf(`INSERT %s values ($1,$2,$3,$4,$5,$6,$7)`, p.TableName)
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
		article.Deleted,
	)
	if err != nil {
		return nil, err
	}
	
	return article, nil
}

func (p *ArticleRepository) NewIdentity() string {
	return uuid.NewV4().String()
}