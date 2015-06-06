package main

import "testing"
import "os"


func RunQIFSample(t *testing.T, name string) {
  in,err := os.Open(os.Getenv("GOPATH")+"/samples/"+name+".in.qif")
  check(err)
  defer in.Close()

  out,err := os.Create("/tmp/result.qif")
  check(err)
  defer out.Close()
  defer os.Remove("/tmp/result.qif")

  processQIF(in,out, func(transaction []string) []string {return transaction} )

  assertEqualFiles("/tmp/result.qif",os.Getenv("GOPATH")+"/samples/"+name+".out.qif",t)
}

func TestQIFHeader(t *testing.T) {
  RunQIFSample(t, "headers")
}


func TestRepeatTransaction(t *testing.T) {
  RunQIFSample(t, "transaction")
}

func TestRepeat2Transactions(t *testing.T) {
  RunQIFSample(t, "2transaction")
}
