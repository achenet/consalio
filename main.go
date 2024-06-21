package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// consalio
// now with more health gage
type Project struct {
	Id int
}

// blue - This project is doing better than expected. You're either very lucky, doing something right, or both. Celebrate!
// green - on time, on budget :)
// orange - uh oh... some overruns
// red - CRITICAL PROBLEMS WITH THIS PROJECT DO SOMETHING NOW
func main() {
	fmt.Printf("Health check of project: \n", HealthCheck(1))
	r := gin.Default()

	r.GET("/:id", func(c *gin.Context) {
		id := c.Params("id")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Run()
}

// I don't actually know how Invoices work
type Invoice struct {
	Total float64
}

func ProjectInvoices(id int) []Invoice {
	switch {
	default:
		return []Invoice{{7}}
	}
}

var expectedSpend = map[int]float64{
	1: 100,
	2: 100,
	3: 5,
	4: 10,
}

type Health uint

const (
	BLUE Health = iota
	GREEN
	ORANGE
	RED
)

func HealthCheck(id int) Health {
	invoices := ProjectInvoices(id)
	spent := 0.0
	for _, inv := range invoices {
		spent += inv.Total
	}
	expect := expectedSpend[id]
	if spent < .5*expect {
		return BLUE
	}
	if spent > expect && spent < 1.5*expect {
		return ORANGE
	}
	if spent > 1.5*expect {
		return RED
	}
	return GREEN
}
