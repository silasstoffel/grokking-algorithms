package main

import "fmt"

type State struct {
	Id        int
	Name      string
	ShortName string
}

func main() {

	states := make(map[string]State)
	ufs := [4]string{
		"ES",
		"MG",
		"RJ",
		"SP",
	}

	states["ES"] = State{
		Id:        1,
		Name:      "Espirito Santo",
		ShortName: "ES",
	}

	states["SP"] = State{
		Id:        2,
		Name:      "Sao Paulo",
		ShortName: "SP",
	}

	states["RJ"] = State{
		Id:        3,
		Name:      "Rio de Janeiro",
		ShortName: "RJ",
	}

	search := "SP"
	if _, exists := states[search]; exists {
		fmt.Println(search, "already exists")
	}

	search = "MG"
	if _, exists := states[search]; !exists {
		fmt.Println(search, "does not exists, adding...")
		states["MG"] = State{
			Id:        4,
			Name:      "Minas Gerais",
			ShortName: "MG",
		}
	}

	fmt.Println("\nPrinting states")

	for _, uf := range ufs {
		fmt.Println("UF:", uf, states[uf])
	}

	fmt.Println("\nRemoving MG...")
	delete(states, "MG")

	fmt.Println("\nPrinting states again...")
	for _, uf := range ufs {
		if _, exists := states[uf]; exists {
			fmt.Println("UF:", uf, states[uf])
		}
	}
	fmt.Println("")
}
