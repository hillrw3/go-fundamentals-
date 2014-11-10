package main

// todo: declare a constant called "goStart" with the type of int and value of 2007
const goStart int = 2007

// todo: declare a variable called "spareChange" with the type of float32 and value of 3.50
var spareChange float32 = 3.50

// todo: declare a variable called "cost" with the type of float32 and a value of 0.25
var cost float32 = 0.25

// todo: declare a variable called "purchased" with the type of float32 and a value of spareChange / cost
var purchased float32 = spareChange / cost

// todo: declare a variable called "snack" with the type of string and a value of "pretzels",
// and another called "bang" with a value of "!!!",
// and another called "excitedSnack" with a value of snack + bang
// !!! Declare them writing 'var' and 'string' only once !!!
var (
  snack string = "pretzels"
  bang = "!!!"
  excitedSnack = snack + bang
)
