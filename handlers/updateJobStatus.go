package handlers

import (
	"github.com/andreasmaier/cimon_jobs/jobs"
	"golang.org/x/net/context"
)

func (s *JobsServer) UpdateJobStatus(ctx context.Context, in *jobs.UpdateStatusRequest) (*jobs.Empty, error) {
	err := jobs.UpdateJobInDb(in.Path, in.Status)

	return &jobs.Empty{}, err
}
