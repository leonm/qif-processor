package main

import "github.com/codegangsta/cli"
import "os"

type CommandContext interface {
    IsSet(string) bool
    String(string) string
}

func commandIn(c *cli.Context) *os.File {
  if (c.GlobalIsSet("input")) {
    f,err := os.Open(c.GlobalString("input"))
    check(err)
    return f;
  } else {
    return os.Stdin
  }
}

func commandOut(c *cli.Context) *os.File {
  if (c.GlobalIsSet("output")) {
    f,err := os.Create(c.GlobalString("output"))
    check(err)
    return f;
  } else {
    return os.Stdout
  }
}
