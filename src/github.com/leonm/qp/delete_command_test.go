package main

import "testing"
import "reflect"

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

func TestNoDeletionByDefault(t *testing.T) {
  c := &MockContext{isSet:false, value:".*"}
  transaction := []string{"PINTERNET TRANSFER"}
  processor := NewDeleteProcessor(c);
  if !reflect.DeepEqual(processor(transaction),transaction) {
    t.Fail()
  }
}

func TestDeleteIfPayeeMatches(t *testing.T)  {
  c := &MockContext{isSet:true, value:".*"}
  transaction := []string{"PINTERNET TRANSFER"}
  processor := NewDeleteProcessor(c);
  if processor(transaction) != nil {
    t.Fail()
  }
}

func TestNoDeletionIfPayeeNotMatches(t *testing.T)  {
  c := &MockContext{isSet:true, value:"_NOMATCH_"}
  transaction := []string{"PINTERNET TRANSFER"}
  processor := NewDeleteProcessor(c);
  if !reflect.DeepEqual(processor(transaction),transaction) {
    t.Fail()
  }
}
