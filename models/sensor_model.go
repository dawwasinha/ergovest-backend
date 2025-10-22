package models

// SensorData merepresentasikan payload JSON dari ESP32 Receiver
type SensorData struct {
    Temp      float32 `json:"temp"`
    Hum       float32 `json:"hum"`
    Gas       float32 `json:"gas"`
    Muscle    int     `json:"muscle"`
    Piezo     int     `json:"piezo"`
    Pitch1    float32 `json:"pitch1"`
    Roll1     float32 `json:"roll1"`
    YawSpeed1 float32 `json:"yawSpeed1"`
    Pitch2    float32 `json:"pitch2"`
    Roll2     float32 `json:"roll2"`
    YawSpeed2 float32 `json:"yawSpeed2"`
    Timestamp int64   `json:"timestamp"` // epoch time
}