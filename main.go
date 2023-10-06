package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SCROLLWH33L/gonomics/genome"
	"github.com/SCROLLWH33L/gonomics/trie"
)

func showMenu() {
	fmt.Println("        Commands:")
	fmt.Println("         c - create new genome library      s - find matching SNiPs")
	fmt.Println("         a - add one genome manually        r - find related genomes (manual)")
	fmt.Println("         l - load one data file             f - find related genomes (file)")
	fmt.Println("         d - load all provided data files   ? - show this menu")
	fmt.Println("         e - find matches exactly           q - quit")
}

func loadFile(filename string, genomes *[]genome.Genome) error {
	f, err := os.Open(filename)

	if err != nil {
		return err
	}

	genomeSource := bufio.NewReader(f)
	return genome.Load(genomeSource, genomes)
}

func main() {
	const defaultMinSearchLength = 10

	fmt.Println("Welcome to the Gee-nomics test harness!")
	fmt.Println("The genome library is initially empty, with a default minSearchLength of", defaultMinSearchLength)
	showMenu()

	// genomes := &[]genome.Genome{}
	// err := loadFile("../Gee-nomics/data/Halorubrum_chaoviator.txt", genomes)

	// Desulfurococcus_mucosus.txt
	// Ferroglobus_placidus.txt
	// Ferroplasma_acidarmanus.txt
	// Halobacterium_jilantaiense.txt
	// Halorientalis_persicus.txt
	// Halorientalis_regularis.txt
	// Halorubrum_californiense.txt
	// Halorubrum_chaoviator.txt

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(len(*genomes))

	t := trie.Trie[uint]{}

	t.Insert("a", 14)
	t.Insert("hi", 9)
	t.Insert("hi", 17)
	t.Insert("to", 22, 23)
	t.Insert("hit", 1, 2)
	t.Insert("hip", 10, 20)
	t.Insert("hat", 7, 8, 9)
	t.Insert("tap", 19, 6, 32)

	fmt.Println(t)
	fmt.Println(t.Find("hit", false))

	t.Reset()
	fmt.Println(t)

	// Private fields start with a lowercase
	// Public fields start with an uppercase

	/* reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		command, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		switch rn, _ := utf8.DecodeRuneInString(command); unicode.ToLower(rn) {
		case 'q':
			return
		case '?':
			showMenu()
		case 'c':
			fmt.Println("TODO")
		case 'a':
			fmt.Println("TODO")
		case 'l':
			fmt.Println("TODO")
		case 'd':
			fmt.Println("TODO")
		case 'e':
			fmt.Println("TODO")
		case 's':
			fmt.Println("TODO")
		case 'r':
			fmt.Println("TODO")
		case 'f':
			fmt.Println("TODO")
		default:
			fmt.Println("Invalid command", command)
		}
	} */
}
