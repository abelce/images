package application

type config struct {
	ENV   string `json:"env"`
	Port  uint16 `json:"port"`
	DB struct {
		DriverName     string `json:"driverName"`      // mysql
		DataSourceName string `json:"dataSourceName"`  // "abelce:Tzx_301214@tcp(111.231.192.70:3306)/images?parseTime=true"
	}
	TableName          string `json:"tableName"`  
	FileDB struct {
		DriverName     string `json:"driverName"`      // mysql
		DataSourceName string `json:"dataSourceName"`  // "abelce:Tzx_301214@tcp(111.231.192.70:3306)/images?parseTime=true"
	}
	AritcleDB struct {
		DriverName     string `json:"driverName"`      // mysql
		DataSourceName string `json:"dataSourceName"`  // "abelce:Tzx_301214@tcp(111.231.192.70:3306)/images?parseTime=true"
	}
}

func (c *config) IsDevelopment() bool {
	if c.ENV == "development" {
		return true
	}
	return false
}

func (c *config) IsProduction() bool {
	if c.ENV == "production" {
		return true
	}
	return false
}

func (c *config) IsTesting() bool {
	if c.ENV == "testing" {
		return true
	}
	return false
}