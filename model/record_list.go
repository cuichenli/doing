package model

// RecordID Data structure represents the Id for one record
type RecordID struct {
	ID int
}

// GetNextID Return the next ID to be used.
func (recordId RecordID) GetNextID() RecordID {
	return RecordID{ID: recordId.ID + 1}
}

// GetHashValue Get the record ID's hash key value
func (recordId RecordID) GetHashValue() int {
	return recordId.ID
}

// RecordList A list of doing records.
type RecordList struct {
	doneRecords  map[int]Record
	doingRecords map[int]Record
	nextID       RecordID
}

// AddRecord Add one record to the record list.
func (recordList *RecordList) AddRecord(record Record) {
	recordList.doingRecords[recordList.nextID.GetHashValue()] = record
	recordList.nextID = recordList.nextID.GetNextID()
}
