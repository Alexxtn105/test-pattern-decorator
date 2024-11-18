package main

import "fmt"

// Coffee Интерфейс, описывающий кофе
type Coffee interface {
	Cost() float64
	Description() string
}

// SimpleCoffee Конкретный компонент кофе
type SimpleCoffee struct{}

// Cost Реализация метода возврата стоимости для обычного кофе
func (s *SimpleCoffee) Cost() float64 {
	return 2.0
}

// Description Реализация метода, возвращающего описание обычного кофе
func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Milk Декоратор № 1 - Молоко
type Milk struct {
	coffee Coffee
}

// Cost Возвращаем цену задекорированного кофе с молоком
func (m *Milk) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

// Description Возвращаем описание задекорированного кофе с молоком
func (m *Milk) Description() string {
	return m.coffee.Description() + ", Milk"
}

// Caramel Декоратор № 2 - Карамель
type Caramel struct {
	coffee Coffee
}

// Cost Возвращаем цену задекорированного кофе с карамелью
func (c *Caramel) Cost() float64 {
	return c.coffee.Cost() + 1.0
}

// Description Возвращаем описание задекорированного кофе с карамелью
func (c *Caramel) Description() string {
	return c.coffee.Description() + ", Caramel"
}

func main() {
	// создаем обычный кофе
	coffee := &SimpleCoffee{}

	// Декорируем кофе молоком
	coffeeWithMilk := &Milk{coffee: coffee}
	fmt.Println("Coffee with Milk:", coffeeWithMilk.Description(), " - Cost:", coffeeWithMilk.Cost())

	coffeeWithCaramel := &Caramel{coffee: coffee}
	fmt.Println("Coffee with Caramel:", coffeeWithCaramel.Description(), " - Cost:", coffeeWithCaramel.Cost())

	coffeeWithMilkAndCaramel := &Milk{coffee: &Caramel{coffee: coffee}}
	fmt.Println("Coffee with Milk And Caramel:", coffeeWithMilkAndCaramel.Description(), " - Cost:", coffeeWithMilkAndCaramel.Cost())

}
