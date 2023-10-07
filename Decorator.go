package main

import "fmt"

type IClothProduct interface {
  GetDetailsAndPrice() (string, int)
}

type Shirt struct{}

func (s *Shirt) GetDetailsAndPrice() (string, int) {
  return "Shirt", 100
}

type Decorations struct {
  ClothProduct IClothProduct
} 

type Embroidery struct {
  Decorations
}

func (e *Embroidery) GetDetailsAndPrice() (string, int) {
  details, price := e.ClothProduct.GetDetailsAndPrice()
  return fmt.Sprintf("%s with Embroidery", details), price + 50
}

type Printing struct {
  Decorations
}

func (p *Printing) GetDetailsAndPrice() (string, int) {
  details, price := p.ClothProduct.GetDetailsAndPrice() 
  return fmt.Sprintf("%s with Printing", details), price + 30
}

func main() {

  shirt := &Shirt{}

  decoratedShirt := &Embroidery{
    Decorations: Decorations{shirt},
  }

  fullyDecorated := &Printing{
    Decorations: Decorations{decoratedShirt},
  }

  details, price := fullyDecorated.GetDetailsAndPrice()

  fmt.Printf("%s - %d\n", details, price)

}