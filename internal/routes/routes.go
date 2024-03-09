package routes

import (
	"strconv"

	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/routes/request"
	"github.com/brunodelucasbarbosa/rinha-backend-2024q1/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func StartRoutes(s services.TransactionsService) {
	r := gin.Default()

	r.POST("/clientes/:id/transacoes", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		body := request.TransactionRequest{ClientId: id}

		if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			c.JSON(400, gin.H{
				"error": "invalid body",
			})
			return
		}

		res, err := s.CreateTransaction(body)

		if err != nil {
			c.JSON(err.Code, gin.H{
				"error": err.Message,
			})
			return
		}

		c.JSON(200, res)
	})

	r.GET("/clientes/:id/extrato", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		res, err := s.GetExtract(id)

		if err != nil {
			c.JSON(err.Code, gin.H{
				"error": err.Message,
			})
			return
		}

		c.JSON(200, res)
	})

	r.Run(":8080")
}
