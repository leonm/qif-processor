package main

import "github.com/codegangsta/cli"
import "os"


func main() {
  app := cli.NewApp()

  app.Name = "QIF Processor"

  app.Commands = []cli.Command{
    {
      Name:    "delete",
      Aliases: []string{"d"},
      Usage:   "delete a transaction",
      Action:  DeleteCommand,
      Flags: []cli.Flag {
        cli.StringFlag{
          Name: "payee, P",
        },
      },
    },
  }

  app.Run(os.Args)
}
