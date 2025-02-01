package utility

func Must[T interface{}](val T, e error) T {
  if e != nil {
    panic(e)
  }

  return val
}
