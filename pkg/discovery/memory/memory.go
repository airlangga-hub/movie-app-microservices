package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/airlangga-hub/movie-app-microservices/pkg/discovery"
)

type Registry struct {
	sync.RWMutex
	serviceAddrs map[string]map[string]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

func New() *Registry {
	return &Registry{serviceAddrs: map[string]map[string]*serviceInstance{}}
}

func (r *Registry) Register(ctx context.Context, instanceID, serviceName, hostPort string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceName]; !ok {
		r.serviceAddrs[serviceName] = make(map[string]*serviceInstance)
	}

	r.serviceAddrs[serviceName][instanceID] = &serviceInstance{hostPort: hostPort, lastActive: time.Now()}

	return nil
}

func (r *Registry) Deregister(ctx context.Context, instanceID, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return nil
	}

	delete(r.serviceAddrs[serviceName], instanceID)

	return nil
}

func (r *Registry) ReportHealthyState(instanceID, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return errors.New("service is not registered yet")
	}

	if _, ok := r.serviceAddrs[serviceName][instanceID]; !ok {
		return errors.New("instance " + instanceID + " of service " + serviceName + " is not registere yet")
	}

	r.serviceAddrs[serviceName][instanceID].lastActive = time.Now()

	return nil
}

func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()

	countInstances := len(r.serviceAddrs[serviceName])

	if countInstances == 0 {
		return nil, discovery.ErrNotFound
	}

	res := make([]string, 0, countInstances)

	for _, instance := range r.serviceAddrs[serviceName] {
		res = append(res, instance.hostPort)
	}

	return res, nil
}
