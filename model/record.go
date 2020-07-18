package model

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// RecordStatus Status of one record
type RecordStatus int

// RecordTag Tag of one record, contains tag name and tag value.
type RecordTag map[string]string

// NewRecordTag Create a new RecordTag instance.
func NewRecordTag() RecordTag {
	return make(RecordTag)
}

// AddTag Add a new tag.
func (tag *RecordTag) AddTag(name string, value string) {
	(*tag)[name] = value
}

const (
	// Done Represent the record is finished.
	Done RecordStatus = iota
	// Doing Represent the recod is still on progress.
	Doing
)

// Record A record of doing things.
type Record struct {
	Status      RecordStatus
	Title       string
	CreatedTime time.Time
	Tag         RecordTag
	Detail      string
}

// NewDoingRecord Factory method for creating one new doing record.
func NewDoingRecord(title string) Record {
	return Record{
		Status:      Doing,
		Title:       title,
		CreatedTime: time.Now(),
		Tag:         NewRecordTag(),
		Detail:      "",
	}
}

// NewDoneRecord Factory method for creating one new done record.
func NewDoneRecord(title string) Record {
	return Record{
		Status:      Done,
		Title:       title,
		CreatedTime: time.Now(),
		Tag:         NewRecordTag(),
		Detail:      "",
	}
}

// ToTaskPaper Convert the record to a string that is compatible to
// task paper syntax.
func (record *Record) ToTaskPaper() string {
	result := fmt.Sprintf("  - %s @created(%s)", record.Title, record.CreatedTime.Format(time.RFC3339))
	if record.Status == Done {
		result += " @done"
	}
	for tag, value := range record.Tag {
		if value == "" {
			result += fmt.Sprintf(" @%s", tag)
		} else {
			result += fmt.Sprintf(" @%s(%s)", tag, value)
		}
	}
	if record.Detail != "" {
		result += fmt.Sprintln()
		result += fmt.Sprintf(record.Detail)
	}
	return result
}

func (record *Record) AddDetail(detail string) {
	record.Detail = detail
}

// AddTag Add a tag to the provided record.
func (record *Record) AddTag(name string, value string) {
	record.Tag.AddTag(name, value)
}

// AddTagFromRawString Based on provided string to genearte a tag.
// The provided text should be either in format "tag=value" or "tag".
func (record *Record) AddTagFromRawString(text string) error {
	nameAndValue := strings.Split(text, "=")
	if len(nameAndValue) == 1 {
		record.AddTag(nameAndValue[0], "")
	} else if len(nameAndValue) == 2 {
		record.AddTag(nameAndValue[0], nameAndValue[1])
	} else {
		return fmt.Errorf("Failed to find tag information for %s", text)
	}
	return nil
}

// AddTagFromRawStringList Based on a list of string to generate tags.
func (record *Record) AddTagFromRawStringList(text []string) error {
	for _, t := range text {
		err := record.AddTagFromRawString(t)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveTag Remove a tag from the tag name.
func (record *Record) RemoveTag(name string) {
	_, ok := record.Tag[name]
	if !ok {
		return
	}
	delete(record.Tag, name)
	return
}

//RemoveTagFromRawString Remove a tag based on provided text.
// The provided text should be either in format "tag=value" or "tag".
func (record *Record) RemoveTagFromRawString(text string) {
	nameAndValue := strings.Split(text, "=")
	name := nameAndValue[0]
	record.RemoveTag(name)
}

// RemoveTagsFromRawStringList Remove tags based on a string list.
func (record *Record) RemoveTagsFromRawStringList(text []string) {
	for _, t := range text {
		record.RemoveTagFromRawString(t)
	}
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
		Tag:    make(RecordTag),
	}
	lines := strings.Split(text, "\n")
	firstLine := lines[0]
	regex := regexp.MustCompile(`((@[^()\s]+(\([^()]+\))?\s*))+$`)
	find := regex.FindIndex([]byte(firstLine))
	if find == nil {
		return record, fmt.Errorf("Failed to get tag of provided record: %s", text)
	}
	title := firstLine[0 : find[0]-1] //Remove the trailing empty space
	record.Title = title
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
		} else {
			record.AddTag(tag, value)
		}
	}
	if len(lines) > 0 {
		details := strings.Join(lines[1:], "\n")
		record.AddDetail(details)
	}
	return record, nil
}

// Done Mark a record is done.
func (record *Record) Done() {
	record.Status = Done
}

// ToDisplayString Conver one record to a string to be displayed.
func (record *Record) ToDisplayString() string {
	timeString := record.CreatedTime.Format("2006-01-02 15:04")
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(fmt.Sprintf("| %s | %s", timeString, record.Title))
	if record.Detail != "" {
		stringBuilder.WriteString(fmt.Sprintf("\n%s", record.Detail))
	}
	return stringBuilder.String()
}
