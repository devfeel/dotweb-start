package tasks

import (
	"github.com/devfeel/dottask"
	"fmt"
	"time"
)

func Task_Print(context *task.TaskContext) error {
	fmt.Println(time.Now(), "Task_Print", context.TaskID)
	return nil
}
