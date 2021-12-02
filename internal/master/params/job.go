package params

type CreateJobRequest struct {
	Name     string `json:"name" binding:"required"`      // 任务名
	Command  string `json:"command" binding:"required"`   // 任务命令行
	CronExpr string `json:"cron_expr" binding:"required"` // 执行周期
}

type CreateJobResponse struct {
	Name     string `json:"name"`      // 任务名
	Command  string `json:"command"`   // 任务命令行
	CronExpr string `json:"cron_expr"` // 执行周期
}
