package jobsupervisor

type dummyJobSupervisor struct {
	status        string
	serviceStatus map[string]interface{}
}

func NewDummyJobSupervisor() JobSupervisor {
	return &dummyJobSupervisor{status: "unknown"}
}

func (s *dummyJobSupervisor) Reload() error {
	return nil
}

func (s *dummyJobSupervisor) Start() error {
	s.status = "running"
	s.serviceStatus = make(map[string]interface{})
	return nil
}

func (s *dummyJobSupervisor) Stop() error {
	s.status = "failing"
	return nil
}

func (s *dummyJobSupervisor) Unmonitor() error {
	return nil
}

func (s *dummyJobSupervisor) Status() (status string) {
	return s.status
}

func (s *dummyJobSupervisor) ServiceStatus() (serviceStatus map[string]interface{}) {
	return s.serviceStatus
}

func (s *dummyJobSupervisor) AddJob(jobName string, jobIndex int, configPath string) error {
	return nil
}

func (s *dummyJobSupervisor) RemoveAllJobs() error {
	return nil
}

func (s *dummyJobSupervisor) MonitorJobFailures(handler JobFailureHandler) error {
	return nil
}
