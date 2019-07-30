package pkg

import (
	"fmt"
)


// Vendor2Creator implements Create
type Vendor2Creator struct {
	value string
}

// NewVendor1Creator returns a Vendor1Creator instance.
func NewVendor2Creator(machine string) *Vendor2Creator  {
	return &Vendor2Creator{
		// Change the values and run.
		// Supported values : machineX. machineY
		value:machine,
	}
}

// Create is Vendor2Creator implementation.
func (pc *Vendor2Creator) Create(/* Other possible args*/) {
	fmt.Println("Creating via Vendor2")
	v1m:=NewVendor1Machine(pc)
	v1m.Create(pc)
}

// Machine abstracts Create at vendor layer
type Machine interface {
	Create(*Vendor2Creator)
}

// NewVendor1Machine returns a machine
func NewVendor1Machine(pc *Vendor2Creator)Machine  {
	if pc.value=="machineX"{
		return &CreationViaX{}
	}

	if pc.value=="machineY"{
		return &CreationViaX{}
	}

	return &CreationNoop{}
}

type CreationViaX struct {}

type CreationViaY struct {}

type CreationNoop struct {}

func (pp *CreationViaX)Create(pc *Vendor2Creator)  {
	fmt.Println("Provisioning via ",pc.value)
}

func (bp *CreationViaY)Create(pc *Vendor2Creator)  {
	fmt.Println("Provisioning via ",pc.value)
}

func (noop *CreationNoop)Create(op *Vendor2Creator)  {
	fmt.Printf("Provisioning via %s not supported\n",op.value)
}














