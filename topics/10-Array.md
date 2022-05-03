# Diziler (Arrays)

Aynı türden bir veri koleksiyonu saklamamız gerektiğinde dizileri kullanabiliriz.
Go'da  diziler diğer programlama dillerindeki yapıya çok benzerdir.

Go'da dizileri iki farklı yolla oluşturabiliriz:

* Var keyword kullanarak
* Shorthand declaration kullanarak  

1. Go'da bir dizi oluşturulurken boyut, tip ve elementleri belirterek oluşturabiliriz.

> var array_variable_name = [size]type{"elements", "elements" ...]

2. Shorthand declaration ise bize esnek bir kullanım sağlar.

> array_variable_name := [size]type{"elements", "elements" ...]

Bu kullanımlarda dizi elemanlarını sonradan ekleyebiliriz. Şimdi örneklerle pekiştirelim.

```
func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("Value: %v , Type: %T\n", numbers, numbers)

	true_or_false := [4]bool{true, false, false, true}

	fmt.Printf("Value: %v , Type: %T\n", true_or_false, true_or_false)

	var fruits [3]string

	fruits[0] = "orange"
	fruits[1] = "banana"
	fruits[2] = "apple"

	fmt.Printf("Value: %v , Type: %T\n", fruits, fruits)

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %v  Value: %v\n", i, fruits[i])
	}
}
```

Çıktı:
```
Value: [1 2 3 4 5] , Type: [5]int
Value: [true false false true] , Type: [4]bool
Value: [orange banana apple] , Type: [3]string
Index: 0  Value: orange                       
Index: 1  Value: banana                       
Index: 2  Value: apple   
```

**Dizilerin özellikleri:**

* Go'daki dizilerin uzunluğu sabittir ve değiştirilemez.
* Dizi içerisindeki elemanların index ve value'larına for loop ile ulaşılabilir.
* Go'daki dizi türü tek boyutludur(one-dimensional). Fakat çok boyutlu şekilde kullanım da yapabiliriz.

```
func main() {
	languages := [2][2]string{{"go", "c++"}, {"python", "ruby"}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(languages[i][j])
		}
	}

	fmt.Println()

	var numbers [2][2]int
	numbers[0][0] = 10
	numbers[0][1] = 20
	numbers[1][0] = 30
	numbers[1][1] = 40
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(numbers[i][j])
		}
	}
}
```

Çıktı:
```
go
c++   
python
ruby  
      
10    
20    
30    
40   
```

* Go'daki dizilerde yinelenen verileri tutmamıza olanak tanır.

```
func main() {
	var fruits = [3]string{"orange", "orange", "orange"}
	fmt.Println(fruits)
}
```

Çıktı:
`[orange orange orange]`

* Eğer bir dizi başlatılmışsa bu dizinin varsayılan değeri(zero value) dizi tipine göre farklılık gösterir.

```
func main() {
	var fruits [3]string
	fmt.Println(fruits)

	var numbers [3]int
	fmt.Println(numbers)

	var boolean [3]bool
	fmt.Println(boolean)
}   
```

Çıktı:
```
[  ]
[0 0 0]            
[false false false]
```

* Go'da dizilerin boyut bölümüne üç nokta konularak girilen eleman kadar boyutu olacağı belirtilebilir.

```
func main() {
	var example = [...]int{1, 2, 3, 4, 5}

	fmt.Printf("Value : %v , Lenght: %v", example, len(example))
}   
```

Çıktı:
`Value : [1 2 3 4 5] , Lenght: 5`

* Go'da dizier "value type" dır. "Reference type" değildir. Yani mevcut dizi  yeni bir diziye atandığında
yeni değişkende yapılan değişiklikler orijinal diziyi etkilemez.

```
func main() {
	var example2 = [...]int{5, 10, 15}

	fmt.Printf("Original array (before) : %v\n", example2)

	new_example_array := example2

	fmt.Printf("New array (before) : %v\n", new_example_array)

	new_example_array[0] = 50

	fmt.Printf("New array (after) : %v\n", new_example_array)

	fmt.Printf("Original array (after) : %v\n", example2)
}  
```

Çıktı:
```
Original array (before) : [5 10 15]
New array (before) : [5 10 15]    
New array (after) : [50 10 15]    
Original array (after) : [5 10 15]  
```

* Go'da iki diziyi doğrudan karşılaştırabiliriz. Tabi bunun için dizi tiplerinin aynı olması gerekir.

```
func main() {
	var arr1 = [...]string{"1", "2", "3"}
	var arr2 = [3]string{"1", "2", "3"}
	var arr3 = [3]string{"1", "2", "6"}
	// var arr4 = [3]int{1, 2, 3}

	fmt.Println(arr1 == arr2)
	fmt.Println(arr1 == arr3)
	fmt.Println(arr2 == arr3)

	/* 
	 *  Böyle bir durumda hata alırız.
	 *	fmt.Println(arr3==arr4)
	 */
}
```

Çıktı:
```
true
false
false
```