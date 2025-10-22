package models

type Alert struct {
	Level   string  `json:"level"` // danger atau warning
	Message string  `json:"message"`
	Sensor  string  `json:"sensor"`
	Value   string  `json:"value"` // Nilai sensor yang memicu
	Time    string  `json:"time"`  // Waktu deteksi
	Timestamp int64 `json:"timestamp"`
}