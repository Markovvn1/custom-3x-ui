package job

import (
	"x-ui/web/service"
)

type StatsNotifyJob struct {
	xrayService  service.XrayService
	tgbotService service.Tgbot
}

func NewStatsNotifyJob() *StatsNotifyJob {
	return new(StatsNotifyJob)
}

// Here run is a interface method of Job interface
func (j *StatsNotifyJob) Run() {
	if !j.xrayService.IsXrayRunning() {
		return
	}
	j.tgbotService.SendReport()
}
