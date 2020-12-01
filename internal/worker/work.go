package worker

import (
	"fmt"
	"time"

	"github.com/oleiade/lane"
	"github.com/omnis-org/omnis-server/config"
	log "github.com/sirupsen/logrus"
)

var workQueue *lane.Queue

type Work struct {
	Job    func(interface{})
	Handle interface{}
}

func initWorkQueue() {
	workQueue = lane.NewQueue()
}

func LaunchWorker() {
	initWorkQueue()
	for true {
		var work *Work
		for workQueue.Head() != nil {
			work = workQueue.Dequeue().(*Work)
			work.Job(work.Handle)
		}
		log.Debug(fmt.Sprintf("Wait for job ... (%d)", config.GetConfig().Worker.WaitWorkTime))
		time.Sleep(time.Duration(config.GetConfig().Worker.WaitWorkTime) * time.Second)
	}
}

func AddWork(work *Work) {
	workQueue.Enqueue(work)
}
