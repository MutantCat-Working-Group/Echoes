package main

import "com.mutantcat.echose/lifecycle"

func main() {

	gin := lifecycle.InitGin()
	lifecycle.RegisterRouter(gin)
	lifecycle.StartGin(gin, "9091")
}
