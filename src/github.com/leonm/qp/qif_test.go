package main

import "testing"
import "os"


func RunQIFSample(t *testing.T, name string, processor processTransaction) {
  in,err := os.Open(os.Getenv("GOPATH")+"/samples/"+name+".in.qif")
  check(err)
  defer in.Close()

  out,err := os.Create("/tmp/result.qif")
  check(err)
  defer out.Close()
  defer os.Remove("/tmp/result.qif")

  processQIF(in, out, processor)

  assertEqualFiles("/tmp/result.qif",os.Getenv("GOPATH")+"/samples/"+name+".out.qif",t)
}

func TestQIFHeader(t *testing.T) {
  RunQIFSample(t, "headers", RepeatTransaction)
}


func TestRepeatTransaction(t *testing.T) {
  RunQIFSample(t, "transaction", RepeatTransaction)
}

func TestRepeat2Transactions(t *testing.T) {
  RunQIFSample(t, "2transaction", RepeatTransaction)
}

func TestRemoveTransactions(t *testing.T) {
  RunQIFSample(t, "delete-transaction", DeleteTransaction)
}

func TestGetTransactionColumn(t *testing.T) {
  transaction := []string{"PINTERNET TRANSFER"}
  if i,c := getColumn(transaction,"P"); i != 0 || c != "PINTERNET TRANSFER" {
    t.Fail()
  }
}

func TestGetNonExistantTransactionColumn(t *testing.T) {
  transaction := []string{"PINTERNET TRANSFER"}
  if i,c := getColumn(transaction,"C"); i != -1 || c != "" {
    t.Fail()
  }
}

func TestSetTransactionColumn(t *testing.T) {
  transaction := []string{"PINTERNET TRANSFER"}
  transaction = setColumn(transaction,"P","TEST")
  if transaction[0] != "PTEST" {
    t.Fail()
  }
}

func TestSetNewTransactionColumn(t *testing.T) {
  transaction := []string{"PINTERNET TRANSFER"}
  transaction = setColumn(transaction,"C","TEST")
  if transaction[0] != "PINTERNET TRANSFER" {
    t.Fail()
  }
  if transaction[1] != "CTEST" {
    t.Fail()
  }
}
