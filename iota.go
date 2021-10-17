package main

import "fmt"

const (
	isPrincipal = 1 << iota
	isCollege
	canGoAccounts 
	canGoCSE
	canGoISE
	canGoMBA
	canGoMCA
	canGoECE
)

func main()  {
	var roles byte = isPrincipal | canGoAccounts | canGoMBA
	fmt.Printf("%b\n", roles)
	// fmt.Printf("Is Principal? %v \n", isPrincipal & roles == isPrincipal)
	fmt.Printf("Is College? %v \n", isCollege & roles == isCollege)

}