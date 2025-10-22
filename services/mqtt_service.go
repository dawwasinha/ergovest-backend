package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/dawwasinha/ergovest-backend/models" // Ganti dengan path modul Anda

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// PENTING: Ganti dengan kredensial dari config.h
const (
	MQTT_BROKER = "ssl://1e92e59ecc3942869a04a67fe409ba84.s1.eu.hivemq.cloud:8883" // Menggunakan SSL/TLS
	MQTT_USER   = "Ergovest"
	MQTT_PASS   = "Ergovest_skripsi123"
	MQTT_TOPIC  = "lowbackpain/sensor/data"
)

// Handler yang dipanggil setiap kali pesan MQTT diterima
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("📥 MQTT Data Received on topic: %s\n", msg.Topic())

	var sensorData models.SensorData
	// Deserialisasi JSON ke struct Golang
	if err := json.Unmarshal(msg.Payload(), &sensorData); err != nil {
		log.Printf("❌ ERROR decoding JSON: %v\n", err)
		return
	}

	log.Printf("✅ Parsed Data - Temp: %.1f, Muscle: %d\n", sensorData.Temp, sensorData.Muscle)

	// --- LOGIKA UTAMA ADA DI SINI ---
	
	// Catatan: Karena AddSensorData dan CheckAndStoreAlert sudah menggunakan Mutex 
	// (di data_service.go) untuk keamanan concurrency, kita bisa memanggilnya di goroutine 
	// agar tidak memblokir thread MQTT.
	
	// 1. Simpan sensorData ke histori In-Memory
	// go AddSensorData(sensorData) // <-- Pindahkan ke goroutine
	go AddSensorData(sensorData)
	
	// 2. Cek Alert Thresholds dan simpan alert yang terdeteksi
	// go CheckAndStoreAlert(sensorData) // <-- Pindahkan ke goroutine
	go CheckAndStoreAlert(sensorData)

	// 3. TODO: Kirim data terbaru ke semua client React yang terhubung via WebSocket
	// go BroadcastWebSocket(sensorData)
}

// StartMQTTClient akan dijalankan sebagai goroutine di main.go
func StartMQTTClient() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(MQTT_BROKER)
	opts.SetClientID("GoBackend_" + fmt.Sprintf("%d", time.Now().Unix()))
	opts.SetUsername(MQTT_USER)
	opts.SetPassword(MQTT_PASS)

	// Atur Handler saat koneksi berhasil
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("✅ Connected to HiveMQ Cloud!")
		token := client.Subscribe(MQTT_TOPIC, 1, messageHandler)
		if token.Wait() && token.Error() != nil {
			log.Fatalf("❌ Error subscribing to topic: %v", token.Error())
		}
		log.Printf("📡 Subscribed to topic: %s\n", MQTT_TOPIC)
	})

	// Atur Handler saat koneksi terputus
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("⚠️ MQTT Connection lost: %v. Reconnecting...", err)
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("❌ Failed to connect to MQTT broker: %v", token.Error())
	}
}