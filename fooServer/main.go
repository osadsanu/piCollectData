package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type device struct {
	id       string
	status   string
	actionID string
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func main() {
	r := gin.Default()

	var fakeID = StringWithCharset(1, "0123")
	var fakeStatus = StringWithCharset(1, "0123")
	var actionID = StringWithCharset(9, charset)
	r.GET("/", func(c *gin.Context) {
		if randomInt(0, 100) > 99 {
			fakeID = StringWithCharset(1, "0123")
			fakeStatus = StringWithCharset(1, "0123")
			actionID = StringWithCharset(9, charset)
		}
		data := gin.H{
			"id":       fakeID,
			"status":   fakeStatus,
			"actionID": actionID,
		}
		c.JSON(http.StatusOK, data)
	})
	fmt.Println("test", &device{"1", "4", "ASCFDAS123"})
	r.Run(":80")
}

// Returns an int >= min, < max
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
