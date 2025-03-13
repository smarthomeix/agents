package director

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/smarthomeix/agents/pkg/service"
)

type device struct {
	device    *service.Device
	driver    service.DriverInterface
	telemetry service.Telemetry
}

type registry map[string]*device

type Director struct {
	service  service.ServiceInterface
	registry registry
	mu       sync.RWMutex // Ensures thread safety
}

func NewDirector(service service.ServiceInterface) *Director {
	return &Director{
		service:  service,
		registry: make(registry), // Initialize map
	}
}

// GetService returns the associated service
func (d *Director) GetService() service.ServiceInterface {
	return d.service
}

// Attach registers a device and stores its driver instance
func (d *Director) Attach(model *service.Device) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Check if device is already registered
	if _, exists := d.registry[model.ID]; exists {
		return fmt.Errorf("device %s is already registered", model.ID)
	}

	// Get integration
	integration, exists := d.service.GetIntegration(model.IntegrationID)

	if !exists {
		return fmt.Errorf("integration %s does not exist", model.IntegrationID)
	}

	// Create device driver
	driver, err := integration.NewDevice(model.Config)

	if err != nil {
		return err
	}

	model.RegisteredAt = time.Now()

	// Store in registry
	d.registry[model.ID] = &device{
		device:    model,
		driver:    driver,
		telemetry: make(service.Telemetry),
	}

	return nil
}

// Detach removes a device from the registry
func (d *Director) Detach(deviceID string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.registry[deviceID]; !exists {
		log.Printf("attempting to detach device %s which does not exist", deviceID)
		return
	}

	delete(d.registry, deviceID)
}

// GetDevice retrieves a device by ID
func (d *Director) GetDevice(deviceID string) (*device, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	device, exists := d.registry[deviceID]

	return device, exists
}
