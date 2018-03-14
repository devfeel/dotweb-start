package demo

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb-start/server/contract"
	"github.com/devfeel/dotweb-start/const"
	"strconv"
	"github.com/devfeel/dotweb-start/protected/model"
	"github.com/devfeel/dotweb-start/protected/service/demo"
)

const (
	PageSize = 20
)


// QueryDemoInfo 查询Demo信息，根据传入demoid
// return code:
// 0 success
// -1001 demoid must set
// -2001 query error & err.Error()
func QueryDemoInfo(ctx dotweb.Context) error {
	var err error
	response := contract.NewResonseInfo()
	demoId := ctx.QueryInt("demoid")
	demoId_param := ctx.GetRouterName("demoid")
	if demoId_param != ""{
		demoId, err = strconv.Atoi(demoId_param)
		if err != nil{
			demoId = 0
		}
	}

	if demoId <= 0{
		return ctx.WriteJson(contract.CreateResponse(-1001, "demoid must set", nil))
	}

	//query data from service
	demoService := demo.NewDemoService()
	demoInfo, err := demoService.QueryDemoInfo(demoId)
	if err != nil {
		response.RetCode = _const.ApiRetCode_ServiceError
		response.RetMsg = err.Error()
	} else {
		response.RetCode = _const.ApiRetCode_Ok
		response.RetMsg = _const.ApiRetMsg_Ok
		response.Message = demoInfo
	}
	return ctx.WriteJson(response)
}


func QueryDemoList(ctx dotweb.Context) error {
	var err error
	response := contract.NewResonseInfo()
	//query data from service
	demoService := demo.NewDemoService()
	demoInfo, err := demoService.QueryDemoList(10)
	if err != nil {
		response.RetCode = _const.ApiRetCode_ServiceError
		response.RetMsg = err.Error()
	} else {
		response.RetCode = _const.ApiRetCode_Ok
		response.RetMsg = _const.ApiRetMsg_Ok
		response.Message = demoInfo
	}
	return ctx.WriteJson(response)
}

func AddDemo(ctx dotweb.Context) error {
	var err error
	response := contract.NewResonseInfo()
	demoInfo := &model.DemoInfo{
		DemoID:10000,
		DemoName : "demoapp",
	}

	//query data from service
	demoService := demo.NewDemoService()
	err = demoService.AddDemo(demoInfo)
	if err != nil {
		response.RetCode = _const.ApiRetCode_ServiceError
		response.RetMsg = err.Error()
	} else {
		response.RetCode = _const.ApiRetCode_Ok
		response.RetMsg = _const.ApiRetMsg_Ok
		response.Message = demoInfo
	}
	return ctx.WriteJson(response)
}