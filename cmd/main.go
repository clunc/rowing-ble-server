package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/clunc/rowing-ble-server/internal/ble"
)

// LogEntry represents BLE data with timestamp, characteristic, and data.
type LogEntry struct {
	Timestamp     float64 `json:"timestamp"`
	Characteristic string  `json:"characteristic"`
	Data          string  `json:"data"`
}

func main() {
	fmt.Println("Starting Rowing BLE Server...")

	// Decide whether to use mock or real BLE producer
	useMock := true // Change to false for real BLE producer

	var producer ble.Producer
	if useMock {
		fmt.Println("Using MOCK BLE Producer...")
		producer = ble.NewMockBLEProducer()
	} else {
		fmt.Println("Using REAL BLE Producer...")
		producer = ble.NewBLEProducer()
	}

	// Discover the BLE device
	deviceAddress, err := producer.Discover()
	if err != nil {
		fmt.Printf("Failed to discover BLE device: %v\n", err)
		return
	}
	fmt.Printf("Connected to device at: %s\n", deviceAddress)

	// Channels for BLE data and graceful shutdown
	packetChan := make(chan ble.BLEPacket)
	stopChan := make(chan bool)

	// Start streaming data
	go producer.Start(packetChan, stopChan)

	// Handle graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// Process BLE packets
	go func() {
		for packet := range packetChan {
			timestamp := float64(time.Now().UnixNano()) / 1e9
			logEntry := LogEntry{
				Timestamp:     timestamp,
				Characteristic: packet.Characteristic,
				Data:          fmt.Sprintf("%x", packet.Data),
			}

			// Format as JSON
			jsonOutput, err := json.Marshal(logEntry)
			if err != nil {
				fmt.Printf("Error formatting log entry: %v\n", err)
				continue
			}

			// Print the log entry
			fmt.Println(string(jsonOutput))
		}
	}()

	// Wait for interrupt signal
	<-signalChan
	fmt.Println("Interrupt received, stopping BLE server...")
	stopChan <- true
}
