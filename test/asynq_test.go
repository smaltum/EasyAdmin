package test

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"os"
	"testing"
	"time"
)

var c *asynq.Client

//TestMain 测试
func TestMain(m *testing.M) {
	r := asynq.RedisClientOpt{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	}
	c = asynq.NewClient(r)
	ret := m.Run()
	_ = c.Close()
	os.Exit(ret)
}

func createTask() *asynq.Task {
	payload := map[string]interface{}{"user_id": 1, "message": "i'm immediately message"}
	str, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("marshal failed: %v", err)
	}
	task := asynq.NewTask("msg", str)
	return task
}

// TestEnqueue 即时消费
func TestEnqueue(t *testing.T) {
	task := createTask()
	res, err := c.Enqueue(task)
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}
	fmt.Printf("Enqueued Result: %+v\n", res)
}

//延时消费
func Test_EnqueueDelay(t *testing.T) {
	task := createTask()
	res, err := c.Enqueue(task, asynq.ProcessIn(5*time.Second))
	// res, err := c.Enqueue(task, asynq.ProcessAt(time.Now().Add(5*time.Second)))
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}
	fmt.Printf("Enqueued Result: %+v\n", res)
}

// 超时、重试、过期
func Test_EnqueueOther(t *testing.T) {
	task := createTask()
	// 10秒超时，最多重试3次，20秒后过期
	res, err := c.Enqueue(task, asynq.MaxRetry(3), asynq.Timeout(10*time.Second), asynq.Deadline(time.Now().Add(20*time.Second)))
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}
	fmt.Printf("Enqueued Result: %+v\n", res)
}
