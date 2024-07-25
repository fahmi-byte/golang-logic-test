package test

import (
	"fmt"
	"testing"
)

type FruitType string

const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

type Fruits struct {
	Id    int
	Name  string
	Type  FruitType
	Stock int
}

func TestFruitsCase(t *testing.T) {

	fruits := []Fruits{
		{
			Id:    1,
			Name:  "Apel",
			Type:  "IMPORT",
			Stock: 10,
		},
		{
			Id:    2,
			Name:  "Kurma",
			Type:  "IMPORT",
			Stock: 20,
		}, {
			Id:    3,
			Name:  "Apel",
			Type:  "IMPORT",
			Stock: 50,
		},
		{
			Id:    4,
			Name:  "Manggis",
			Type:  "LOCAL",
			Stock: 100,
		},
		{
			Id:    5,
			Name:  "Jeruk Bali",
			Type:  "LOCAL",
			Stock: 10,
		},
		{
			Id:    5,
			Name:  "Kurma",
			Type:  "IMPORT",
			Stock: 20,
		},
		{
			Id:    5,
			Name:  "Salak",
			Type:  "LOCAL",
			Stock: 150,
		},
	}

	uniqueFruitNames := UniqueFruitNames(fruits)

	fmt.Println("Buah yang andi miliki:")
	for _, name := range uniqueFruitNames {
		fmt.Println("- " + name)
	}

	groupedFruits, totalStock := GroupFruitsByType(fruits)

	fmt.Printf("\nTotal wadah yang dibutuhkan: %d\n", len(groupedFruits))

	for fruitType, names := range groupedFruits {
		fmt.Printf("\nWadah %s\n", fruitType)
		fmt.Println("Buahnya adalah:")
		for _, name := range names {
			fmt.Printf("- %s\n", name)
		}
		fmt.Printf("Total Stock dari wadah %s: %d\n", fruitType, totalStock[fruitType])
	}

}

func UniqueFruitNames(fruits []Fruits) []string {
	uniqueNames := make(map[string]bool)
	var fruitNames []string

	for _, fruit := range fruits {
		if !uniqueNames[fruit.Name] {
			uniqueNames[fruit.Name] = true
			fruitNames = append(fruitNames, fruit.Name)
		}
	}

	return fruitNames
}

// Function to group fruits by type
func GroupFruitsByType(fruits []Fruits) (map[FruitType][]string, map[FruitType]int) {
	groupedFruits := make(map[FruitType][]string)
	totalStock := make(map[FruitType]int)

	for _, fruit := range fruits {
		groupedFruits[fruit.Type] = append(groupedFruits[fruit.Type], fruit.Name)
		totalStock[fruit.Type] += fruit.Stock
	}

	return groupedFruits, totalStock
}
