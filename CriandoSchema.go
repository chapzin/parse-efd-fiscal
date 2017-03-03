package main

import "createMigration"

func main() {
	createMigration.Drop()
	createMigration.Create()
}
