package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string) (string, error) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(prompt)
  input, err := reader.ReadString('\n')
  return strings.TrimSpace(input), err
}

func createBill() bill {
  name, _ := getInput("Create a new bill name: ")
  b := newBill(name)
  fmt.Println("Created the bill - ", b.name)
  return b
}

func promptOptions(b bill) string {
  opt, _ := getInput("Choose option (a - add item, s - save item, t - add tip): ")

  if opt == "a"{
    name, _ := getInput("Name of item : ")
    price, _ := getInput("Price of item($) : ")
    p, err := strconv.ParseFloat(price, 64)
    if err != nil{
      fmt.Println("Yuu can only pass in a number")
    }
    b.addItem(name, p)
    fmt.Println("Item added - ", name, p)
    promptOptions(b)
  }else if opt == "t"{
    tip, _ := getInput("Tip amount: ")
    t, err := strconv.ParseFloat(tip, 64)
    if err != nil{
      fmt.Println("Tip must be a number") 
    }
    fmt.Println("Tip added", t)
    b.updateTip(t)
    promptOptions(b)
  }else if opt == "s"{
    fmt.Println("Bill Saved!")
    fmt.Println("--------------------------------")
    b.save()
    fmt.Println("--------------------------------")
  }else{
    fmt.Println("Invalid command...")
    promptOptions(b)
  }
  return opt
}


func main() {
  myBill :=  createBill()
  promptOptions(myBill)
}
