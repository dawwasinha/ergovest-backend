package services

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/dawwasinha/ergovest-backend/models"
	"go.etcd.io/bbolt"
)

const (
	DB_FILE       = "ergovest.db"
	BUCKET_SENSOR = "sensor_history"
	BUCKET_ALERT  = "alert_history"
)

var db *bbolt.DB

func InitDB() {
	var err error
	db, err = bbolt.Open(DB_FILE, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatalf("failed to open bbolt DB: %v", err)
	}

	// Ensure buckets exist
	err = db.Update(func(tx *bbolt.Tx) error {
		if _, e := tx.CreateBucketIfNotExists([]byte(BUCKET_SENSOR)); e != nil {
			return e
		}
		if _, e := tx.CreateBucketIfNotExists([]byte(BUCKET_ALERT)); e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to create buckets: %v", err)
	}

	// Load last N entries into memory to initialize histories
	loadRecentSensors(200)
	loadRecentAlerts(100)
}

func CloseDB() {
	if db != nil {
		_ = db.Close()
	}
}

func SaveSensor(data models.SensorData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(BUCKET_SENSOR))
		// use timestamp as key (ms)
		key := []byte(fmtInt64(data.Timestamp))
		return bkt.Put(key, b)
	})
}

func SaveAlert(a models.Alert) error {
	b, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return db.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(BUCKET_ALERT))
		key := []byte(fmtInt64(a.Timestamp))
		return bkt.Put(key, b)
	})
}

func loadRecentSensors(limit int) {
	// load all keys and return last `limit` into SensorHistory
	tmp := make([]models.SensorData, 0)
	_ = db.View(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(BUCKET_SENSOR))
		if bkt == nil {
			return nil
		}
		c := bkt.Cursor()
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			var s models.SensorData
			if err := json.Unmarshal(v, &s); err == nil {
				tmp = append(tmp, s)
				if len(tmp) >= limit {
					break
				}
			}
		}
		return nil
	})

	// reverse order to oldest->newest
	for i := len(tmp) - 1; i >= 0; i-- {
		AddSensorData(tmp[i])
	}
}

func loadRecentAlerts(limit int) {
	tmp := make([]models.Alert, 0)
	_ = db.View(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte(BUCKET_ALERT))
		if bkt == nil {
			return nil
		}
		c := bkt.Cursor()
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			var a models.Alert
			if err := json.Unmarshal(v, &a); err == nil {
				tmp = append(tmp, a)
				if len(tmp) >= limit {
					break
				}
			}
		}
		return nil
	})
	for i := len(tmp) - 1; i >= 0; i-- {
		AddAlert(tmp[i])
	}
}

// helper to convert int64 to string bytes (no import fmt to keep minimal)
func fmtInt64(v int64) string {
	return strconv.FormatInt(v, 10)
}
