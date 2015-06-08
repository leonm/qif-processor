package main

import "regexp"
import "github.com/codegangsta/cli"
import "os"

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


func DeleteCommand(c *cli.Context)  {
  processor := NewDeleteProcessor(c)
  processQIF(os.Stdin, os.Stdout, processor)
}
