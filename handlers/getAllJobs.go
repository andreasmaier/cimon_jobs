package handlers

import (
	"github.com/andreasmaier/cimon_jobs/jobs"
	"golang.org/x/net/context"
)

func (s *JobsServer) GetAllJobs(ctx context.Context, in *jobs.Empty) (*jobs.Jobs, error) {
	jjobs, err := jobs.GetAllJobsFromDb()

	if err != nil {
		panic(err)
	}

	result := make([]*jobs.Job, len(jjobs))

	for idx, jj := range jjobs {
		result[idx] = jj.ToJobMessage()
	}

	return &(jobs.Jobs{Jobs: result}), nil
}
