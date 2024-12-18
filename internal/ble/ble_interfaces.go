package ble

// BLEPacket combines characteristic UUID and raw data.
type BLEPacket struct {
	Characteristic string
	Data           []byte
}

// Discoverer is responsible for discovering BLE devices.
type Discoverer interface {
	Discover() (string, error)
}

// Streamer is responsible for streaming BLE data.
type Streamer interface {
	Start(chan<- BLEPacket, <-chan bool)
}

// Producer combines Discoverer and Streamer functionalities.
type Producer interface {
	Discoverer
	Streamer
}
