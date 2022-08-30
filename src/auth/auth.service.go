package auth

import "fmt"

func create(data interface{}) string {
	return "This action adds a new auth"
}

func findAll() string {
	return "This action adds all moto"
}

func findOne(id string) string {
	return fmt.Sprintf("This action returns a #%s auth", id)
}

func update(id string, data interface{}) string {
	return fmt.Sprintf("This action updates a #%s auth", id)
}

func remove(id string) string {
	return fmt.Sprintf("This action removes a #%s auth", id)
}
