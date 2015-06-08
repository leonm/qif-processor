package main

import "regexp"
import "github.com/codegangsta/cli"
import "os"

func NewSetValueProcessor(c CommandContext) processTransaction {

  payeeMatch := func(transaction []string) bool {return false }

  if(c.IsSet("payee")) {
    regex,_ := regexp.Compile(c.String("payee"))
    payeeMatch = func(transaction []string) bool {
      return regex.MatchString(getColumn(transaction,"P"))
    }
  }

  processor := func (transaction []string) []string {
    if (payeeMatch(transaction)) {
      return transaction
    }
    return transaction
  }

  return processor
}


func SetValueCommand(c *cli.Context)  {
  processor := NewSetValueProcessor(c)
  processQIF(os.Stdin, os.Stdout, processor)
}
