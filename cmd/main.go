package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pichead/go-fiber-shorturl/configs"
)

type Url struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Key       string    `gorm:"uniqueIndex;not null" json:"key"`
	Link      string    `gorm:"not null" json:"link"`
	Exp       int       `json:"exp"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func main() {

	config := configs.LoadConfig()

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "3333"
	}

	app := fiber.New(fiber.Config{
		Prefork:      false, // ถ้า true, ใช้ multi-process mode
		ServerHeader: "Fiber",
		AppName:      "GO Fiber Short URL Service",
	})

	app.Static("/public-file/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GO Fiber Short Url Service")
	})

	// appPort := "3333"

	app.Listen(":" + config.App.Port)

}
