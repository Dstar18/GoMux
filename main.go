package main

import "GoMux/config"

func main() {
	config.InitDB()
	config.InitMigrate()
}
