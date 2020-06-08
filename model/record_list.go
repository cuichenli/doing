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

// InitializeRecordID Create a new recordID.
func InitializeRecordID() RecordID {
	return RecordID{ID: 0}
}

// RecordList A list of doing records.
type RecordList struct {
	DoneRecords  map[int]Record
	DoingRecords map[int]Record
	NextID       RecordID
}

// NewRecordList Create a new RecordList instance.
func NewRecordList() RecordList {
	return RecordList{
		DoneRecords:  make(map[int]Record),
		DoingRecords: make(map[int]Record),
		NextID:       InitializeRecordID(),
	}
}

// AddRecord Add one record to the record list.
func (recordList *RecordList) AddRecord(record Record) {
	recordList.DoingRecords[recordList.NextID.GetHashValue()] = record
	recordList.NextID = recordList.NextID.GetNextID()
}

// MarkItemDone Mark one item is done.
func (recordList *RecordList) MarkItemDone(recordID RecordID) (*Record, error) {
	hashKey := recordID.GetHashValue()
	targetRecord, ok := recordList.DoingRecords[hashKey]
	if !ok {
		return nil, errors.New("The provided record ID is not found in doing record list")
	}
	targetRecord.Done()
	recordList.DoneRecords[hashKey] = targetRecord
	delete(recordList.DoingRecords, hashKey)
	return &targetRecord, nil
}
