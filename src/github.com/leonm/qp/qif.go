package main

import "bufio"
import "os"
import "strings"

type processTransaction func([]string) []string

func processQIF(input *os.File, output *os.File, processor processTransaction) {

  scanner := bufio.NewScanner(input)
  transaction := []string{}

  for scanner.Scan() {
    line := scanner.Text()

    if strings.HasPrefix(line, "!") {
      output.WriteString(line+"\n")
    } else if line == "^" {
      for _,t := range processor(transaction) {
        output.WriteString(t+"\n")
      }
      output.WriteString("^\n")
      transaction = []string{}
    } else {
      transaction = append(transaction,line)
    }
  }
}
