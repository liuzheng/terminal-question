package terminal_question

import "golang.org/x/crypto/ssh/terminal"

type Question struct {
  ID       string
  Question string
  Retry    int
  Checker  func(string, *terminal.Terminal) bool
}
