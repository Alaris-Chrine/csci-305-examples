package main

import "fmt"
     

func main(){
    //var declares 1 or more variables.
    var a = "initial"
    fmt.Println(a)
    aaa := "second String"
    fmt.Println(aaa)
   // a  = 7     //can't re declard a var
    var mine  int = 8
    fmt.Println(mine)
//You can declare multiple variables at once.
    var b, c int = 1, 2
    var x, y = 1, "you"
    fmt.Println(b, c)
    fmt.Println(x, y)



//Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)
  


   //Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
    var e,m,n int
    fmt.Println(e, m, n)
    

    //The := syntax is shorthand for declaring and initializing a variable, 
    //e.g. for var f string = "short" in this case.

    f := "short"
    fmt.Println(f)
}
