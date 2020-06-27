package model

import (
	"fmt"
	"regexp"
	"strings"
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

// NewDoingRecord Factory method for creating one new doing record.
func NewDoingRecord(detail string) Record {
	return Record{
		Status:      Doing,
		Detail:      detail,
		CreatedTime: time.Now(),
	}
}

// NewDoneRecord Factory method for creating one new done record.
func NewDoneRecord(detail string) Record {
	return Record{
		Status:      Done,
		Detail:      detail,
		CreatedTime: time.Now(),
	}
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

// ParseTags Parse provided tag information and return tag's name and value
func ParseTags(text string) (string, string, error) {
	trimtedText := strings.Trim(text, " ")
	regex := regexp.MustCompile(`^([^()\s]+)(\(([^()]+)\))?$`)
	groups := regex.FindStringSubmatch(trimtedText)
	groupsNum := len(groups)
	if groupsNum == 1 {
		return groups[0], "", nil
	}
	if groupsNum == 4 {
		return groups[1], groups[3], nil
	}
	return "", "", fmt.Errorf("Failed to find tag information for %s", text)
}

// FromTaskPaper Generate Record based on the provided TaskPaper string.
func FromTaskPaper(text string) (Record, error) {
	record := Record{
		Status: Doing,
	}
	regex := regexp.MustCompile(`((@[^()\s]+(\([^()]+\))?\s*))+$`)
	find := regex.FindIndex([]byte(text))
	if find == nil {
		return record, fmt.Errorf("Failed to get tag of provided record: %s", text)
	}
	detail := text[0 : find[0]-1] //Remove the trailing empty space
	record.Detail = detail
	tagsText := text[find[0]+1 : find[1]] // Skip the first "@"
	tagTexts := strings.Split(tagsText, "@")
	for _, tagText := range tagTexts {
		tag, value, err := ParseTags(tagText)
		if err != nil {
			return record, err
		}
		if tag == "done" {
			record.Done()
			continue
		}
		if tag == "created" {
			t, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return record, fmt.Errorf("Failed to parse time %s for item %s", value, text)
			}
			record.CreatedTime = t
			continue
		}
	}
	return record, nil
}

// Done Mark a record is done.
func (record *Record) Done() {
	record.Status = Done
}
