package main

import (
	"fmt"
	"github.com/sonasingh46/exp/pkg"
)

// Background :

// There are different vendors in market who creates some imaginary items
// Every vendors again have different specialized machines to create thos items.

// But there is a single supreme creator in the market and every customer go to the supreme creator and
// the supreme creator finally creates the imaginary item as specified by customer create configuration
// which decides the final vendor and creating machine of the vendor.

// Vendor1 supports creation via  'machineX' and 'machineY'

// Vendor2 supports creation via  'machineX' and 'machineY'

/*

pkg/vendor1.go contains vendor1 creation implementation via a functional approach of map

pkg/vendor2.go contains vendor2 creation implementation via interface approach.

I am trying to compare and contrast which one is a suitable and preferable approach with respect to
following factors:

1. Extensibility
2. Maintainability
3. Simplicity
4. Idiomatic Go Patterns

 */

func main(){
	// Test Code for vendor 1 creation
	// Get create config of vendor1
	vendor1CreateConfig:=pkg.NewVendor1Creator("machineX")

	// Create the supreme creator by passing the create config
	// This should run to finally create the item.
	pcv1:=NewCreateConfig(vendor1CreateConfig)

	pcv1.Run()

	// Test Code for vendor 2 creation
	// Get create config of vendor1
	vendor2CreateConfig:=pkg.NewVendor2Creator("machineY")

	// Create the supreme creator by passing the create config
	// This should run to finally create the item.
	pcv2:=NewCreateConfig(vendor2CreateConfig)

	pcv2.Run()



}

// Following code from supreme creator side
// that has an interface Create and this is implemented
// by different vendors

// CreateConfig knows how to create something.
type CreateConfig struct {
	creator Creator
}

// Creator interfaces abstract Create()
type Creator interface {
	Create()
}

// NewCreateConfig returns an instance of CreateConfig.
func NewCreateConfig(p Creator) *CreateConfig {
	return &CreateConfig{
		creator:p,
	}
}
// Run function is executed to create finally based on CreateConfig
func(pc *CreateConfig)Run(){
	fmt.Println("Creating...")
	pc.creator.Create()
}













