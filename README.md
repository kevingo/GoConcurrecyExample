這個專案主要用來學習 golang 的 concurrency.

首先我們可以看到下面這段程式碼會印出 Foo0~Foo9，接著是 Bar0~Bar9，相當直覺，應該沒有什麼問題：

```

```

而在 golang 中，想要 concurrecy 執行 foo() 和 bar()，要怎麼做呢？

### goroutine

goroutine 是 go 提供的一種功能，讓你可以在 go code 裡面實現 concurrency。使用的方法相當容易，只要使用 `go` 這個關鍵字即可，來看程式碼：

```
package main

import (
	"fmt"
)

func main() {
	go foo()
	go bar()
	fmt.Println("All Done.")
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
}

```

在這段程式碼中，我們在 foo() 和 bar() 前面加上了 `go` 宣告，這樣一來就完成了，你可以想像我們現在有三個 goroutine 被呼叫，分別是：main、foo 和 bar。

接著，你可以用 `go run goroutine.go` 來執行這段程式，你會預期它的執行結果應該是隨機的印出 Foo 和 Bar，最後再印出 All Done.

但是你會發現好像不是這樣，就像在我們電腦上，我會印出類似如下的結果：

```
$ go run goroutine.go
Foo: 0
Bar: 0
All Done.
Foo: 1
Bar: 1
Foo: 2
Bar: 2
```

你多跑幾次，也會發現執行結果不同：

```
$ go run goroutine.go
All Done.
Foo: 0
Bar: 0
Foo: 1
Bar: 1
```

這是為什麼呢？

### channel