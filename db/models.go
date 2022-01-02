package db

import (
	"strconv"
	"time"
)

type DomainStatus struct {
	Available    bool
	LastSearched time.Time
}

func (s DomainStatus) ToHashMap() map[string]string {
	return map[string]string{
		"available":     strconv.FormatBool(s.Available),
		"last_searched": s.LastSearched.Format("2006-01-02 15:04:05"),
	}
}

func NewDomainStatus(available bool, time time.Time) DomainStatus {
	return DomainStatus{Available: available, LastSearched: time}
}

func NewDomainStatusFromHashMap(h map[string]string) (status DomainStatus, err error) {
	status.Available, err = strconv.ParseBool(h["available"])
	if err != nil {
		return
	}

	status.LastSearched, err = time.Parse("2006-01-02 15:04:05", h["last_searched"])
	if err != nil {
		return
	}
	return
}
