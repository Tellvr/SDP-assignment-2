package main

import "fmt"

type ISexToy interface {
	GetName() string
	GetPower() int
}

type Toy struct {
	name  string
	power int
}

func (t *Toy) GetName() string {
	return t.name
}

func (t *Toy) GetPower() int {
	return t.power
}

type Vibrator struct {
	Toy
}

func newVibrator() ISexToy {
	return &Vibrator{
		Toy: Toy{
			name:  "Magic Wand Vibrator",
			power: 7,
		},
	}
}

type Dildo struct {
	Toy
}

func newDildo() ISexToy {
	return &Dildo{
		Toy: Toy{
			name:  "Realistic Dildo",
			power: 5,
		},
	}
}

func getToy(toyType string) (ISexToy, error) {
	if toyType == "vibrator" {
		return newVibrator(), nil
	}
	if toyType == "dildo" {
		return newDildo(), nil
	}
	return nil, fmt.Errorf("Invalid toy type")
}

func main() {
	vibrator, _ := getToy("vibrator")
	dildo, _ := getToy("dildo")

	printDetails(vibrator)
	printDetails(dildo)
}

func printDetails(t ISexToy) {
	fmt.Println(t.GetName())
	fmt.Println(t.GetPower())
}
