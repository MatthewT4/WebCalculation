package http

import (
	"awesomeProject2/internal/blogic"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Router struct {
	blogic blogic.IBlogic
}

func NewRouter() *Router {
	return &Router{blogic: &blogic.Blogic{}}
}

func (r *Router) Start() {
	app := fiber.New()

	app.Get("/api/plus", r.Plus)

	app.Listen(":80")
}

func (r *Router) Plus(c *fiber.Ctx) error {
	if len(c.Query("first")) == 0 {
		c.Status(400)
		c.Write([]byte("argument 'first' not found"))
	}

	if len(c.Query("second")) == 0 {
		c.Status(400)
		c.SendString("argument 'second' not found")
	}

	stFirst := c.Query("first")
	first, fErr := strconv.ParseFloat(stFirst, 64)

	if fErr != nil {
		c.Status(400)
		c.SendString("argument 'first' is not float")
	}

	stSecond := c.Query("second")
	second, sErr := strconv.ParseFloat(stSecond, 64)

	if sErr != nil {
		c.Status(400)
		c.SendString("argument 'second' is not float")
	}
	result := r.blogic.Plus(first, second)
	c.SendString(fmt.Sprint(result))
	return nil
}

func float64ToByte(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
