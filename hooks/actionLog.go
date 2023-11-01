package hooks

import (
	"github.com/captainhook-go/captainhook/config"
	"github.com/captainhook-go/captainhook/io"
)

type ActionLog struct {
	log []*ActionLogItem
}

func NewActionLog() *ActionLog {
	return &ActionLog{}
}

func (a *ActionLog) Add(item *ActionLogItem) {
	a.log = append(a.log, item)
}

func (a *ActionLog) Logs() []*ActionLogItem {
	return a.log
}

func (a *ActionLog) HasLogs() bool {
	for _, l := range a.log {
		if l.CollectorIO.HasCollectedMessages() {
			return true
		}
	}
	return false
}

type ActionLogItem struct {
	Conf        *config.Action
	CollectorIO *io.CollectorIO
	Status      int
}

func NewActionLogItem(action *config.Action, collectorIO *io.CollectorIO, status int) *ActionLogItem {
	return &ActionLogItem{action, collectorIO, status}
}
