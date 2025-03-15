package director

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/smarthomeix/agents/pkg/messages"
	"github.com/smarthomeix/agents/pkg/service"
	"github.com/smarthomeix/pkg/mqtt/broker"
)

const (
	defaultQoS             = 0
	defaultPublishInterval = 5 * time.Second
	defaultPublishTimeout  = 5 * time.Second
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
	wg            sync.WaitGroup
}

type registry map[string]*Device

type Director struct {
	service  service.ServiceInterface
	broker   *broker.Client
	registry registry
	mu       sync.RWMutex
}

func NewDirector(svc service.ServiceInterface, bkr *broker.Client) *Director {
	return &Director{
		service:  svc,
		broker:   bkr,
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

	device := &Device{
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

	d.registry[device.ID] = device

	device.wg.Add(1)

	go func(dev *Device) {
		defer dev.wg.Done()
		ticker := time.NewTicker(defaultPublishInterval)
		defer ticker.Stop()

		for {
			select {
			case <-dev.ctx.Done():
				log.Printf("Telemetry loop stopped for device %s", dev.ID)
				return
			case <-ticker.C:
				telemetryCtx, cancel := context.WithTimeout(dev.ctx, defaultPublishTimeout)
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

				d.publishMessage(dev, "director")
			}
		}
	}(device)

	return device, nil
}

func (d *Director) Detach(deviceID string) {
	d.mu.Lock()
	dev, exists := d.registry[deviceID]
	if !exists {
		d.mu.Unlock()
		log.Printf("Attempting to detach device %s which does not exist", deviceID)
		return
	}

	// Stop telemetry updates
	dev.cancel()
	delete(d.registry, deviceID)
	d.mu.Unlock()

	// Ensure only this device's goroutine is awaited
	dev.wg.Wait()

	log.Printf("Device %s detached successfully", deviceID)
}

func (d *Director) GetDevice(deviceID string) (*Device, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	device, exists := d.registry[deviceID]

	if !exists {
		return nil, false
	}

	return device, true
}

func (d *Director) publishMessage(dev *Device, source string) {
	d.mu.RLock()

	payload := messages.Telemetry{
		DeviceID:      dev.ID,
		Source:        source,
		IntegrationID: dev.IntegrationID,
		Telemetry:     dev.Telemetry.Data,
		PublishedAt:   time.Now().UTC().Format(time.RFC3339),
	}

	if dev.Telemetry.UpdatedAt != nil {
		updatedAt := dev.Telemetry.UpdatedAt.Format(time.RFC3339)
		payload.UpdatedAt = &updatedAt
	}

	d.mu.RUnlock()

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal telemetry payload for device %s: %v", dev.ID, err)
		return
	}

	// Publish telemetry
	topic := fmt.Sprintf("telemetry/%s", dev.ID)
	token := d.broker.Publish(topic, defaultQoS, false, data)

	if !token.WaitTimeout(defaultPublishTimeout) {
		log.Printf("Timeout publishing to topic %s", topic)
		return
	}

	if err := token.Error(); err != nil {
		log.Printf("Error publishing to topic %s: %v", topic, err)
	} else {
		log.Printf("Successfully published to topic %s", topic)
	}
}
