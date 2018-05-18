package task

import (
	"github.com/devfeel/dottask"
	"github.com/devfeel/dotweb-start/global"
	"github.com/devfeel/dotweb-start/task/tasks"
	"fmt"
)

func registerTask(service *task.TaskService) {
	//TODO register task to service
	service.RegisterHandler("Task_Print", tasks.Task_Print)
}

func StartTaskService(configPath string) {
	global.DotTask = task.StartNewService()

	//register all task handler
	registerTask(global.DotTask)

	//load config file
	global.DotTask.LoadConfig(configPath + "/dottask.conf")

	//start all task
	global.DotTask.StartAllTask()

	global.InnerLogger.Debug(fmt.Sprint("StartTaskService", " ", configPath, " ", global.DotTask.PrintAllCronTask()))
}

func StopTaskService() {
	if global.DotTask != nil {
		global.DotTask.StopAllTask()
	}
}
