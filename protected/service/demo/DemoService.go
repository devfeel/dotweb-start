package demo

import (
	"database/sql"
	"errors"
	"github.com/devfeel/cache"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/dotweb-start/config"
	"github.com/devfeel/dotweb-start/const"
	"github.com/devfeel/dotweb-start/protected/model"
	"github.com/devfeel/dotweb-start/protected/repository/demo"
	"github.com/devfeel/dotweb-start/protected/service"
	"strconv"
)

var (
	defaultDemoRepository *demo.DemoRepository
	defaultLogger         dotlog.Logger
)

const (
	defaultRedisID         = "DefaultRedis"
	RedisKey_DemoInfoID    = _const.RedisKey_ProjectPre + "DemoInfoID:"
	loggerName_DemoService = "DemoServiceLogger"
)

type DemoService struct {
	service.BaseService
	demoRepository *demo.DemoRepository
}

func init() {
	defaultLogger = dotlog.GetLogger(loggerName_DemoService)
}

// NewDemoService create ConfService use default repository config
func NewDemoService() *DemoService {
	service := &DemoService{
		demoRepository: demo.GetDemoRepository(),
	}

	redisInfo, exists := config.GetRedisInfo(defaultRedisID)
	if !exists || redisInfo.ServerUrl == "" {
		err := errors.New("no config " + defaultRedisID + " redis config")
		panic(err)
	}
	service.RedisCache = cache.GetRedisCache(redisInfo.ServerUrl)
	return service
}

// QueryDemoInfo 根据指定DemoID查询DemoInfo
func (service *DemoService) QueryDemoInfo(demoId int) (*model.DemoInfo, error) {
	if demoId <= 0 {
		return nil, errors.New("must set demoId")
	}
	result := new(model.DemoInfo)
	var err error
	redisKey := RedisKey_DemoInfoID + strconv.Itoa(demoId)
	//get from redis
	err = service.RedisCache.GetJsonObj(redisKey, result)
	if err == nil {
		return result, nil
	}

	err = service.demoRepository.QueryDemoInfo(result, demoId)
	if err == nil {
		service.RedisCache.SetJsonObj(redisKey, result)
	} else if err == sql.ErrNoRows {
		result = nil
		err = errors.New("not exists this demo info")
	}
	return result, err
}

// QueryDemoList 根据指定记录数查询记录
func (service *DemoService) QueryDemoList(rowNum int) ([]*model.DemoInfo, error) {
	if rowNum <= 0 {
		return nil, errors.New("must set rowNum")
	}
	var results []*model.DemoInfo
	var err error
	err = service.demoRepository.QueryTopDemoList(&results, rowNum)
	if err == nil {
		if len(results) <= 0 {
			results = nil
			err = errors.New("not exists this demo info")
		}
	}
	return results, err
}

func (service *DemoService) AddDemo(demo *model.DemoInfo) error {
	if demo == nil {
		return errors.New("must set demoinfo")
	}
	_, err := service.demoRepository.InsertDemo(demo)
	defaultLogger.InfoS("AddDemo", *demo, err)
	return err
}
