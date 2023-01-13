package bootstrap

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"time"
)

var AsyNq = new(nqDemo)

type nqDemo struct {
}

func (*nqDemo) _init() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     fmt.Sprintf(":%d", Cfg.GetVal("database.redis.port").Int()),
			Password: Cfg.GetVal("database.redis.pass").String(),
			DB:       Cfg.GetVal("database.redis.db").Int(),
		},
		asynq.Config{Concurrency: 20},
	)
	defer srv.Stop()

	mux := asynq.NewServeMux()

	// some middlewares
	mux.Use(func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			fmt.Println("receive")
			fmt.Println(fmt.Printf("[%s] log - %v", time.Now().Format("2006-01-02 15:04:05"), t.Payload))

			return next.ProcessTask(ctx, t)
		})
	})

	// some workers
	mux.HandleFunc("msg", func(ctx context.Context, task *asynq.Task) error {
		fmt.Println("------HandleMsg start------")
		fmt.Println(task)
		return nil
	})

	// start server
	if err := srv.Start(mux); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
