package handlers

import (
	"golang.org/x/net/context"
	"github.com/andreasmaier/cimon_jobs/jobs"
)

func (s *JobsServer) CreateJob(ctx context.Context, in *jobs.CreateJobRequest) (*jobs.Job, error) {
	job := in.ToJenkinsJob()

	createdJob, err := jobs.CreateJobInDb(job)

	if err != nil {
		return nil, err
	}

	return &jobs.Job{
		Id: createdJob.Id,
		Path: createdJob.Path,
		Status: createdJob.Status,
		Alias: createdJob.Alias,
	}, nil
}