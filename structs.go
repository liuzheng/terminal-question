package terminal_question

type Question struct {
  ID       string
  Question string
  Retry    int
  Checker  func(string) (bool, error)
}
