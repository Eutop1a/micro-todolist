package script

import (
	"context"
	"todo_list/app/task/repository/mq/task"
	log "todo_list/pkg/logger"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskService(ctx)
	if err != nil {
		log.LogrusObj.Infof("RunTaskCreate:%s", err)
	}
}
