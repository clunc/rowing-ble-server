package ble

import (
	"encoding/json"
	"fmt"
	"time"
	_ "embed"
)

//go:embed mock/recorded_rowing_data.json
var recordedData []byte

// MockBLEProducer simulates BLE device interactions.
type MockBLEProducer struct{}

// NewMockBLEProducer creates a new MockBLEProducer instance.
func NewMockBLEProducer() *MockBLEProducer {
	return &MockBLEProducer{}
}

// Discover simulates discovering a BLE device.
func (m *MockBLEProducer) Discover() (string, error) {
	fmt.Println("Simulating BLE device discovery...")
	return "MOCK-ADDRESS-01", nil
}

// Start streams mock BLE data loaded from the embedded file.
func (m *MockBLEProducer) Start(packetChan chan<- BLEPacket, stopChan <-chan bool) {
	fmt.Println("Starting Mock BLE Producer...")

	var mockData []struct {
		Characteristic string `json:"characteristic"`
		Data           string `json:"data"`
	}

	if err := json.Unmarshal(recordedData, &mockData); err != nil {
		fmt.Printf("Error decoding mock data: %v\n", err)
		close(packetChan)
		return
	}

	for _, packet := range mockData {
		select {
		case <-stopChan:
			fmt.Println("Stopping Mock BLE Producer...")
			close(packetChan)
			return
		default:
			rawData := HexStringToBytes(packet.Data)
			packetChan <- BLEPacket{
				Characteristic: packet.Characteristic,
				Data:           rawData,
			}
			fmt.Printf("Mock Data Sent: %+v\n", packet)
			time.Sleep(1 * time.Second)
		}
	}
	close(packetChan)
}

// HexStringToBytes converts a hex string into a byte slice.
func HexStringToBytes(hexStr string) []byte {
	data := make([]byte, len(hexStr)/2)
	for i := 0; i < len(data); i++ {
		fmt.Sscanf(hexStr[2*i:2*i+2], "%x", &data[i])
	}
	return data
}
