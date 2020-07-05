package model

// RecordList A list of doing records.
type RecordList struct {
	Records []Record
}

// NewRecordList Create a new RecordList instance.
func NewRecordList() RecordList {
	return RecordList{
		Records: make([]Record, 0),
	}
}

// AddRecord Add one record to the record list.
func (recordList *RecordList) AddRecord(record Record) {
	recordList.Records = append(recordList.Records, record)
}

// GetAllRecords Get a reference to all the records.
func (recordList *RecordList) GetAllRecords() *[]Record {
	return &recordList.Records
}

// EditRecords Edit the provided records based on the handler.
func (recordList *RecordList) EditRecords(start int, end int, handler func(*Record) error) error {
	records := recordList.GetAllRecords()
	for i := start; i < end; i++ {
		err := handler(&(*records)[i])
		if err != nil {
			return err
		}
	}
	return nil
}
