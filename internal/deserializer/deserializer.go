package protocol

// PM5Service represents the available services on the PM5 Bluetooth interface.
type PM5Service int

type UUID string

const (
    // PM5 Base UUID Reference: Page 9
    UUIDGAPService       UUID = "0x1800" // GAP Service (Page 10)
    UUIDGATTService      UUID = "0x1801" // GATT Service (Page 10)
    UUIDDeviceInfo       UUID = "0x0010" // Device Information Service (Page 11)
    UUIDControlService   UUID = "0x0020" // PM Control Service (Page 12)
    UUIDRowingService    UUID = "0x0030" // Rowing Service (Page 12)
)

// PM5ServiceUUID maps services to their corresponding UUIDs.
var PM5ServiceUUID = map[PM5Service]UUID{
    ServiceGAP:       UUIDGAPService,
    ServiceGATT:      UUIDGATTService,
    ServiceDeviceInfo: UUIDDeviceInfo,
    ServiceControl:   UUIDControlService,
    ServiceRowing:    UUIDRowingService,
}

// String provides a human-readable representation of the PM5Service.
func (s PM5Service) String() string {
    switch s {
    case ServiceGAP:
        return "GAP Service" // Table 3, Page 10
    case ServiceGATT:
        return "GATT Service" // Table 3, Page 10
    case ServiceDeviceInfo:
        return "Device Information Service" // Table 3, Page 11
    case ServiceControl:
        return "Control Service" // Table 3, Page 12
    case ServiceRowing:
        return "Rowing Service" // Table 3, Page 12
    default:
        return "Unknown Service"
    }
}

// GAPService represents the GAP service characteristics.
// Reference: Table 3, Page 10
type GAPService struct {
    DeviceName       string             // 0x2A00: Device name (Page 10)
    Appearance       uint16             // 0x2A01: Appearance (Page 10)
    PeripheralPrivacy bool              // 0x2A02: Privacy (Page 10)
    ConnectionParams ConnectionParameters // 0x2A04: Connection Params (Page 10)
}

// ConnectionParameters defines the preferred connection parameters.
// Reference: Table 3, Page 10
type ConnectionParameters struct {
    MinInterval       uint16 // Preferred min connection interval
    MaxInterval       uint16 // Preferred max connection interval
    SlaveLatency      uint16 // Slave latency
    SupervisionTimeout uint16 // Supervision timeout
}

// DeviceInfoService contains details about the PM5 device.
// Reference: Table 3, Page 11
type DeviceInfoService struct {
    ModelNumber       string // 0x0011: Model Number (Page 11)
    SerialNumber      string // 0x0012: Serial Number (Page 11)
    HardwareRevision  string // 0x0013: Hardware Revision (Page 11)
    FirmwareRevision  string // 0x0014: Firmware Revision (Page 11)
    ManufacturerName  string // 0x0015: Manufacturer Name (Page 11)
}

// PMControlService handles sending and receiving CSAFE frames.
// Reference: Table 3, Page 12
type PMControlService struct {
    CommandFrame  []byte // 0x0021: CSAFE command frame (Page 12)
    ResponseFrame []byte // 0x0022: CSAFE response frame (Page 12)
}

// CSAFECommandType enumerates common CSAFE command types.
// Reference: CSAFE Protocol Specification
type CSAFECommandType uint8

const (
    CSAFEGetWorkoutType CSAFECommandType = 0x89
    CSAFEGetRowingState CSAFECommandType = 0x8D
    CSAFESetWorkout     CSAFECommandType = 0x23
)

func IsValidCSAFECommand(cmd CSAFECommandType) bool {
    switch cmd {
    case CSAFEGetWorkoutType, CSAFEGetRowingState, CSAFESetWorkout:
        return true
    default:
        return false
    }
}

// RowingGeneralStatus represents the 0x0031 General Status characteristic.
// Reference: Table 3, Page 13
type RowingGeneralStatus struct {
    ElapsedTime     uint32 // 3 bytes (0.01 sec lsb)
    Distance        uint32 // 3 bytes (0.1 m lsb)
    WorkoutType     uint8  // Enum for workout type
    IntervalType    uint8  // Enum for interval type
    WorkoutState    uint8  // Enum for workout state
    RowingState     uint8  // Enum for rowing state
    StrokeState     uint8  // Enum for stroke state
    DragFactor      uint8  // Drag factor
}

// RowingStrokeData represents the 0x0035 Stroke Data characteristic.
// Reference: Table 3, Page 17
type RowingStrokeData struct {
    DriveLength       uint8  // Drive length (0.01 m)
    DriveTime         uint8  // Drive time (0.01 sec)
    StrokeRecovery    uint16 // Recovery time (0.01 sec)
    PeakDriveForce    uint16 // Peak force (0.1 lbs)
    AverageDriveForce uint16 // Avg force (0.1 lbs)
    StrokeDistance    uint16 // Distance/stroke (0.01 m)
    WorkPerStroke     uint16 // Work/stroke (0.1 Joules)
    StrokeCount       uint16 // Total stroke count
}

// WorkoutType represents the type of workout.
// Reference: Appendix A, Page 36
type WorkoutType uint8

const (
    WorkoutJustRow WorkoutType = iota
    WorkoutFixedDistance
    WorkoutFixedTime
    WorkoutInterval
)

// RowingState represents the rowing activity state.
// Reference: Appendix A, Page 37
type RowingState uint8

const (
    RowingInactive RowingState = iota
    RowingActive
)

// StrokeState represents the state of the rowing stroke.
// Reference: Appendix A, Page 37
type StrokeState uint8

const (
    StrokeWaiting StrokeState = iota
    StrokeDriving
    StrokeRecovery
)

// MultiplexedRowingData represents a combined rowing data packet.
// Reference: Table 4, Page 24
type MultiplexedRowingData struct {
    Identifier uint8  // Identifies which data stream this packet belongs to
    Data       []byte // Up to 19 bytes of packed data
}
