package main

import "fmt"

func main() {
	var G = [6]string{"Alexander the Great", "Julius Caesar", "Napoleon Bonaparte", "Genghis Kan", "George Washington", "Duke of Caxias"}
	g := len(G)
	var M = []string{}
	var B = []string{}
	var b int

	for i := 0; i < g; i++ {
		M = append(M, ReadMessage(G[i]))
	}

	MajorityDecision := CheckMajority(G, M)
	if (MajorityDecision) {
		fmt.Println("A supermajority of Generals agree to Attack")
	}

	B = CheckBesieged(G, M, MajorityDecision)
	b = len(B)

	fmt.Println("Names of besieged generals:")
	for i := 0; i < b; i++ {
		fmt.Println(B[i])
	}
}

func ReadMessage(general string) string {
	return "Message from " + general
}

func CheckMajority(G [6]string, M []string) bool {
	return true
}

func CheckBesieged(G [6]string, M []string, MajorityDecision bool) []string {
	slicedG := G[:]
	return slicedG
}