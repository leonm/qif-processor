package main

import "github.com/codegangsta/cli"
import "os"

func processDelete(c *cli.Context)  {

  processor := func (transaction []string) []string {
    return transaction
  }

  processQIF(os.Stdin, os.Stdout, processor)
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
      Flags: []cli.Flag {
        cli.StringFlag{
          Name: "payee, P",
        },
      },
    },
  }

  app.Run(os.Args)
}
