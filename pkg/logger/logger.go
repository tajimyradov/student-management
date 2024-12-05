package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewWithSampler(name, version, env, level string, initial, thereafter int, opts ...zap.Option) (*zap.Logger, error) {
	var config zap.Config
	switch env {
	case "local":
		fallthrough
	case "docker":
		config = zap.NewDevelopmentConfig()
	default:
		config = zap.NewProductionConfig()
	}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	err := config.Level.UnmarshalText([]byte(level))
	if err != nil || len(level) == 0 {
		config.Level.SetLevel(zap.DebugLevel)
	}

	if initial != 0 && thereafter != 0 {
		if config.Sampling == nil {
			config.Sampling = &zap.SamplingConfig{}
		}

		config.Sampling.Initial = initial
		config.Sampling.Thereafter = thereafter
	}

	logger, err := config.Build(opts...)
	if err != nil {
		return nil, err
	}

	logger = logger.With(
		zap.String("name", name),
		zap.String("ver", version),
		zap.String("env", env),
	)

	return logger, nil
}
