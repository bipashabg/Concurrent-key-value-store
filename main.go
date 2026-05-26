package main

import (
    "time"
    "NanoKV/kvstore"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    kv := kvstore.NewKeyValueStore()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("This is a Simple Key-Value store like Redis in Go.")
    })

    app.Get("/get/:key", func(c *fiber.Ctx) error {
        key := c.Params("key")
        value, ok := kv.Get(key)
        if !ok {
            return c.SendString("The Key " + key + " doesn't exist")
        }
        return c.SendString("The Key " + key + " has Value " + value)
    })

    app.Post("/set/:key/:value", func(c *fiber.Ctx) error {
        key := c.Params("key")
        value := c.Params("value")
        kv.Set(key, value, 10*time.Minute)
        return c.SendString("Key " + key + " Value " + value)
    })

    app.Delete("/delete/:key", func(c *fiber.Ctx) error {
        key := c.Params("key")
        ok := kv.Delete(key)
        if !ok {
            return c.SendString("The Key " + key + " doesn't exist")
        }
        return c.SendString("Successfully Deleted!!")
    })

    app.Listen(":3000")
}
