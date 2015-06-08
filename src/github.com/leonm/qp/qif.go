package main

import "bufio"
import "os"
import "strings"

type processTransaction func([]string) []string

func RepeatTransaction (transaction []string) []string {
  return transaction
}

func DeleteTransaction (transaction []string) []string {
  return nil
}

func getColumn(transaction []string, column string) (int, string) {
  for i,t := range transaction {
    if strings.HasPrefix(t,column) {
      return i,t
    }
  }
  return -1, ""
}

func setColumn(transaction []string, column string, value string) []string{
  i,_ := getColumn(transaction, column)
  if i == -1 {
    transaction = append(transaction,column+value)
    return transaction
  } else {
    newTransaction := make([]string,len(transaction),len(transaction))
    copy(newTransaction, transaction)
    newTransaction[i] = column+value
    return newTransaction
  }
}

func processQIF(input *os.File, output *os.File, processor processTransaction) {

  scanner := bufio.NewScanner(input)
  transaction := []string{}

  for scanner.Scan() {
    line := scanner.Text()

    if strings.HasPrefix(line, "!") {
      output.WriteString(line+"\n")
    } else if line == "^" {
      transaction = processor(transaction)
      if transaction != nil {
        for _,t := range processor(transaction) {
          output.WriteString(t+"\n")
        }
        output.WriteString("^\n")
      }
      transaction = []string{}
    } else {
      transaction = append(transaction,line)
    }
  }
}
