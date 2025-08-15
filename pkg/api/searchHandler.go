package api
import (
	"github.com/gofiber/fiber/v2"	
	"net/http"
	"net/url"
    "io/ioutil"
)

func SemanticSearch(c *fiber.Ctx) error {
        //valor por default
        threshold := "0.45"
		query := c.Query("query")
        queryEncoded := url.QueryEscape(query)

        resp, err := http.Get("http://localhost:8000/buscar?query="+queryEncoded+"&threshold="+threshold)
        
		if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }

        defer resp.Body.Close()

        body, _ := ioutil.ReadAll(resp.Body)
        return c.Send(body)
    }
