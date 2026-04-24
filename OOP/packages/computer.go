package packages

// Polymorthism
type Machine interface {
	CalculateWorth() int
}

type Computer struct {
	Ram     int
	Storage int
	Cpu     int
	Model   string
	Price   string
	owner   string
}

// Inthern
type Server struct {
	Computer
}

func (c Computer) SetOwner(newOwner string) {
	c.owner = newOwner
}

// Encapsulation
func (c Computer) ShowPriceAndOwner() string {
	return c.owner + " " + c.Price
}

// Abstraction
func (c Computer) CalculateWorth() int {
	return c.Ram * c.Cpu
}

func NewComputer(ram int, storage int, cpu int, model string, price string, owner string) Computer {
	return Computer{
		Ram:     ram,
		Storage: storage,
		Cpu:     cpu,
		Model:   model,
		Price:   price,
		owner:   owner,
	}
}
