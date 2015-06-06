package main

import "bufio"
import "os"
import "strings"

func processQIF(input *os.File, output *os.File) {

  scanner := bufio.NewScanner(input)
  transaction := []string{}

  for scanner.Scan() {
    line := scanner.Text()

    if strings.HasPrefix(line, "!") {
      output.WriteString(line+"\n")
    } else if line == "^" {
      for _,t := range transaction {
        output.WriteString(t+"\n")
      }
      output.WriteString("^\n")
    } else {
      transaction = append(transaction,line)
    }
  }
}
