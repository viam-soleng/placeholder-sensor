package models

import (
	"context"
	"errors"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/utils/rpc"
)

var (
	Sensor           = resource.NewModel("bill", "placeholder", "sensor")
	errUnimplemented = errors.New("unimplemented")
)

func init() {
	resource.RegisterComponent(sensor.API, Sensor,
		resource.Registration[sensor.Sensor, *Config]{
			Constructor: newPlaceholderSensor,
		},
	)
}

type Config struct {
	// Generic structure to hold any kind of data
	Readings map[string]interface{} `json:"readings"`
}

// Validate ensures all parts of the config are valid.
func (cfg *Config) Validate(path string) ([]string, error) {
	// Basic validation to ensure the readings field exists
	if cfg.Readings == nil {
		return nil, errors.New("readings field must be provided in configuration")
	}
	return nil, nil
}

type placeholderSensor struct {
	name   resource.Name
	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()
}

func newPlaceholderSensor(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (sensor.Sensor, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &placeholderSensor{
		name:       rawConf.ResourceName(),
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *placeholderSensor) Name() resource.Name {
	return s.name
}

// Reconfigure updates the sensor with new configuration data.
func (s *placeholderSensor) Reconfigure(ctx context.Context, deps resource.Dependencies, newConf resource.Config) error {
	s.logger.Info("Reconfiguring placeholder sensor")

	// Parse the new configuration
	conf, err := resource.NativeConfig[*Config](newConf)
	if err != nil {
		return err
	}

	// Update the stored configuration
	s.cfg = conf
	s.logger.Info("Placeholder sensor reconfigured successfully")

	return nil
}

func (s *placeholderSensor) NewClientFromConn(ctx context.Context, conn rpc.ClientConn, remoteName string, name resource.Name, logger logging.Logger) (sensor.Sensor, error) {
	return sensor.NewClientFromConn(ctx, conn, remoteName, name, logger)
}

func (s *placeholderSensor) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
	// Simply return the readings from the config
	return s.cfg.Readings, nil
}

func (s *placeholderSensor) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	// Return the current readings
	return map[string]interface{}{
		"readings": s.cfg.Readings,
	}, nil
}

func (s *placeholderSensor) Close(context.Context) error {
	s.cancelFunc()
	return nil
}
