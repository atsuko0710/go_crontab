package v1

import (
	"encoding/json"
	"master/internal/master/params"
	"master/internal/master/store/etcd"
)

const (
	JOB_SAVE_DIR = "/cron/jobs/"  // 任务保存目录
)

func CreateJob(job *params.CreateJobRequest) (oldJob *params.CreateJobResponse, err error) {
	jobKey := getJobKey(job.Name)

	jobValue, err := json.Marshal(job)
	if err != nil {
		return
	}
	putResp, err := etcd.Put(jobKey, string(jobValue))
	if err != nil {
		return
	}

	if putResp.PrevKv != nil {
		if err = json.Unmarshal(putResp.PrevKv.Value, &oldJob); err != nil {
			return
		}
	}
	return
}

func getJobKey(key string) string {
	return JOB_SAVE_DIR + key
}