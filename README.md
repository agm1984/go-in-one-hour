# Go in One Hour

## GoDoc

``` golang
// Shows information about fmt package, Println function
godoc fmt Println
```

## Variables, fmt.Println(), and fmt.Printf()

- names can contain alphanumeric (a-z, A-Z, 0-9)
- names can contain underscores (_)

`SWEET IMMUTABILITY:` Variables cannot be changed once set

`EXTRANEOUS DECLARATIONS:` Program will detonate if variable is declared but not used.

- Do not do arithmetic with float types (use `int`)

``` golang
// Int
const pi float64 = 3.14159265
fmt.Printf("%f \n", pi) // Use newlines with Printf or it won't skip to the next line
fmt.Printf("%f.3f \n", pi) // to 3 decimal places
var (
  varA = 2
  varB = 3
)
fmt.Printf("%d \n", 100) //returns 100 because %d is int

// String
// can use " or `
var myName string = "Adam"
fmt.Println(len(myName)) // returns 4
fmt.Println(myName + " is a robot")

fmt.Println("I like \n \n ")
fmt.Println("new lines")

// Boolean
var isOver40 bool = false
fmt.Printf("%T \n", pi) // return float64 (%T shows the Type)
fmt.Printf("%t \n", pi) // returns false because %t is bool

```



``` golang
fmt.Println("Hello World")
var age int = 33
var favNum float64 = 1.6180339
randNum := 1
fmt.Println(age, " ", favNum)
fmt.Println(age, favNum)
  ```

- Spaces are added automatically
- explicitly adding space adds a second one

``` golang
num1 := 1000
num2 := 337
fmt.Println(num1 + num2)
```

### Operators

`+` `-` `*` `/` `%`

#### Logical Operators

``` golang
fmt.Println("true && false =", true && false)
fmt.Println("true || false =", true || false)
fmt.Println("!true", !true)
```

#### Relational Operators

`==` `!=` `<` `>` `<=` `>=`

## Loops

``` golang
i := 1
for i <= 10 {
  fmt.Println(i)
  i++
}

for j := 0; j < 5; j++ {
  fmt.Println(j)
}
```

## If statements

``` golang
yourAge := 18
if yourAge >= 16 {
  fmt.Println("You can drive")
} else {
  fmt.Printlb("You can't drive!")
}

// Go exits when the first condition is true
// In this example, the else if is not evaluated
yourAge := 18
if yourAge >= 16 {
  fmt.Println("You can drive")
} else if yourAge >= 18 {
  fmt.Println("You can vote")
} else {
  fmt.Println("You can have fun")
}
```

``` golang
yourAge := 5
switch yourAge {
  case 16: fmt.Println("Go drive")
  case 18: fmt.Println("Go vote")
  default: fmt.Println("Go have fun")
}
```

## Arrays

``` golang
var favNums2[5] float64
favNums2[0] = 163
favnums2[1] = 78557
favnums2[2] = 691
favnums2[3] = 3.141
favnums2[4] = 1.618
```