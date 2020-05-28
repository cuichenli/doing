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

// ToTaskPaper Convert the record to a string that is compatible to
// task paper syntax.
func (record *Record) ToTaskPaper() string {
	result := fmt.Sprintf("  - %s @created(%s)", record.Detail, record.CreatedTime.Format(time.RFC3339))
	if record.Status == Done {
		result += " @done"
	}
	return result
}

// Done Mark a record is done.
func (record *Record) Done() {
	record.Status = Done
}
