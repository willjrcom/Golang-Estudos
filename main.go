package main

import httpServer "projetoGo/infra/http"

func main() {
	httpServer.InitHttpServer()
}

func teste(x interface{}) interface{} {
	return x
}
