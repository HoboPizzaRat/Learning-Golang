package main

// Go supports the ability to return early from a function.

// you can use it make guard clauses that return early
// when given conditions are not met or errors occur

// example of guard clause function
/*
func getInsuranceAmount(status insuranceStatus) int {
  if !status.hasInsurance(){
    return 1
  }
  if status.isTotaled(){
    return 10000
  }
  if !status.isDented(){
    return 0
  }
  if status.isBigDent(){
    return 270
  }
  return 160
}
*/
