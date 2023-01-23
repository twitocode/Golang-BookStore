package common

type Status struct {
  Code int
  Message string
}

type Response[T any]  struct {
  Code int
  Data T
}
