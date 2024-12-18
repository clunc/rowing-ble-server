package ble

import (
	"fmt"
	"time"
)

// BLEProducer handles real BLE communication.
type BLEProducer struct{}

// NewBLEProducer creates a new BLEProducer instance.
func NewBLEProducer() *BLEProducer {
	return &BLEProducer{}
}

// Discover simulates discovering a BLE device.
func (b *BLEProducer) Discover() (string, error) {
	fmt.Println("Simulating BLE device discovery...")
	return "REAL-BLE-ADDRESS-01", nil
}

// Start streams simulated real BLE data.
func (b *BLEProducer) Start(packetChan chan<- BLEPacket, stopChan <-chan bool) {
	fmt.Println("Starting Real BLE Producer...")

	for i := 0; ; i++ {
		select {
		case <-stopChan:
			fmt.Println("Stopping Real BLE Producer...")
			close(packetChan)
			return
		default:
			packet := BLEPacket{
				Characteristic: "ce060031-43e5-11e4-916c-0800200c9a66",
				Data:           []byte{0x10, 0x20, 0x30, byte(i)},
			}
			packetChan <- packet
			fmt.Printf("Real BLE Data Sent: %x\n", packet.Data)
			time.Sleep(2 * time.Second)
		}
	}
}
