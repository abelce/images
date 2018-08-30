package mysql 

import (
	"database/sql"
	"fmt"

	"github.com/satori/go.uuid"

	"images/domain/model"
)

type ImageRepository struct {
	Client     *sql.DB
	TableName  string
}

func NewImageRepository(client *sql.DB, tableName string) *ImageRepository {
	return &ImageRepository{
		Client:    client,
		TableName: tableName,
	}
}

func (p *ImageRepository) Save(image *model.Image) error {
	id := p.NewIdentity()
	image.ID = id
	queryStr := fmt.Sprintf(`INSERT %s VALUES(?,?,?,?,?,?,?,?)`, p.TableName)
	_, err := p.Client.Exec(queryStr, 
		image.ID, 
		image.Url,
		image.SvgUrl,
		image.Width,
		image.Height, 
		image.Deleted,
		image.CreateTime,
		image.LastUpdateTime,
	)
	if err != nil {
		return err
	}
	return nil
}

// func (p *ArticleRepository)FindByID(id string) (*model.Article, error) {
// 	queryStr := fmt.Sprintf(`SELECT * FROM %s WHERE id=?`, p.TableName)
// 	article := model.Article{}
// 	err := p.Client.QueryRow(queryStr, id).Scan(
// 		&article.ID, 
// 		&article.Title, 
// 		&article.Markdowncontent,
// 		&article.Private,
// 		&article.Tags,
// 		&article.Categories,
// 		&article.Type,
// 		&article.Description,
// 		&article.CreateTime,
// 		&article.LastUpdateTime,
// 		&article.Status,
// 		&article.Deleted,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &article, nil
// }

// func (p *ArticleRepository) UpdateByID(id string, article *model.Article) (*model.Article, error) {
// 	queryStr := fmt.Sprintf(`UPDATE %s SET title=?,markdowncontent=?,private=?,tags=?,categories=?,type=?,description=?,status=?,deleted=? WHERE id=?`, p.TableName)
// 	_, err := p.Client.Exec(
// 		queryStr, 
// 		article.Title, 
// 		article.Markdowncontent, 
// 		article.Private, 
// 		article.Tags,
// 		article.Categories,
// 		article.Type,
// 		article.Description,
// 		article.Status,
// 		article.Deleted,
// 		id,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return article, nil
// }

func (p *ImageRepository) NewIdentity() string {
	id, _ := uuid.NewV4()
	return id.String()
}

func (p *ImageRepository) Find(offsetNum, limit int) (total int, images []*model.Image, err error) {
	queryStr := fmt.Sprintf(`SELECT SQL_CALC_FOUND_ROWS id, url, width, height, deleted, createTime, lastUpdateTime FROM %s LIMIT ?,?`, p.TableName)
	rows, err := p.Client.Query(queryStr, offsetNum, limit)
	if err != nil {
		return 0, nil, err
	}
	queryStr = fmt.Sprintf(`SELECT count(*) FROM %s`, p.TableName)
	err = p.Client.QueryRow(queryStr).Scan(&total)
	if err != nil {
		return 0, nil, err
	}
	if total == 0 {
		return 0, nil, nil
	}

	for rows.Next() {
		image := new(model.Image)
		rows.Scan(
			&image.ID,
			&image.Url,
			&image.SvgUrl,
			&image.Width,
			&image.Height,
			&image.Deleted,
			&image.CreateTime,
			&image.LastUpdateTime,
		)
		images = append(images, image)
	}
	
	return total, images, nil
}