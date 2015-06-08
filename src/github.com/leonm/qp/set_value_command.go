package main

import "regexp"
import "github.com/codegangsta/cli"
import "os"

func NewSetValueProcessor(c CommandContext) processTransaction {

  payeeMatch := func(transaction []string) bool {return false }

  if(c.IsSet("payee")) {
    regex,_ := regexp.Compile(c.String("payee"))
    payeeMatch = func(transaction []string) bool {
      _,payee := getColumn(transaction,"P")
      return regex.MatchString(payee)
    }
  }

  processor := func (transaction []string) []string {
    if (payeeMatch(transaction)) {
      if c.IsSet("p") { transaction = setColumn(transaction,"P",c.String("p"))}
      if c.IsSet("l") { transaction = setColumn(transaction,"C",c.String("l"))}
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
