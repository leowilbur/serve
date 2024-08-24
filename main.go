package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Static File Server",
		Usage: "Serve a folder over HTTP",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "directory",
				Aliases: []string{"d"},
				Usage:   "Path to the folder to serve. Defaults to the current folder",
				Value:   "build", // Default to current directory
			},
			&cli.StringFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Port to serve on. Defaults to 3000",
				Value:   "3000", // Default to port 3000
			},
		},
		Action: func(c *cli.Context) error {
			directory := c.String("directory")
			port := c.String("port")

			gin.SetMode(gin.ReleaseMode)
			r := gin.Default()

			r.Static("/", "./"+directory)

			r.NoRoute(func(c *gin.Context) {
				c.File("./" + directory + "/index.html")
			})

			// Start the server on the specified port
			fmt.Printf("Serving %s on http://localhost:%s", directory, port)
			return r.Run(":" + port)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
