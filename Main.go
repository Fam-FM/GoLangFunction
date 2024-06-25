package main
import "fmt"

func getName(name string) string {
  fmt.Println("What is your name?")
  fmt.Scan(&name)
  return name
}

func getAge() int {
  var age int
  fmt.Println("What is your age?")
  fmt.Scan(&age)

   for age <= 0 || age > 100 {
    fmt.Println("That cannot be, please enter your age again")
    fmt.Scan(&age)
  }
  return age
}

func printTicket(age int, ticketPrice float64) float64 {
 if age <= 13 {
    ticketPrice = 9.99
 } else if age > 13 && age < 65 {
    ticketPrice = 19.99 
 } else if age >= 65 {
    ticketPrice = 12.99
 }
  return ticketPrice
}

func main() {
  var name string
  var ticketprice float64

  defer fmt.Println("Your name is:", getName(name))
  var age int = getAge()
  defer fmt.Println("Your age is:", age)
  defer fmt.Println("Your ticket price is:", printTicket(age, ticketprice))
}
