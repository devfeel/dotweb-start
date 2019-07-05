package demo

import (
	"errors"
	"github.com/devfeel/dotweb-start/config"
	"github.com/devfeel/dotweb-start/protected/model"
	"github.com/devfeel/dotweb-start/protected/repository"
	"sync"
)

const defaultDatabaseID = "demodb"

var defaultDemoRepository *DemoRepository
var demoRepositoryLocker *sync.Mutex

func init() {
	demoRepositoryLocker = new(sync.Mutex)
}

type DemoRepository struct {
	repository.BaseRepository
}

// GetMessageRepository return MessageRepository which is inited
func GetDemoRepository() *DemoRepository {
	//check default repository is init
	if defaultDemoRepository == nil {
		demoRepositoryLocker.Lock()
		defer demoRepositoryLocker.Unlock()
		if defaultDemoRepository == nil {
			defaultDemoRepository = NewDemoRepository()
		}
	}
	return defaultDemoRepository
}

// NewMessageRepository return new MessageRepository
func NewDemoRepository() *DemoRepository {
	dbInfo, exists := config.GetDataBaseInfo(defaultDatabaseID)
	if !exists || dbInfo.ServerUrl == "" {
		err := errors.New("no config " + defaultDatabaseID + " database config")
		panic(err)
	}
	repository := new(DemoRepository)
	repository.Init(dbInfo.ServerUrl)
	repository.InitLogger()
	return repository
}

func (repository *DemoRepository) QueryDemoInfo(dest interface{}, demoId int) error {
	sql := "SELECT * FROM [Demo] WITH(NOLOCK) WHERE DemoID = ? "
	return repository.FindOne(dest, sql, demoId)
}

func (repository *DemoRepository) QueryTopDemoList(dest interface{}, rowCount int) error {
	sql := "SELECT TOP 10 * FROM [Demo] WITH(NOLOCK)"
	return repository.FindList(dest, sql)
}

func (repository *DemoRepository) InsertDemo(demo *model.DemoInfo) (n int64, err error) {
	sql := "INSERT INTO [Demo] ([DemoID], [DemoName]) VALUES(?,?)"
	return repository.Insert(sql, demo.DemoID, demo.DemoName)
}

func (repository *DemoRepository) QueryByPage(dest interface{}, skip, take int) error {
	fields := "*"
	tableName := "Demo"
	where := ""
	orderBy := "ID ASC, ID DESC"
	return repository.FindListByPage(dest, tableName, fields, where, orderBy, skip, take)
}
