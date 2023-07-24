package utils

import (
	"fmt"
	"sort"
	"time"
)

func Soma(n1, n2 int) int {
	return n1 + n2
}

func goroutines() {
	canal := make(chan string)

	go addCanal(canal)
	go enviarMsg(canal)

	var entrada string
	fmt.Scanln(&entrada)
}

func addCanal(canal chan string) {
	for i := 0; ; i++ {
		canal <- "teste"
	}
}

func enviarMsg(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func sorts() {
	nomes := []string{
		"William",
		"Nicolas",
	}
	sort.Sort(Textos(nomes))

	fmt.Println(nomes)
}

type Textos []string

func (txt Textos) Len() int {
	return len(txt)
}

func (txt Textos) Less(i, j int) bool {
	return txt[i] < txt[j]
}

func (txt Textos) Swap(i, j int) {
	txt[i], txt[j] = txt[j], txt[i]
}

func floatToString(num float64) string {
	return fmt.Sprintf("%.2f", num)
}
