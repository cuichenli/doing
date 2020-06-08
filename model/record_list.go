package model

import "errors"

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

// NewRecordList Create a new RecordList instance.
func NewRecordList() RecordList {
	return RecordList{
		doneRecords:  make(map[int]Record),
		doingRecords: make(map[int]Record),
		nextID:       RecordID{ID: 0},
	}
}

// AddRecord Add one record to the record list.
func (recordList *RecordList) AddRecord(record Record) {
	recordList.doingRecords[recordList.nextID.GetHashValue()] = record
	recordList.nextID = recordList.nextID.GetNextID()
}

// MarkItemDone Mark one item is done.
func (recordList *RecordList) MarkItemDone(recordID RecordID) (*Record, error) {
	hashKey := recordID.GetHashValue()
	targetRecord, ok := recordList.doingRecords[hashKey]
	if !ok {
		return nil, errors.New("The provided record ID is not found in doing record list")
	}
	targetRecord.Done()
	recordList.doneRecords[hashKey] = targetRecord
	delete(recordList.doingRecords, hashKey)
	return &targetRecord, nil
}
