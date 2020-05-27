package model

import (
	"fmt"
	"time"
)

// RecordStatus Status of one record
type RecordStatus int

const (
	// Done Represent the record is finished.
	Done RecordStatus = iota
	// Doing Represent the recod is still on progress.
	Doing
)

// Record A record of doing things.
type Record struct {
	Status      RecordStatus
	Detail      string
	CreatedTime time.Time
}

func (record *Record) toTaskPaper() string {
	result := fmt.Sprintf("  - %s@created(%s)", record.Detail, record.CreatedTime)
	if record.Status == Done {
		result += "@done"
	}
	return result
}
