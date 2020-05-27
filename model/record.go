package model

import "time"

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
	status      RecordStatus
	detail      string
	createdTime time.Time
}
