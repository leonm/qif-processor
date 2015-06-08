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
    {
      Name:    "set-value",
      Aliases: []string{"s"},
      Usage:   "sets a column value on a transaction",
      Action:  SetValueCommand,
      Flags: []cli.Flag {
        cli.StringFlag{
          Name: "payee, P",
        },
        cli.StringFlag{
          Name: "set-payee, p",
        },
        cli.StringFlag{
          Name: "set-category, l",
        },
      },
    },
  }

  app.Run(os.Args)
}
