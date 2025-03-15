package director

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/smarthomeix/agents/pkg/service"
)

type Device struct {
	ID            string
	IntegrationID string
	Config        service.DeviceConfig
	RegisteredAt  time.Time
	Telemetry     service.Telemetry
	driver        service.DriverInterface
	ctx           context.Context
	cancel        context.CancelFunc
}

type registry map[string]*Device

type Director struct {
	service  service.ServiceInterface
	registry registry
	mu       sync.RWMutex
	wg       sync.WaitGroup
}

func NewDirector(service service.ServiceInterface) *Director {
	return &Director{
		service:  service,
		registry: make(registry),
	}
}

func (d *Director) GetService() service.ServiceInterface {
	return d.service
}

func (d *Director) Attach(model *Device) (*Device, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.registry[model.ID]; exists {
		return nil, fmt.Errorf("device %s is already registered", model.ID)
	}

	integration, exists := d.service.GetIntegration(model.IntegrationID)
	if !exists {
		return nil, fmt.Errorf("integration %s does not exist", model.IntegrationID)
	}

	driver, err := integration.NewDriver(model.Config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	dev := &Device{
		ID:            model.ID,
		IntegrationID: model.IntegrationID,
		Config:        model.Config,
		RegisteredAt:  time.Now(),
		Telemetry: service.Telemetry{
			Data: make(service.TelemetryData),
		},
		driver: driver,
		ctx:    ctx,
		cancel: cancel,
	}

	d.registry[dev.ID] = dev
	d.wg.Add(1)

	go func(dev *Device) {
		defer d.wg.Done()
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-dev.ctx.Done():
				log.Printf("Telemetry loop stopped for device %s", dev.ID)
				return
			case <-ticker.C:
				telemetryCtx, cancel := context.WithTimeout(dev.ctx, 5*time.Second)
				telemetry, err := dev.driver.GetTelemetry(telemetryCtx)
				cancel()

				updatedAt := time.Now()

				d.mu.Lock()
				if err != nil {
					log.Printf("Failed to get telemetry for device %s: %v", dev.ID, err)
					dev.Telemetry.UpdatedAt = &updatedAt
				} else {
					dev.Telemetry.Data = telemetry
					dev.Telemetry.UpdatedAt = &updatedAt
				}
				d.mu.Unlock()

				log.Printf("Telemetry update for device %s: %+v", dev.ID, dev.Telemetry)
			}
		}
	}(dev)

	return dev, nil
}

func (d *Director) Detach(deviceID string) {
	d.mu.Lock()
	dev, exists := d.registry[deviceID]
	if !exists {
		d.mu.Unlock()
		log.Printf("Attempting to detach device %s which does not exist", deviceID)
		return
	}

	dev.cancel()
	delete(d.registry, deviceID)
	d.mu.Unlock()

	d.wg.Wait()

	log.Printf("Device %s detached successfully", deviceID)
}

func (d *Director) GetDevice(deviceID string) (Device, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	device, exists := d.registry[deviceID]
	if !exists {
		return Device{}, false
	}

	return *device, true
}
