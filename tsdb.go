package arrowdb

import (
	"encoding/json"
	"time"
)

const (
	defaultRetentionTime = 14 * 24 * time.Hour // 2 weeks
	defaultBlockDuration = 2 * time.Hour
	defaultWalBufferSize = 16 * 1024
	defaultDataPath      = "./data"
)

type DBOptions struct {
	Retention     time.Duration // time series data retention time
	BlockDuration time.Duration // disk block file duration time
	WalBufferSize int           // batch size to flush to disk
	DataPath      string        // path to save time series data
}

var DefaultDBOptions = &DBOptions{
	Retention:     defaultRetentionTime,
	BlockDuration: defaultBlockDuration,
	WalBufferSize: defaultWalBufferSize,
	DataPath:      defaultDataPath,
}

type DB struct {
	options *DBOptions
}

// OpenTSDB open a new time series db with options, it will use DefaultDBOptions if options is nil
func OpenTSDB(options *DBOptions) (*DB, error) {
	if options == nil {
		options = DefaultDBOptions
	}

	db := &DB{
		options: options,
	}

	db.loadDiskBlocks()
	db.replayWAL()

	go db.backgroundWorker()
	return db, nil
}

func (db *DB) loadDiskBlocks() {

}

func (db *DB) replayWAL() {

}

func (db *DB) backgroundWorker() {

}

func (db *DB) InsertJsonRows(rowsData []byte) error {
	var rows []Row
	err := json.Unmarshal(rowsData, &rows)
	if err != nil {
		return err
	}

	return db.InsertRows(rows)
}

func (db *DB) InsertRows(rows []Row) error {
	return nil
}

// QueryMetrics query metrics that match input pattern with regexp, use ".*" to query all metrics in the db
func (db *DB) QueryMetrics(metricPattern string) ([]string, error) {
	return nil, nil
}

func (db *DB) QueryTagNames(metric string) ([]string, error) {
	return nil, nil
}

func (db *DB) QueryTagValues(metric string, tms TagMatcherSet) ([]string, error) {
	return nil, nil
}

func (db *DB) QuerySeries(metric string, tms TagMatcherSet) ([]map[string]string, error) {
	return nil, nil
}

func (db *DB) QueryRange(start int64, end int64, metric string, tms TagMatcherSet) (*SeriesDataPoints, error) {
	return nil, nil
}

func (db *DB) QueryTopN(start int64, end int64, metric string, groupByName string, tms TagMatcherSet) ([]TopNData, error) {
	return nil, nil
}
