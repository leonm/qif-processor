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

func getColumn(transaction []string, column string) string {
  for _,t := range transaction {
    if strings.HasPrefix(t,column) {
      return t
    }
  }
  return ""
}

func processQIF(input *os.File, output *os.File, processor processTransaction) {

  scanner := bufio.NewScanner(input)
  transaction := []string{}

  for scanner.Scan() {
    line := scanner.Text()

    if strings.HasPrefix(line, "!") {
      output.WriteString(line+"\n")
    } else if line == "^" {
      transaction = append(transaction,line)
      for _,t := range processor(transaction) {
        output.WriteString(t+"\n")
      }
      transaction = []string{}
    } else {
      transaction = append(transaction,line)
    }
  }
}
