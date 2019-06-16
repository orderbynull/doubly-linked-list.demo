**Пример использования:**
```
package main

import "fmt"
import "github.com/orderbynull/doubly-linked-list.demo"

func printList(list dllist.List) {
	item := list.First()
	for {
		if item == nil {
			break
		}

		fmt.Printf(" %d ", item.Value())

		item = item.Next()
	}
}

func main() {
	var list dllist.List


	for i := 1; i <= 10; i++ {
		list.PushBack(i)
	}

	fmt.Printf("Значения после инициализации списка: ")
	printList(list)

	fmt.Printf("\nЗначения после удаления первого элемента: ")
	list.First().Remove()
	printList(list)

	fmt.Printf("\nЗначения после удаления последнего элемента: ")
	list.Last().Remove()
	printList(list)

	fmt.Printf("\nЗначения после удаления третьего элемента: ")
	list.First().Next().Next().Remove()
	printList(list)

	fmt.Printf("\nПервый элемент: %d", list.First().Value())
	fmt.Printf("\nПоследний элемент: %d", list.Last().Value())
}

```

**Вывод:**
```
Значения после инициализации списка:  1  2  3  4  5  6  7  8  9  10 
Значения после удаления первого элемента:  2  3  4  5  6  7  8  9  10 
Значения после удаления последнего элемента:  2  3  4  5  6  7  8  9 
Значения после удаления третьего элемента:  2  3  5  6  7  8  9 
Первый элемент: 2
Последний элемент: 9
```

**Playground:** https://play.golang.org/p/CR32pniGKCu