# Slices

Slice'lar go'da dizilere göre daha güçlü, flex ve elverişlidir. Hafig bir veri yapısıdır. Slice'ların boyutu diziler gibi sabit değildir.
Slice da diziler gibi yinelenen öğeleri tutabilir. Slice ve dizi birbirine bağlıdır.
Bir slice temelde bir diziye referanstır.

Go'da nasıl slice oluşturabiliriz?

`var slice_name[]type`

```
func main() {
    var mySlice1 = []string{}
    
    var mySlice2 = []string{"red", "blue", "dark"}
    
    fmt.Println(mySlice1)
    
    fmt.Println(mySlice2)
}
```
Çıktı:
```
[]
[red blue dark]
```

### Diziden Slice'lar oluşturma

```
func main() {
	var colors = [6]string{"red", "blue", "yellow", "orange", "purple", "gray"}

	slice1 := colors[:6]
	slice2 := colors[0:]
	slice3 := colors[:]
	slice4 := colors[:3]
	slice5 := colors[2:4]

	fmt.Println(colors)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
}
```

```
[red blue yellow orange purple gray]
[red blue yellow orange purple gray]
[red blue yellow orange purple gray]
[red blue yellow orange purple gray]
[red blue yellow]                   
[yellow orange]  
```

### Make fonksiyonu ile Slice oluşturma 

Make fonksiyonunu kullanarak da slice oluşturabiliriz.

`make([]type, lenght, capacity)`

Capacity(kapasite) opsiyoneldir. Genellikle make fonksiyonu boş slice oluşturmak için kullanılır.

```
func main() {
	var mySlice1 = make([]int, 5, 7)
	fmt.Printf("Slice 1: %v, Lenght: %d , Capacity: %d\n", mySlice1, len(mySlice1), cap(mySlice1))

	mySlice2 := make([]int, 4)
	fmt.Printf("Slice 2: %v, Lenght: %d , Capacity: %d\n", mySlice2, len(mySlice2), cap(mySlice2))
}
```

Çıktı:
```
Slice 1: [0 0 0 0 0], Lenght: 5 , Capacity: 7
Slice 2: [0 0 0 0], Lenght: 4 , Capacity: 4  
```

### Slice hakkında önemli noktalar

Slice'ın bir referans türü olduğunu zaten biliyoruz. Bu yüzden slice üzerinde yapılan değişiklikler 
başvurulan dizide de gerçekleşir. Yani slice'daki herhangi bir değişiklik direkt olarak diziye yansır.


```
func main() {
	var myArr = [6]int{1, 2, 3, 4, 5, 6}

	mySlice := myArr[:3]

	fmt.Printf("Original Array : %v\n", myArr)
	fmt.Printf("Original Slice : %v\n", mySlice)

	fmt.Println()

	mySlice[0] = 10
	mySlice[1] = 20
	mySlice[2] = 30

	fmt.Printf("New Array : %v\n", myArr)
	fmt.Printf("New Slice : %v\n", mySlice)
}
```

Çıktı:
```
Original Array : [1 2 3 4 5 6]
Original Slice : [1 2 3]    
                            
New Array : [10 20 30 4 5 6]
New Slice : [10 20 30]  
```

Slice'ın nil olup olmadığını == operatörü kullanarak kontrol edebiliyoruz.

```
func main() {
	mySlice1 := []int{5, 10, 15}
	var mySlice2 []int

	fmt.Println(mySlice1 == nil)
	fmt.Println(mySlice2 == nil)
}
```

Çıktı:
```
false
true
```

**Not:** Ayrıca dizilerdeki iki slice karşılaştırmaya çalışırsak hata alırız.

Multi-dimensional slice(çok boyutlu slice) multi-dimensional diziye benzer.
Slice farklı olarak boyut içermez.

```
func main() {
	mySlice1 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	mySlice2 := [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}, {"g", "h"}, {"j", "z"}}

	fmt.Println("multi-dimensional slice 1 :", mySlice1)
	fmt.Println("multi-dimensional slice 1 :", mySlice2)
}
```

Çıktı:
```
multi-dimensional slice 1 : [[1 2] [3 4] [5 6]]
multi-dimensional slice 1 : [[a b] [c d] [e f] [g h] [j z]]
```

Slice'da bulunan öğeler sıralanabilir. Go standart kütüphanesi int, float64 ve string 
sıralamaya olanak tanır. Her zaman artan düzende sıralama yapar.

```
func main() {
	mySlice1 := []int{2, 5, 3, 7, 1, 9}
	mySlice2 := []string{"b", "a", "d", "c", "g", "e"}

	fmt.Println("Before Slice 1 :", mySlice1)
	fmt.Println("Before Slice 2 :", mySlice2)

	// sorting
	sort.Ints(mySlice1)
	sort.Strings(mySlice2)

	fmt.Println()

	fmt.Println("After Slice 1 :", mySlice1)
	fmt.Println("After Slice 2 :", mySlice2)
}
```

Çıktı:
```
Before Slice 1 : [2 5 3 7 1 9]
Before Slice 2 : [b a d c g e]
                              
After Slice 1 : [1 2 3 5 7 9] 
After Slice 2 : [a b c d e g] 
```

**Not:** IntsAreSorted fonksiyonu ise slice'ın artan sırada olup olmadığını kontrol eder, bool değer döndürür.