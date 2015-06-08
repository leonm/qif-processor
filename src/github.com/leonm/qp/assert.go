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

type MockContext struct {
  isSet bool
  value string
}

func (c *MockContext) IsSet(string) bool {
  return c.isSet
}

func (c *MockContext) String(string) string {
  return c.value
}
