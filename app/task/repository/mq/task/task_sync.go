package task

import (
	"context"
	"encoding/json"
	"todo_list/app/task/repository/mq"
	"todo_list/app/task/service"
	"todo_list/consts"
	"todo_list/idl/pb"
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
			// 落库
			req := new(pb.TaskRequest)
			err = json.Unmarshal(d.Body, req)
			if err != nil {
				return
			}
			err = service.TaskMQ2DB(ctx, req)
			if err != nil {
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}
