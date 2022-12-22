package main

import (
	"log"

	"github.com/beomdevops/go-grpc/client/service"
	pb "github.com/beomdevops/go-grpc/proto"
	fiber "github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pb.NewGreeterClient(conn)
	helloSrv := service.NewHelloService(client)

	app := fiber.New()

	app.Get("/:test", func(c *fiber.Ctx) error {
		text := c.Params("test")
		if res := helloSrv.GetGrpcMessage(text); res != "" {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": res,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "err",
		})
	})
	log.Fatal(app.Listen(":3000"))

}
