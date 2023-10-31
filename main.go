package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/SCROLLWH33L/gonomics/genome"
)

func createNewLibrary(library *genome.GenomeMatcher) {
	fmt.Print("Enter minimum search length (3-100): ")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)
	len, err := strconv.Atoi(line)
	if err != nil || len < 3 || len > 100 {
		fmt.Println("Invalid prefix size.")
		return
	}

	*library = genome.NewGenomeMatcher(len)
}

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

	library := genome.NewGenomeMatcher(defaultMinSearchLength)
	reader := bufio.NewReader(os.Stdin)

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
			createNewLibrary(&library)
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
	}
}
