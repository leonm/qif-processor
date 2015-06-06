package main

import "io/ioutil"
import "bytes"
import "testing"

func assertEqualFiles(file1 string, file2 string, t *testing.T) {
  f1, err := ioutil.ReadFile(file1)
  check(err)

  f2, err := ioutil.ReadFile(file2)
  check(err)

  if (!bytes.Equal(f1, f2)) {
    t.Fail()
  }
}
