package model

// RecordID Data structure represents the Id for one record
type RecordID struct {
	ID int
}

// GetNextID Return the next ID to be used.
func (recordId RecordID) GetNextID() RecordID {
	return RecordID{ID: recordId.ID + 1}
}

// RecordList A list of doing records.
type RecordList struct {
	doneRecords  map[RecordID]Record
	doingRecords map[RecordID]Record
	nextID       RecordID
}

}
