package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	port = flag.String("port", ":1337", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in producion")
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	app.Post("/img/upload", func(ctx *fiber.Ctx) error {
		file, err := ctx.FormFile("document")
		if err != nil {
			return err
		}

		err = os.Mkdir("img", 0755)
		if err != nil {
			log.Println(err)
		}

		url := RandStringRunes(7)

		err = ctx.SaveFile(file, fmt.Sprintf("./img/%s", url))
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			ctx.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"url": fmt.Sprintf("https://img.zackmyers.io/%s", url),
		})
	})

	app.Static("/", "./img")

	log.Fatal(app.Listen(*port))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
