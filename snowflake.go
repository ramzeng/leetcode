package main

import (
	"errors"
	"fmt"
	"time"
)

func NewSnowflakeServer(datacenterID, workerID int64) *SnowflakeServer {
	parsedTime, _ := time.Parse("2006-01-02 15-04-05", "2023-11-07 23:48:44")

	workerIDBits := int64(5)
	datacenterIDBits := int64(5)
	sequenceBits := int64(12)

	server := &SnowflakeServer{
		StartTimestamp:    parsedTime.Unix(),
		WorkerIDBits:      workerIDBits,
		DatacenterIDBits:  datacenterIDBits,
		WorkerID:          workerID,
		MaxWorkerID:       -1 ^ (-1 << workerIDBits),
		MaxDatacenterID:   -1 ^ (-1 << datacenterID),
		SequenceBits:      sequenceBits,
		WorkerIDShift:     sequenceBits,
		DatacenterIDShift: workerIDBits + sequenceBits,
		TimeStampShift:    datacenterIDBits + workerIDBits + sequenceBits,
		SequenceMask:      -1 ^ (-1 << sequenceBits),
	}

	return server
}

type SnowflakeServer struct {
	StartTimestamp    int64
	WorkerIDBits      int64
	DatacenterIDBits  int64
	MaxWorkerID       int64
	MaxDatacenterID   int64
	SequenceBits      int64
	WorkerIDShift     int64
	DatacenterIDShift int64
	TimeStampShift    int64
	SequenceMask      int64
	WorkerID          int64
	DatacenterID      int64
	Sequence          int64
	LastTimestamp     int64
}

func (s *SnowflakeServer) NextID() (int64, error) {
	timestamp := time.Now().UnixMilli()

	// 发生了时间回拨，无法正确生成 ID，需要报错
	if timestamp < s.LastTimestamp {
		return 0, errors.New("clock moved backwards")
	}

	// 如果是同一毫秒，则按序列号递增
	if timestamp == s.LastTimestamp {
		// 按位与位运算，判断是否溢出
		s.Sequence = (s.Sequence + 1) & s.SequenceMask

		// 序列溢出
		if s.Sequence == 0 {
			return 0, errors.New("sequence overflow")
		}
	} else {
		// 重置序列号
		s.Sequence = 0
	}

	s.LastTimestamp = timestamp

	return (timestamp-s.StartTimestamp)<<s.TimeStampShift |
		(s.DatacenterID << s.DatacenterIDShift) |
		(s.WorkerID << s.WorkerIDShift) |
		s.Sequence, nil
}

func main() {
	server := NewSnowflakeServer(100, 100)

	if id, err := server.NextID(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
}
