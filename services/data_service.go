package services

import (
	"github.com/dawwasinha/ergovest-backend/models" // Ganti dengan path modul Anda
	"sync"
)

// Global Variables untuk penyimpanan In-Memory (Simulasi Database)
var (
	// SensorHistory menyimpan semua data sensor yang masuk dari MQTT
	SensorHistory = make([]models.SensorData, 0)
	// AlertHistory menyimpan semua alert yang terdeteksi
	AlertHistory  = make([]models.Alert, 0) 
	
	// Mutex untuk concurrency safety karena ini diakses dari MQTT Goroutine dan API Controller
	dataMutex = sync.RWMutex{}
)

// AddSensorData menyimpan data sensor baru ke histori. Dipanggil dari MQTT Service.
func AddSensorData(data models.SensorData) {
	dataMutex.Lock()
	defer dataMutex.Unlock()
	
	// Batasi ukuran histori agar memori tidak penuh (misalnya, 1000 data)
	if len(SensorHistory) >= 1000 {
		// Hapus data tertua
		SensorHistory = SensorHistory[1:] 
	}
	SensorHistory = append(SensorHistory, data)
}

// GetSensorHistory mendapatkan histori data sensor (Contoh: 100 data terbaru)
func GetSensorHistory(limit int) []models.SensorData {
	dataMutex.RLock()
	defer dataMutex.RUnlock()
	
	if limit > len(SensorHistory) {
		limit = len(SensorHistory)
	}
	// Mengembalikan N data terbaru
	return SensorHistory[len(SensorHistory)-limit:]
}

// AddAlert menyimpan alert yang terdeteksi. Dipanggil dari Alert Service.
func AddAlert(alert models.Alert) {
	dataMutex.Lock()
	defer dataMutex.Unlock()
	
	if len(AlertHistory) >= 50 { // Batasi riwayat alert
		AlertHistory = AlertHistory[1:]
	}
	AlertHistory = append(AlertHistory, alert)
}

// GetAlertHistory mendapatkan histori alert
func GetAlertHistory() []models.Alert {
	dataMutex.RLock()
	defer dataMutex.RUnlock()
	
	return AlertHistory
}

// TODO: Tambahkan fungsi untuk SurveyHistory di sini (AddSurvey, GetSurveyStats, dll.)