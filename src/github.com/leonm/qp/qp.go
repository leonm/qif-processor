package main

import "github.com/codegangsta/cli"
import "os"
import "strings"
import "regexp"

func getColumn(transaction []string, column string) string {
  for _,t := range transaction {
    if strings.HasPrefix(t,column) {
      return t
    }
  }
  return ""
}

func NewDeleteProcessor(c CommandContext) processTransaction {

  payeeMatch := func(transaction []string) bool {return false }

  if(c.IsSet("payee")) {
    regex,_ := regexp.Compile(c.String("payee"))
    payeeMatch = func(transaction []string) bool {
      return regex.MatchString(getColumn(transaction,"P"))
    }
  }

  processor := func (transaction []string) []string {
    if (payeeMatch(transaction)) {
      return nil
    }
    return transaction
  }

  return processor
}


func processDelete(c *cli.Context)  {
  processor := NewDeleteProcessor(c)
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
