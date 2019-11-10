package repository

import (
	"github.com/devfeel/database/mssql"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/dotweb-start/const"
)

type BaseRepository struct {
	mssql.MsSqlDBContext
	databaseLogger dotlog.Logger
}

func (base *BaseRepository) InitLogger() {
	base.databaseLogger = dotlog.GetLogger(_const.LoggerName_Repository)
	base.GetCommand().SetOnTrace(base.Trace)
	base.GetCommand().SetOnDebug(base.Debug)
	base.GetCommand().SetOnInfo(base.Info)
	base.GetCommand().SetOnWarn(base.Warn)
	base.GetCommand().SetOnError(base.Error)
}

func (base *BaseRepository) Trace(content interface{}) {
	if base.databaseLogger != nil {
		base.databaseLogger.Trace(content)
	}
}

func (base *BaseRepository) Debug(content interface{}) {
	if base.databaseLogger != nil {
		base.databaseLogger.Debug(content)
	}
}

func (base *BaseRepository) Info(content interface{}) {
	if base.databaseLogger != nil {
		base.databaseLogger.Info(content)
	}
}

func (base *BaseRepository) Warn(content interface{}) {
	if base.databaseLogger != nil {
		base.databaseLogger.Warn(content)
	}
}

func (base *BaseRepository) Error(err error, content interface{}) {
	if base.databaseLogger != nil {
		base.databaseLogger.Error(err, content)
	}
}
