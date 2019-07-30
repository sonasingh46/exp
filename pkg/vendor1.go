package pkg

import "fmt"

// Vendor1Creator implements Create
type Vendor1Creator struct {
	value string
}

// NewVendor1Creator returns a Vendor1Creator instance.
func NewVendor1Creator(machine string) *Vendor1Creator  {
	return &Vendor1Creator{
		// Change the values and run.
		// Supported values : machineX. machineY
		value:machine,
	}
}

// Create is Vendor1Creator implementation.
func (pc *Vendor1Creator) Create(/* Other possible args*/) {
	fmt.Println("Creating via Vendor1")
	f,ok:=createFuncMap[pc.value]
	if !ok{
		fmt.Println("Creation not supported via",pc.value)
		return
	}
	f(pc)
}

// createFunc is a typed function for
// vendor1 machine creators
type createFunc func( *Vendor1Creator)

// createFuncMap holds several
// creator machines supported be vendor1
var createFuncMap= map[string]createFunc{
	"machineX":Vendor1CreatorViaMachineX,
	"machineY":Vendor1CreatorViaMachineY,
}

// Vendor1CreatorViaMachineX creates via machineX
func Vendor1CreatorViaMachineX(pc *Vendor1Creator)  {
	fmt.Println("Creating via ",pc.value)
}

// Vendor1CreatorViaMachineX creates via machineY
func Vendor1CreatorViaMachineY(pc *Vendor1Creator)  {
	fmt.Println("Creating via ",pc.value)
}
