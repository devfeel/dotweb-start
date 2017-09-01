package repository

var (
	DriverName     string
	DataSourceName string
)

const Mysql_DirverName = "mysql"

// InitConfig init database's config
func InitConfig(dataSourceName string) {
	DriverName = Mysql_DirverName
	DataSourceName = dataSourceName
}
