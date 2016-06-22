package jobs

type JenkinsJob struct {
	Id int32 `json:"id"`
	Server string `json:"server"`
	Path string `json:"path"`
	Status string `json:"status"`
	Alias string `json:"alias"`
}

func (r *CreateJobRequest) ToJenkinsJob() (*JenkinsJob) {
	return &JenkinsJob{
		Path: r.Path,
		Status: r.Status,
		Alias: r.Alias,
	}
}

func (jj *JenkinsJob) ToJobMessage() (*Job) {
	return &Job{
		Id: jj.Id,
		Path: jj.Path,
		Alias: jj.Alias,
		Status: jj.Status,
	}
}