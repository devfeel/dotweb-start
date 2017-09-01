package test

import (
	"fmt"
	"github.com/devfeel/dotlog"
	"github.com/devfeel/task"
)

func DealTest(ctx *task.TaskContext) error {
	fmt.Println("test deal")
	dotlog.GetLogger("TastTestLogger").Debug("test deal")
	return nil
}
