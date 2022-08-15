package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ledger(c *gin.Context) {
	command := c.Query("command")
	log.Println(command)

	parameters := strings.Fields(command)

	//cmd := []string{"b", "--flat", "--no-total"}
	result, err := runLedger(parameters...)

	if err != nil {
		log.Fatal(err)
	}

	var output = SplitLines(result)

	c.JSON(http.StatusOK, output)
}

// Returns an image.
// Call with a random parameter so it is not cached on the client-side.
func hello_img(c *gin.Context) {
	// Base64 encoded pixel
	pixelEncoded := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg=="

	sDec, err := base64.StdEncoding.DecodeString(pixelEncoded)
	if err != nil {
		msg := fmt.Sprintf("Error decoding string: %s ", err.Error())
		fmt.Println(msg)
		//return nil
		panic(msg)
	}

	byteReader := bytes.NewReader(sDec)
	var length int64 = int64(byteReader.Len())

	c.DataFromReader(http.StatusOK, length, "image/png", byteReader, nil)
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func shutdown(c *gin.Context) {
	os.Exit(1)
}
