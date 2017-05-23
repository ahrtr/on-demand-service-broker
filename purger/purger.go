package purger

import (
	fmt "fmt"
	"log"
)

const errorMessageTemplate = "Purger Failed: %s"

type Purger struct {
	deleter   Deleter
	registrar Registrar
	cfClient  CloudFoundryClient
	logger    *log.Logger
}

//go:generate counterfeiter -o fakes/fake_deleter.go . Deleter
type Deleter interface {
	DeleteAllServiceInstances(string) error
}

//go:generate counterfeiter -o fakes/fake_registrar.go . Registrar
type Registrar interface {
	Deregister(string) error
}

//go:generate counterfeiter -o fakes/fake_cloud_foundry_client.go . CloudFoundryClient
type CloudFoundryClient interface {
	DisableServiceAccessForServiceOffering(serviceOfferingID string, logger *log.Logger) error
}

func New(d Deleter, r Registrar, cfClient CloudFoundryClient, logger *log.Logger) *Purger {
	return &Purger{
		deleter:   d,
		registrar: r,
		cfClient:  cfClient,
		logger:    logger,
	}
}

func (p Purger) DeleteInstancesAndDeregister(serviceCatalogID, brokerName string) error {

	p.logger.Println("Disabling service access for all plans")
	err := p.cfClient.DisableServiceAccessForServiceOffering(serviceCatalogID, p.logger)
	if err != nil {
		return fmt.Errorf(errorMessageTemplate, err.Error())
	}

	p.logger.Println("Deleting all service instances")
	err = p.deleter.DeleteAllServiceInstances(serviceCatalogID)
	if err != nil {
		return fmt.Errorf(errorMessageTemplate, err.Error())
	}

	p.logger.Println("Deregistering service brokers")
	err = p.registrar.Deregister(brokerName)
	if err != nil {
		return fmt.Errorf(errorMessageTemplate, err.Error())
	}
	return nil
}
