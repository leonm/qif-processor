package main

import "testing"
import "reflect"

func TestNoSetValueByDefault(t *testing.T) {
  c := &MockContext{isSet:false, value:".*"}
  transaction := []string{"PINTERNET TRANSFER"}
  processor := NewSetValueProcessor(c);
  if !reflect.DeepEqual(processor(transaction),transaction) {
    t.Fail()
  }
}

func TestSetValue(t *testing.T) {
  c := &MockContext{isSet:true, value:"INT"}
  transaction := []string{"PINTERNET TRANSFER"}
  processor := NewSetValueProcessor(c);
  processor(transaction)
  if !reflect.DeepEqual(processor(transaction),[]string{"PINT","CINT"}) {
    t.Fail()
  }
}
