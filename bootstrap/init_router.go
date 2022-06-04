package bootstrap

import (
	"context"
	"easy-admin/plugin/resp"
	"easy-admin/router/admin"
	"easy-admin/router/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Router = new(router)

type router struct {
}

// _init 路由初始化
func (r *router) _init() {

	// 监听端口
	listenAddr := Cfg.GetVal("server.port").Int()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", listenAddr),
		Handler: engine(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	fmt.Printf(`
[swagger:] http://127.0.0.1:%d/swagger/index.html
[server :] http://127.0.0.1:%d/info
[github :] https://github.com/smaltum/EasyAdmin.git
`,
		listenAddr, listenAddr)

	// 监听退出
	shutdown(server)
}

func engine() *gin.Engine {
	eng := gin.Default()

	maxMultipartMemory := Cfg.GetVal("server.maxMultipartMemory").Int64()

	//限制最大文件大小
	if maxMultipartMemory > 0 {
		eng.MaxMultipartMemory = maxMultipartMemory
	}

	// 路由
	eaGroup := eng.Group("")

	// 公开路由
	publicGroup := eaGroup.Group("")
	{
		publicGroup.Any("/info", func(c *gin.Context) {
			resp.New(c).ResponseOK()
		})
		// api
		api.RouterGroup.InitPublic(publicGroup)
	}

	// 私有路由
	privateGroup := eaGroup.Group("")
	{
		// admin
		admin.RouterGroup.Init(privateGroup)
		// api
		api.RouterGroup.InitPrivate(privateGroup)
	}

	return eng
}

// shutdown kill -9
// 监听标准信号
func shutdown(server *http.Server) {
	sdSign := make(chan os.Signal)
	signal.Notify(sdSign, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sign := <-sdSign
	log.Printf("ServerShutdown by == > %gvar", sign)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	//执行退出
	defer cancelFunc()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server.Shutdown() err %gvar", err)
	}
	log.Printf("ServerExit by == > %gvar", sign)
}
