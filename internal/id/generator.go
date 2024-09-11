package id

import (
	"fmt"
	"sync"
	"time"
)

const (
	sequenceBits      = 10
	nodeIDBits        = 5
	dataCenterIDBits  = 5
	sequenceMask      = (1 << sequenceBits) - 1
	maxNodeID         = (1 << nodeIDBits) - 1
	maxDataCenterID   = (1 << dataCenterIDBits) - 1
	nodeIDShift       = sequenceBits
	dataCenterIDShift = sequenceBits + nodeIDBits
	timestampShift    = sequenceBits + nodeIDBits + dataCenterIDBits
)

type config struct {
	dataCenterID int
	nodeID       int
	epoch        int64
}

type UniqueIDGenerator struct {
	config
	sequence      int64
	lastTimestamp int64
	mutex         sync.Mutex
}

func NewUniqueIDGenerator(dataCenterID, nodeID int, epoch int64) (*UniqueIDGenerator, error) {
	if dataCenterID < 0 || dataCenterID > maxDataCenterID {
		return nil, fmt.Errorf("'dataCenterID' cannot be less than 0 or greater than %d ", maxDataCenterID)
	}
	if nodeID < 0 || nodeID > maxNodeID {
		return nil, fmt.Errorf("'nodeID' cannot be less than 0 or greater than %d", maxNodeID)
	}
	if epoch < 0 {
		return nil, fmt.Errorf("'epoch' cannot be less than 0")
	}
	return &UniqueIDGenerator{config: config{dataCenterID, nodeID, epoch}}, nil
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
		(int64(u.nodeID) << nodeIDShift) |
		u.sequence, nil
}
