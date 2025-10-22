package services

import (
	"fmt"
	"math"
	"time"
	"github.com/dawwasinha/ergovest-backend/models" // Ganti dengan path modul Anda
)

const (
	PITCH_THRESHOLD = 30.0  // Pitch1/Pitch2 > 30 deg
	ROLL_THRESHOLD  = 30.0  // Roll1/Roll2 > 30 deg
	YAW_THRESHOLD   = 100.0 // YawSpeed > 100 deg/s
	MUSCLE_THRESHOLD = 3000 // Muscle > 3000
	TEMP_THRESHOLD   = 35.0 // Temp > 35 C
)

// CheckAndStoreAlert memeriksa data sensor terhadap ambang batas dan menyimpan alert
func CheckAndStoreAlert(data models.SensorData) {
	alerts := make([]models.Alert, 0)
	
	currentTime := time.Now().Format("02-01-2006 15:04:05")

	// Helper untuk membuat dan menambahkan alert
	addAlert := func(level, message, sensor string, value float32, unit string) {
		alerts = append(alerts, models.Alert{
			Level:   level,
			Message: message,
			Sensor:  sensor,
			Value:   fmt.Sprintf("%.1f%s", value, unit),
			Time:    currentTime,
			Timestamp: time.Now().UnixMilli(),
		})
	}

	// 1. MPU1 (Upper Back) - Pitch dan Roll
	if math.Abs(float64(data.Pitch1)) > PITCH_THRESHOLD {
		addAlert("danger", "🚨 DANGER: Extreme upper back bending!", "MPU1 Pitch", data.Pitch1, "°")
	}
	if math.Abs(float64(data.Roll1)) > ROLL_THRESHOLD {
		addAlert("danger", "🚨 DANGER: Extreme upper back side bending!", "MPU1 Roll", data.Roll1, "°")
	}
	if math.Abs(float64(data.YawSpeed1)) > YAW_THRESHOLD {
		addAlert("warning", "⚠️ WARNING: Rapid upper back twisting!", "MPU1 Yaw", data.YawSpeed1, "°/s")
	}

	// 2. MPU2 (Lower Back) - Pitch dan Roll
	if math.Abs(float64(data.Pitch2)) > PITCH_THRESHOLD {
		addAlert("danger", "🚨 DANGER: Extreme lower back bending!", "MPU2 Pitch", data.Pitch2, "°")
	}
	if math.Abs(float64(data.Roll2)) > ROLL_THRESHOLD {
		addAlert("danger", "🚨 DANGER: Extreme lower back side bending!", "MPU2 Roll", data.Roll2, "°")
	}
	if math.Abs(float64(data.YawSpeed2)) > YAW_THRESHOLD {
		addAlert("warning", "⚠️ WARNING: Rapid lower back twisting!", "MPU2 Yaw", data.YawSpeed2, "°/s")
	}

	// 3. EMG (Muscle)
	if data.Muscle > MUSCLE_THRESHOLD {
		addAlert("warning", "⚠️ WARNING: High muscle strain!", "EMG", float32(data.Muscle), "")
	}
	
	// 4. BME (Temperature)
	if data.Temp > TEMP_THRESHOLD {
		addAlert("warning", "⚠️ WARNING: High temperature!", "Temperature", data.Temp, "°C")
	}
	
	// Simpan semua alert yang terdeteksi
	for _, alert := range alerts {
		AddAlert(alert)
		// TODO: Tambahkan logic untuk WebSockets broadcast di sini
	}
}