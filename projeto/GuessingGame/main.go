package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um numero aleatorio sera sorteado. Tente acertar. O numero e um inteiro entre 0 e 100")

	x1 := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual e o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("O seu chute tem que ser um numero inteiro")
			return
		}

		switch {
		case chuteInt < x1:
			fmt.Println("Voce errou. O numero sorteado e maior que ", chuteInt)
		case chuteInt > x1:
			fmt.Println("Voce errou. O numero sorteado e menor que ", chuteInt)
		case chuteInt == x1:
			return
		}

		chutes[i] = chuteInt
	}
}
