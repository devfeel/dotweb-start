package repository

import "github.com/jmoiron/sqlx"

type BaseImpl struct {
	driverName, dataSourceName string
}

func (b *BaseImpl) SetConnect(driverName, dataSourceName string) {
	b.driverName = driverName
	b.dataSourceName = dataSourceName
}

func (b *BaseImpl) Connect() (*sqlx.DB, error) {
	if b.driverName == "" {
		b.driverName = DriverName
	}
	if b.dataSourceName == "" {
		b.dataSourceName = DataSourceName
	}
	db, err := sqlx.Connect(b.driverName, b.dataSourceName)
	return db, err
}

func (b *BaseImpl) MustConnect() *sqlx.DB {
	db, err := b.Connect()
	if err != nil {
		panic(err)
	}
	return db
}
