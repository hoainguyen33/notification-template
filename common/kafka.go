package common

import "time"

var (
	// KafkaURLs              = []string{}
	// GroupID                = ""
	// Topic                  = ""
	MinBytes      = 1
	MaxBytes      = 1 * 1024 * 1024
	QueueCapacity = 10000
	// HeartbeatInterval      = 3 * time.Second
	// CommitInterval         = 0 * time.Second
	// PartitionWatchInterval = 5 * time.Second
	MaxAttempts           = 3
	DialTimeout           = 2 * 60 * time.Second
	WriterMaxAttempts     = 10
	WriterReadTimeout     = 10 * time.Second
	WriterWriteTimeout    = 10 * time.Second
	BatchTimeout          = 10 * time.Millisecond
	BatchSize             = 10
	PartitionNotification = 0
	CommitInterval        = 0 * time.Second
)
