package main

import (
	"fmt"
	"os"
)

type bill struct{
  name string
  items map[string]float64
  tip float64
}

func newBill(name string) bill {
  b := bill{
    name: name,
    items: map[string]float64{},
    tip: 0,
  }
  return b
}

// Format bill

func (b bill) format() string{
  fs := "Bill breakdown: \n"
  var total float64 = 0
  for k, v := range b.items{
    fs += fmt.Sprintf("%v:       ...$%0.2f\n", k , v)
    total += v
  }
  
  fs += fmt.Sprintf("Tip:         ....$%0.2f\n", b.tip)
  total += b.tip
  fmt.Println("--------------------------------")
  fs += fmt.Sprintf("\nTotal:       ....$%0.2f\n", total)
  return fs
}

func (b *bill) updateTip(tip float64){
  b.tip = tip
}

func (b *bill) addItem(name string, price float64){
  b.items[name] = price
}

func (b *bill) addMultipleItem(items map[string]float64){
  b.items = items
}

func (b *bill) save(){
  data := []byte(b.format())
  err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
  if err != nil{
    fmt.Println(err)
  }else{
    fmt.Println("Bill saved to bills folder")
  }
}