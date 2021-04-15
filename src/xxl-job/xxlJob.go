package xxl_job

import (
	"fmt"
	"github.com/gin-gonic/gin"
	xxl_job_executor_gin "github.com/gin-middleware/xxl-job-executor"
	"github.com/xxl-job/xxl-job-executor-go"
	"golang.org/x/net/context"
	"log"
)

func Job() {
	exec := xxl.NewExecutor(
		xxl.ServerAddr("http://120.79.223.58:8010/xxl-job-admin"),
		xxl.AccessToken(""),
		xxl.RegistryKey("xxl-job-executor-sample"), //执行器名称
		xxl.SetLogger(&logger{}),                   //自定义日志
	)
	exec.Init()

	//设置日志查看handler
	exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
			FromLineNum: req.FromLineNum,
			ToLineNum:   2,
			LogContent:  "这个是自定义日志handler",
			IsEnd:       true,
		}}
	})

	//添加到gin路由
	r := gin.Default()
	xxl_job_executor_gin.XxlJobMux(r, exec)

	//注册gin的handler
	r.GET("ping", func(cxt *gin.Context) {
		cxt.JSON(200, "pong")
	})

	//注册任务handler
	exec.RegTask("Task_Test", Task_Test)
	log.Fatal(r.Run(":9999"))
}

//xxl.Logger接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}

func Task_Test(cxt context.Context, param *xxl.RunReq) (msg string) {
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "test done"
}
