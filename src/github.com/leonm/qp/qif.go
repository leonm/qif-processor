package main

import "bufio"
import "os"

func processQIF(input *os.File, output *os.File) {
  scanner := bufio.NewScanner(input)
  for scanner.Scan() {
    //println(scanner.Text())
    output.WriteString(scanner.Text()+"\n")
  }
}
