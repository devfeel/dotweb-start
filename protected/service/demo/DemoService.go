package demo

import (
	"github.com/devfeel/dotweb-start/const"
	"github.com/devfeel/dotweb-start/protected"
	"github.com/devfeel/dotweb-start/protected/model"
	"github.com/devfeel/dotweb-start/protected/repository"
	"github.com/devfeel/dotweb-start/protected/service"
	"errors"
	"strconv"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/cache"
	"github.com/devfeel/mapper"
)

var (
	defaultDemoRepository *repository.DemoRepository
	defaultLogger dotlog.Logger
)

const (
	RedisKey_DemoInfoID = _const.RedisKey_ProjectPre + "DemoInfoID:"
	loggerName_DemoService = "DemoServiceLogger"
)

type DemoService struct {
	service.BaseService
	demoRepository *repository.DemoRepository
	}

func init() {
	protected.RegisterServiceLoader("conf", serviceLoader)
}

func serviceLoader() {
	defaultDemoRepository = repository.NewDemoRepository(protected.DefaultConfig)
	defaultLogger = dotlog.GetLogger(loggerName_DemoService)
}

// NewDemoService create ConfService use default repository config
func NewDemoService() *DemoService {
	service := &DemoService{
		demoRepository: defaultDemoRepository,
	}
	service.RedisCache = cache.GetRedisCache(protected.DefaultConfig.DefaultRedisConn)
	return service
}

// QueryDemoInfo 根据指定DemoID查询DemoInfo
func (service *DemoService) QueryDemoInfo(demoId int) (*model.DemoInfo, error) {
	if demoId <= 0 {
		return nil, errors.New("must set demoId")
	}
	result := new(model.DemoInfo)
	var resultMap map[string]interface{}
	var err error
	redisKey := RedisKey_DemoInfoID + strconv.Itoa(demoId)
	//get from redis
	err = service.RedisCache.GetJsonObj(redisKey, result)
	if err == nil {
		return result, nil
	}

	err = service.demoRepository.QueryDemoInfo(result, demoId)
	if err == nil {
		if len(resultMap) >0 {
			//convert to struct
			errMapper := mapper.MapperMap(resultMap, result)
			if errMapper != nil {
				return nil, errMapper
			}
			//set to redis
			service.RedisCache.SetJsonObj(redisKey, result)
		}else{
			result = nil
			err = errors.New("not exists this demo info")
		}
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
		if len(results) <=0 {
			results = nil
			err = errors.New("not exists this demo info")
		}
	}
	return results, err
}

func (service *DemoService) AddDemo(demo *model.DemoInfo)error{
	if demo == nil{
		return errors.New("must set demoinfo")
	}
	_, err:= service.demoRepository.InsertDemo(demo)
	defaultLogger.InfoFormat("AddDemo", *demo, err)
	return err
}