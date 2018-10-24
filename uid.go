package uid 

import (
	"bytes"
	"errors"
	"math/big"
	"sync/atomic"
	"time"
)

const (
	// Base : [0-9,a-z]
	Base = 36

	// InstanceIDMax :
	InstanceIDMax = 1000

	// LoopNumMax :
	LoopNumMax = int64(1000)
	// LoopNumStep :
	LoopNumStep = int64(1)

	// LenTime :
	LenTime = int32(8)
	// LenInst :
	LenInst = int32(2)
	// LenLoop :
	LenLoop = int32(2)
)

// UIDGenerator :
type UIDGenerator struct {
	prefix     string
	instanceID int64
	loopNum    *int64
}

// NewUIDGenerator : 1000 different ID per 10ms
func NewUIDGenerator(prefix string, instanceID int32) (gnrtr UIDGenerator, err error) {
	if instanceID > int32(InstanceIDMax) {
		err = errors.New("instanceID is larger than 1000")
		return
	}
	var loopNum = int64(0)
	gnrtr = UIDGenerator{
		prefix,
		int64(instanceID),
		&loopNum,
	}
	return
}

// ID : prefix + 8 time(10ms) + 2 instance + 2 loop number
func (u *UIDGenerator) ID() (ID string) {
	var b bytes.Buffer
	b.WriteString(u.prefix)

	timestamp := big.NewInt(time.Now().UnixNano() / 100000)
	inst := big.NewInt(u.instanceID)
	loopNum := big.NewInt(atomic.AddInt64(u.loopNum, LoopNumStep) % LoopNumMax)

	str := LeftPadString(timestamp.Text(Base), LenTime)
	b.WriteString(str)
	str = LeftPadString(inst.Text(Base), LenInst)
	b.WriteString(str)
	str = LeftPadString(loopNum.Text(Base), LenLoop)
	b.WriteString(str)

	return b.String()
}
