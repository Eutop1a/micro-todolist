package script

import (
	"context"
	"todo_list/app/task/repository/mq/task"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskService(ctx)
	if err != nil {
		return
	}
}
