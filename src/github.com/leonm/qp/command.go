package main

type CommandContext interface {
    IsSet(string) bool
    String(string) string
}
