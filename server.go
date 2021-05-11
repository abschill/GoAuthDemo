package main

import (
	"fmt"
"github.com/gofiber/fiber/v2"
)
type User struct{
	First string
	Last string
}
func handleLogin(c *fiber.Ctx) error{
	user := new(User)

    if err := c.BodyParser(user); err != nil {
      fmt.Println("error = ",err)
      return c.SendStatus(400)
    }
    
    // getting user if no error
    fmt.Println("user = ", user)
    fmt.Println("username:", user.First)
    fmt.Println("password:", user.Last)
  
    return c.JSON(user)
}
func main() {
	app := fiber.New()
	app.Static("/", "./public")
	app.Post("/login", handleLogin)
	app.Listen(":3000")
}
