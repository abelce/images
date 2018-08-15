package application

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	// "fmt"

	_ "github.com/go-sql-driver/mysql"

	"images/domain/model"
	_mysql "images/port/persistence/repository/mysql"
	

)
type Context struct {
	config        *config
	service       *service
	db            *sql.DB
	repository    model.Repository
	queryService  model.QueryService
}

func NewContext(path string)(*Context, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}	

	cfg := &config{}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	c := &Context{
		config: cfg,
	}

	repository, err := c.Repository()
	if err != nil {
		return nil, err
	}
	// fmt.Println(repository)
	model.DomainRegistry.Repository = repository

	qs, err := c.QueryService();
	if err != err {
		return nil, err
	}
	model.DomainRegistry.QueryService = qs

	return c, nil
}


func (c *Context) Mysql() (*sql.DB, error){
	if c.db != nil {
		return c.db, nil
	}
	db, err := sql.Open(c.config.DB.DriverName, c.config.DB.DataSourceName)
	if err != nil {
		return nil, err
	}

	c.db = db
	return c.db, err
}

func (c *Context) Repository() (model.Repository, error){

	if c.repository != nil {
		return c.repository, nil
	}

	db, err := c.Mysql()
	if err != nil {
		return nil, err
	}
	repository := &_mysql.ImageRepository{
		Client: db,
		TableName: c.config.TableName,
	}
	c.repository = repository

	return c.repository, nil
}


func (c *Context)QueryService() (model.QueryService, error) {
	if c.queryService != nil {
		return c.queryService, nil
	}

	db, err := c.Mysql()
	if err != nil {
		return nil, err
	}

	queryService := _mysql.NewImageRepository(db, c.config.TableName)
	c.queryService = queryService

	return c.queryService, nil
}

func (c *Context)Service() (*service, error){
	repository, err := c.Repository()
	if err != nil {
		return nil, err
	}

	queryService, err := c.QueryService()
	if err != nil {
		return nil, err
	}

	if c.service == nil {
		c.service = &service{
			Repository: repository,
		    QueryService: queryService,
		}
	}

	return c.service, nil
}

func (c *Context) Config() *config{
	return c.config
}