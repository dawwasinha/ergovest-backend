package services

import (
	"testing"
	"time"

	"github.com/dawwasinha/ergovest-backend/models"
)

func TestCheckAndStoreAlert(t *testing.T) {
	// Initialize DB for persistence used by AddAlert
	InitDB()
	defer CloseDB()

	// reset history
	dataMutex.Lock()
	AlertHistory = make([]models.Alert, 0)
	dataMutex.Unlock()

	sd := models.SensorData{
		Pitch1:    45.0,
		Roll1:     40.0,
		YawSpeed1: 200.0,
		Pitch2:    50.0,
		Roll2:     35.0,
		YawSpeed2: 150.0,
		Muscle:    4000,
		Temp:      36.0,
		Timestamp: time.Now().UnixMilli(),
	}

	CheckAndStoreAlert(sd)

	// Wait briefly for goroutines (if any) to finish
	time.Sleep(100 * time.Millisecond)

	alerts := GetAlertHistory()
	if len(alerts) == 0 {
		t.Fatalf("expected alerts to be generated, got 0")
	}
}
