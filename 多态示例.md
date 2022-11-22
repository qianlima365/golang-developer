在go语言中，没有类的概念，但是仍然可以用**struct+interface**来实现类的功能，下面的这个简单的例子演示了如何用go来模拟c++中的多态的行为。

```
package main

import "os"
import "fmt"

type Human interface { 
  sayHello()
} 
type Chinese struct { 
  name string
}
 
type English struct { 
  name string
}

func (c *Chinese) sayHello() { 
  fmt.Println(c.name,"说：你好，世界")
}
 
func (e *English) sayHello() { 
  fmt.Println(e.name,"says: hello,world")
}

func main() {
  fmt.Println(len(os.Args))
  c := Chinese{"汪星人"}
  e := English{"jorn"} 
  m := map[int]Human{}
  m[0] = &c
  m[1] = &e
  for i:=0;i<2;i++ {
   m[i].sayHello()
  }
}
```
