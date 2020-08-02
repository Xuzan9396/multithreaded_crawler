package crontab

import (
	"errors"
	"fmt"
	"github.com/gorhill/cronexpr"
	"pacong/zhenai/engine"
	"pacong/zhenai/parser"
	"time"
)

// 任务调度计划表
type JobSchedulerPlan struct {
	Job      *Job
	Expr     *cronexpr.Expression // 解析好的cronnxpr 表达式
	NextTime time.Time
}

type Job struct {
	Name     string `json:"name"`     // 任务名
	Command  string `json:"command"`  // shell 命令
	CronExpr string `json:"cronExpr"` // cron 表达式
	Url      string `json:"url"`
	PushWx   []string
}

type Scheduler struct {
	//jobEventChan      chan *JobEvent
	//jobConfigEventChan chan int
	jobPlanTable map[string]*JobSchedulerPlan // 执行计划表
}

func InitCrontab() *Scheduler {
	model := &Scheduler{
		jobPlanTable: make(map[string]*JobSchedulerPlan),
	}
	//jobs := make([]*Job,0)

	jobs := []Job{
		{
			Name:     "weibo",
			Command:  "1",
			CronExpr: "*/5 * * * * * *",
			Url:      "https://s.weibo.com/top/summary",
			PushWx:   []string{"SCU64514T1d2bceaaf7386be63d2e6da3d22e46995da98d09d8ca7"},
		},

		{
			Name:     "agu",
			Command:  "1",
			CronExpr: "*/30 * * * * * *",
			Url:      "https://cn.investing.com/indices/shanghai-se-a-share",
			PushWx:   []string{"SCU64514T1d2bceaaf7386be63d2e6da3d22e46995da98d09d8ca7"},
		},
	}

	for _, job := range jobs {
		model.jobPlanTable[job.Name], _ = BuildSchedulerPlan(&job)

	}
	return model
}

// 构建任务执行计划
func BuildSchedulerPlan(job *Job) (jobSchedulerPlan *JobSchedulerPlan, err error) {
	var (
		expr *cronexpr.Expression
	)
	if expr, err = cronexpr.Parse(job.CronExpr); err != nil {
		fmt.Println(err, "解析错误了")
		return
	}
	nextNow := expr.Next(time.Now())
	now := time.Now()
	if nextNow.Before(now) {
		err = errors.New("时间过期了")
		return
	}

	jobSchedulerPlan = &JobSchedulerPlan{
		Job:      job,
		Expr:     expr,
		NextTime: nextNow,
	}
	return

}

// 调度协程
func (scheduler *Scheduler) SchedulerLoop() {
	// 定时任务
	var (
		schedulerAfter time.Duration
		schedulerTimer *time.Timer
	)

	// 计算调度的时间
	schedulerAfter = scheduler.TrySchedule()
	// 调度延时器
	schedulerTimer = time.NewTimer(schedulerAfter)
	// 调度延迟事件

	for {
		select {
		case <-schedulerTimer.C:

		}

		// 重新调度一次任务
		schedulerAfter = scheduler.TrySchedule()
		// 重置任务定时器
		schedulerTimer.Reset(schedulerAfter)
	}
}

// 尝试遍历所有任务
func (scheduler *Scheduler) TrySchedule() (schedulerAfter time.Duration) {
	var (
		//jobPlan  *JobSchedulerPlan
		now      time.Time
		nearTime *time.Time
	)

	// 没有任务睡一s
	if len(scheduler.jobPlanTable) == 0 {
		schedulerAfter = 1 * time.Second
		return
	}

	now = time.Now()
	for key, jobPlan := range scheduler.jobPlanTable {

		if jobPlan.NextTime.Unix() < 0 {
			// 过期的删除
			fmt.Println("我删除了%s", jobPlan.Job.Name)
			delete(scheduler.jobPlanTable, key)
			continue
		}
		if jobPlan.NextTime.Before(now) || jobPlan.NextTime.Equal(now) {

			//scheduler.realScheduleJobPlan(jobPlan)
			//fmt.Println("执行");
			var parserFunc func([]byte) engine.ParseResult
			switch jobPlan.Job.Name {
			case "weibo":
				parserFunc = func(bytes []byte) engine.ParseResult {
					return parser.WeiboList(bytes, jobPlan.Job.Url)
				}
			case "agu":
				parserFunc = func(bytes []byte) engine.ParseResult {
					return parser.Aparser(bytes, jobPlan.Job.Url)
				}
			default:

			}

			engine.SetRequestChan(engine.Request{
				Url:        jobPlan.Job.Url,
				ParserFunc: parserFunc,
			})

			//log.Printf("执行")
			// 更新下次执行时间
			jobPlan.NextTime = jobPlan.Expr.Next(now)

		}

		if nearTime == nil || jobPlan.NextTime.Before(*nearTime) {
			nearTime = &jobPlan.NextTime
		}
	}

	// 睡眠多少时间
	schedulerAfter = (*nearTime).Sub(now)

	return
}
