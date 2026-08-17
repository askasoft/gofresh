package freshdesk

const (
	//JobStatusInProgress  = "IN PROGRESS"
	JobStatusInProgress = "in_progress"
	JobStatusCompleted  = "completed"
)

type Job struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Status string `json:"status,omitempty"`

	DownloadURL string `json:"download_url,omitempty"`

	CreatedAt Time `json:"created_at,omitzero"`

	UpdatedAt Time `json:"updated_at,omitzero"`

	StatusUpdatedAt Time `json:"status_updated_at,omitzero"`

	Progress int `json:"progress,omitempty"`
}

func (job *Job) IsCompleted() bool {
	return job.Status == JobStatusCompleted
}

func (job *Job) IsInProgress() bool {
	return job.Status == JobStatusInProgress
}

func (job *Job) String() string {
	return toString(job)
}
