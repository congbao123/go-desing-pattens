package main

import (
    "fmt"
)

// type iTransport interface {
// 	setName(name string)
// 	setPower(power int)
// 	getName() string
// 	getPower() int
// }

// type transport struct {
// 	name  string
// 	power int
// }

// func (g *transport) setName(name string) {
// 	g.name = name
// }

// func (g *transport) getName() string {
// 	return g.name
// }

// func (g *transport) setPower(power int) {
// 	g.power = power
// }

// func (g *transport) getPower() int {
// 	return g.power
// }

// type truck struct {
// 	transport
// }

// func newTruck() iTransport {
// 	return &truck{
// 		transport: transport{
// 			name:  "Truck",
// 			power: 4,
// 		},
// 	}
// }

// type ships struct {
// 	transport
// }

// func newShips() iTransport {
// 	return &ships{
// 		transport: transport{
// 			name:  "Ships",
// 			power: 5,
// 		},
// 	}
// }

// func getTransport(transportType string) (iTransport, error) {
// 	if transportType == "truck" {
// 		return newTruck(), nil
// 	} else if transportType == "ships" {
// 		return newShips(), nil
// 	}
// 	return nil, fmt.Errorf("Wrong transport type")
// }

// func printDetails(g iTransport) {
// 	fmt.Printf("Type: %s\n", g.getName())
// 	fmt.Printf("Power: %d\n", g.getPower())
// 	fmt.Println()
// }

// func main() {
// 	truck, _ := getTransport("truck")
// 	ships, _ := getTransport("ships")
// 	printDetails(truck)
// 	printDetails(ships)
// }

// Tạo đối tượng đặt ghế hiện đại và vip 
type Chair interface {
	SitOn()        
	GetDescription() string 
}

// Sofa interface
type Sofa interface {
	LayOn()         
	GetDescription() string
}

// CoffeeTable interface
type CoffeeTable interface {
	PlaceItems() 
	GetDescription() string 
}
// ModernChair struct implements Chair interface
type ModernChair struct{}

func (m *ModernChair) SitOn() {
	fmt.Println("Sitting on a modern chair.")
}

func (m *ModernChair) GetDescription() string {
	return "This is a modern chair."
}

// ModernSofa struct implements Sofa interface
type ModernSofa struct{}

func (m *ModernSofa) LayOn() {
	fmt.Println("Lying on a modern sofa.")
}

func (m *ModernSofa) GetDescription() string {
	return "This is a modern sofa."
}

// ModernCoffeeTable struct implements CoffeeTable interface
type ModernCoffeeTable struct{}

func (m *ModernCoffeeTable) PlaceItems() {
	fmt.Println("Placing items on a modern coffee table.")
}

func (m *ModernCoffeeTable) GetDescription() string {
	return "This is a modern coffee table."
}

// VictorianChair class
type VictorianChair struct{}

func (v *VictorianChair) SitOn() {
	fmt.Println("Sitting on a Victorian chair.")
}

func (v *VictorianChair) GetDescription() string {
	return "This is a Victorian chair."
}

// VictorianSofa class
type VictorianSofa struct{}

func (v *VictorianSofa) LayOn() {
	fmt.Println("Lying on a Victorian sofa.")
}

func (v *VictorianSofa) GetDescription() string {
	return "This is a Victorian sofa."
}

// VictorianCoffeeTable class
type VictorianCoffeeTable struct{}

func (v *VictorianCoffeeTable) PlaceItems() {
	fmt.Println("Placing items on a Victorian coffee table.")
}

func (v *VictorianCoffeeTable) GetDescription() string {
	return "This is a Victorian coffee table."
}

// Abstract Factory interface
type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
	CreateCoffeeTable() CoffeeTable
}
// ModernFurnitureFactory class
type ModernFurnitureFactory struct{}

func (m *ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (m *ModernFurnitureFactory) CreateSofa() Sofa {
	return &ModernSofa{}
}

func (m *ModernFurnitureFactory) CreateCoffeeTable() CoffeeTable {
	return &ModernCoffeeTable{}
}

// VictorianFurnitureFactory class
type VictorianFurnitureFactory struct{}

func (v *VictorianFurnitureFactory) CreateChair() Chair {
	return &VictorianChair{}
}

func (v *VictorianFurnitureFactory) CreateSofa() Sofa {
	return &VictorianSofa{}
}

func (v *VictorianFurnitureFactory) CreateCoffeeTable() CoffeeTable {
	return &VictorianCoffeeTable{}
}

func GetFactory(style string) FurnitureFactory {
	switch style {
	case "modern":
		return &ModernFurnitureFactory{}
	case "victorian":
		return &VictorianFurnitureFactory{}
	default:
		panic("Unknown style!")
	}
}

func main() {
	// Example: Initialize the factory based on configuration
	style := "modern" // Can be "modern" or "victorian"
	factory := GetFactory(style)

	// Use the factory to create products
	chair := factory.CreateChair()
	sofa := factory.CreateSofa()
	table := factory.CreateCoffeeTable()

	// Interact with the products
	chair.SitOn()
	fmt.Println(chair.GetDescription())

	sofa.LayOn()
	fmt.Println(sofa.GetDescription())

	table.PlaceItems()
	fmt.Println(table.GetDescription())
}

