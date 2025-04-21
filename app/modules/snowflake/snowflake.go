package snowflake

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	epoch          = 1738713600000 // February 5, 2025 00:00:00 UTC
	machineIDBits  = 10            // 10 bits for machine ID
	sequenceBits   = 12            // 12 bits for sequence number
	maxMachineID   = (1 << machineIDBits) - 1
	maxSequence    = (1 << sequenceBits) - 1
	timestampShift = machineIDBits + sequenceBits
	machineIDShift = sequenceBits
)

type SnowflakeID int64

func (id SnowflakeID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strconv.FormatInt(int64(id), 10) + `"`), nil
}

func (id *SnowflakeID) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return fmt.Errorf("invalid snowflake string: %w", err)
	}

	parsed, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid snowflake number: %w", err)
	}

	*id = SnowflakeID(parsed)
	return nil
}

// Snowflake represents a unique ID generator.
type Snowflake struct {
	mutex     sync.Mutex
	timestamp int64
	machineID int64
	sequence  int64
	timeFunc  func() int64 // Custom time function for testing
}

// NewSnowflake creates a new Snowflake instance with the given machine ID.
func NewSnowflake(machineID int64, timeFunc func() int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, errors.New("invalid machine ID: must be between 0 and " + strconv.Itoa(maxMachineID))
	}
	if timeFunc == nil {
		timeFunc = time.Now().UnixMilli // Default time function
	}
	return &Snowflake{
		machineID: machineID,
		timeFunc:  timeFunc,
	}, nil
}

// Generate creates a new unique ID.
func (s *Snowflake) Generate() SnowflakeID {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := s.timeFunc()
	if now == s.timestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// Sequence exhausted, wait until next millisecond
			for now <= s.timestamp {
				time.Sleep(1 * time.Millisecond) // Sleep for 1 millisecond
				now = s.timeFunc()
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now
	id := (now-epoch)<<timestampShift | (s.machineID << machineIDShift) | s.sequence
	return SnowflakeID(id)
}
