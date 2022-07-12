package main

import "golang-mongodb/config/delivery"

func main() {
	delivery.NewServer().Run()
}