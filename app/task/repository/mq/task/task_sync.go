package task

import (
	"context"
	"encoding/json"
	"todo_list/app/task/repository/mq"
	"todo_list/app/task/service"
	"todo_list/consts"
	"todo_list/idl/pb"
	log "todo_list/pkg/logger"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskService(ctx context.Context) (err error) {
	rabbitMqQueue := consts.RabbitMQTaskQueueName
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		return
	}
	//var forever = make(chan struct{})
	var forever chan struct{}

	go func() {

		for d := range msgs {
			log.LogrusObj.Infof("Received run Task: %s", d.Body)

			// 落库
			req := new(pb.TaskRequest)
			err = json.Unmarshal(d.Body, req)
			if err != nil {
				log.LogrusObj.Infof("Received run Task: %s", err)
				//return
			}
			err = service.TaskMQ2MySQL(ctx, req)
			if err != nil {
				log.LogrusObj.Infof("Received run Task: %s", err)
				//return
			}
			d.Ack(false)
		}
	}()

	log.LogrusObj.Infoln(err)
	<-forever
	return nil
}
