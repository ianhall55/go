package main

import (
  "log"
  "os"
  "text/template"
  "strings"
)

var fm = template.FuncMap{
  "uc": strings.ToUpper,
  "ft": firstThree,
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.New("").Funcs(fm).ParseFiles("temp.gohtml"))
  // tpl = template.Must(template.Funcs(fm).ParseFiles("temp.gohtml"))
}

func firstThree(s string) string {
  s = strings.TrimSpace(s)
  s = s[:3]
  return s
}

type sage struct {
  Name string
  Motto string
}

func main() {

  b := sage{
    Name: "Buddha",
    Motto: "The belief of no beliefs",
  }

  g := sage{
    Name: "Ghandi",
    Motto: "Be the change",
  }

  sages := []sage{b, g}

  data := struct {
    Wisdom []sage
  }{
    sages,
  }

  err := tpl.ExecuteTemplate(os.Stdout, "temp.gohtml", data)
  if err != nil {
    log.Fatalln(err)
  }
}
