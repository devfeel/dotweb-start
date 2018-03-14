package repository

import (
	"github.com/devfeel/dotweb-start/protected"
	"github.com/devfeel/dotweb-start/protected/model"
)

type DemoRepository struct {
	BaseRepository
}

func NewDemoRepository(conf *protected.ServiceConfig) *DemoRepository {
	repository := new(DemoRepository)
	repository.Init(conf.DefaultDBConn)
	repository.InitLogger()
	return repository
}

func (repository *DemoRepository) QueryDemoInfo(demoId int) (result map[string]interface{}, err error) {
	sql := "SELECT * FROM [Demo] WITH(NOLOCK) WHERE DemoID = ? "
	return repository.FindOne(sql, demoId)
}

func (repository *DemoRepository) QueryTopDemoList(rowCount int) (result []map[string]interface{}, err error) {
	sql := "SELECT TOP 10 * FROM [Demo] WITH(NOLOCK)"
	return repository.FindList(sql)
}

func (repository *DemoRepository) InsertDemo(demo *model.DemoInfo) (n int64, err error) {
	sql := "INSERT INTO [Demo] ([DemoID], [DemoName]) VALUES(?,?)"
	return repository.Insert(sql, demo.DemoID, demo.DemoName)
}

