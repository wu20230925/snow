package jobs

import (
	"fmt"
	"time"

	"github.com/qit-team/work"
)

func test(task work.Task) work.TaskResult {
	time.Sleep(time.Millisecond * 5)
	s, err := work.JsonEncode(task)
	if err != nil {
		// work.StateFailed 不会进行ack确认
		// work.StateFailedWithAck 会进行actk确认
		// return work.TaskResult{Id: task.Id, State: work.StateFailed}
		return work.TaskResult{Id: task.Id, State: work.StateFailedWithAck}
	} else {
		// work.StateSucceed 会进行ack确认
		fmt.Println("do task", s)
		return work.TaskResult{Id: task.Id, State: work.StateSucceed}
	}

}
