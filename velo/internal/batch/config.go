package batch

import (
	"time"

	"github.com/velo-api/velo/pkg/config"
)

type BatchConfig struct {
	Enabled       bool
	Window        time.Duration
	MaxBatchSize  int
	FlushInterval time.Duration
	Classifiers   []string
}

func ParseConfig(cfg config.BatchConfig) *BatchConfig {
	window, err := time.ParseDuration(cfg.Window)
	if err != nil {
		window = 5 * time.Millisecond
	}

	flushInterval, err := time.ParseDuration(cfg.FlushInterval)
	if err != nil {
		flushInterval = 10 * time.Millisecond
	}

	maxBatchSize := cfg.MaxBatchSize
	if maxBatchSize <= 0 {
		maxBatchSize = 100
	}

	classifiers := cfg.Classifiers
	if len(classifiers) == 0 {
		classifiers = []string{"endpoint"}
	}

	return &BatchConfig{
		Enabled:       cfg.Enabled,
		Window:        window,
		MaxBatchSize:  maxBatchSize,
		FlushInterval: flushInterval,
		Classifiers:   classifiers,
	}
}
