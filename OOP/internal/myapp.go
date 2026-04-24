package internal

import (
	"log"

	"github.com/Isgender-dev/leetcode-solutions/packages"
)

func DoMyCode() {
	computers := []packages.Computer{}
	computers = append(computers, packages.NewComputer(16, 512, 6, "Lenovo LOQ", "800$", "Men"))
	computers = append(computers, packages.NewComputer(32, 1, 16, "Asus ROG", "2800$", "Sen"))
	computers = append(computers, packages.NewComputer(24, 512, 8, "Asus TUF", "1800$", "Ol"))
	//log.Println(computers)

	for _, c := range computers {
		log.Println(c.Ram, c.Storage, c.Cpu, c.Model, c.Price, c.ShowPriceAndOwner())
	}
}
