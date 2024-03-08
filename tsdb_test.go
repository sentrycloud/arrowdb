package arrowdb

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func prepareInsertRows() []Row {
	ts := time.Now().Unix()
	rows := []Row{
		{Metric: "sys.cpu", Tags: map[string]string{"ip": "127.0.0.1"}, Timestamp: ts, Value: 10},
		{Metric: "sys.load", Tags: map[string]string{"ip": "127.0.0.1"}, Timestamp: ts, Value: 2},
	}

	return rows
}

func TestDB_InsertJsonRows(t *testing.T) {
	rows := prepareInsertRows()
	rowsData, err := json.Marshal(rows)
	if err != nil {
		fmt.Printf("json encode failed: %v\n", err)
		return
	}

	db, err := OpenTSDB(nil)
	if err != nil {
		fmt.Printf("open tsdb failed: %v\n", err)
		return
	}

	err = db.InsertJsonRows(rowsData)
	if err != nil {
		fmt.Printf("tsdb insert json rows failed: %v\n", err)
		return
	}

	fmt.Printf("tsdb insert json rows success\n")
}

func TestDB_InsertRows(t *testing.T) {
	rows := prepareInsertRows()

	db, err := OpenTSDB(nil)
	if err != nil {
		fmt.Printf("open tsdb failed: %v\n", err)
		return
	}

	err = db.InsertRows(rows)
	if err != nil {
		fmt.Printf("tsdb insert rows failed: %v\n", err)
		return
	}

	fmt.Printf("tsdb insert rows success\n")
}
