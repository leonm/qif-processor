package main

import "testing"
import "os"


func TestQIFHeader(t *testing.T) {
  in,err := os.Open(os.Getenv("GOPATH")+"/samples/headers.in.qif")
  check(err)
  defer in.Close()

  out,err := os.Create("/tmp/result.qif")
  check(err)
  defer out.Close()
  defer os.Remove("/tmp/result.qif")

  processQIF(in,out)

  assertEqualFiles("/tmp/result.qif",os.Getenv("GOPATH")+"/samples/headers.out.qif",t)
}


func TestRepeatTransactions(t *testing.T) {
  in,err := os.Open(os.Getenv("GOPATH")+"/samples/transaction.in.qif")
  check(err)
  defer in.Close()

  out,err := os.Create("/tmp/result.qif")
  check(err)
  defer out.Close()
  // defer os.Remove("/tmp/result.qif")

  processQIF(in,out)

  assertEqualFiles("/tmp/result.qif",os.Getenv("GOPATH")+"/samples/transaction.out.qif",t)
}
