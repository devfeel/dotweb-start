package task

import (
	"github.com/devfeel/dotweb-start/task/tasks/test"
	"fmt"
	"github.com/devfeel/dottask"
)

var (
	taskService *task.TaskService
)

func RegisterTask(service *task.TaskService) {
	service.RegisterHandler("Job_DealTest", test.DealTest)
}

func StartTaskService(configPath string) {
	taskService = task.StartNewService()
	//step 2: register all task handler
	RegisterTask(taskService)

	//step 3: load config file
	taskService.LoadConfig(configPath + "/dottask.conf")

	//step 4: start all task
	taskService.StartAllTask()

	fmt.Println(taskService.PrintAllCronTask())
}
