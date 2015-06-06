package main

import "github.com/codegangsta/cli"
import "os"

func processDelete(c *cli.Context)  {
  processQIF(os.Stdin, os.Stdout)
}

func main() {
  app := cli.NewApp()

  app.Name = "QIF Processor"

  app.Commands = []cli.Command{
    {
      Name:    "delete",
      Aliases: []string{"d"},
      Usage:   "delete a transaction",
      Action:  processDelete,
    },
  }

  app.Run(os.Args)
}
