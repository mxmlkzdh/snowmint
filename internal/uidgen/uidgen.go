package uidgen

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	sequenceBits     = 10
	machineIDBits    = 5
	dataCenterIDBits = 5

	sequenceMask = (1 << sequenceBits) - 1

	maxMachineID    = (1 << machineIDBits) - 1
	maxDataCenterID = (1 << dataCenterIDBits) - 1

	machineIDShift    = sequenceBits
	dataCenterIDShift = sequenceBits + machineIDBits
	timestampShift    = sequenceBits + machineIDBits + dataCenterIDBits
)

type config struct {
	dataCenterID int
	machineID    int
	epoch        int64
}

type UniqueIDGenerator struct {
	config
	sequence      int64
	lastTimestamp int64
	mutex         sync.Mutex
}

func NewUniqueIDGenerator(dataCenterID, machineID int, epoch int64) (*UniqueIDGenerator, error) {
	if dataCenterID < 0 || dataCenterID > maxDataCenterID {
		return nil, fmt.Errorf("'dataCenterID' cannot be less than 0 or greater than %d ", maxDataCenterID)
	}
	if machineID < 0 || machineID > maxMachineID {
		return nil, fmt.Errorf("'machineID' cannot be less than 0 or greater than %d", maxMachineID)
	}
	if epoch < 0 {
		return nil, fmt.Errorf("'epoch' cannot be less than 0")
	}
	log.Printf("NewUniqueIDGenerator: dataCenterID: %d, machineID: %d, epoch: %d", dataCenterID, machineID, epoch)
	return &UniqueIDGenerator{config: config{dataCenterID, machineID, epoch}}, nil
}

func (u *UniqueIDGenerator) GenerateUniqueID() (int64, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	currentTimestamp := time.Now().UnixMilli()
	if currentTimestamp < u.lastTimestamp {
		return 0, fmt.Errorf("cannot generate id because clock moved backwards")
	}
	if currentTimestamp == u.lastTimestamp {
		u.sequence = (u.sequence + 1) & sequenceMask
		if u.sequence == 0 {
			for currentTimestamp <= u.lastTimestamp {
				currentTimestamp = time.Now().UnixMilli()
			}
		}
	} else {
		u.sequence = 0
	}
	u.lastTimestamp = currentTimestamp
	return ((currentTimestamp - u.epoch) << timestampShift) |
		(int64(u.dataCenterID) << dataCenterIDShift) |
		(int64(u.machineID) << machineIDShift) |
		u.sequence, nil
}
