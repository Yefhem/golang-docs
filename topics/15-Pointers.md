# İşaretleyiciler (Pointers)

Bilindiği üzerine değişkenler hafızada "değişken adresi" ve "değeri" şeklinde tutulurlar. Bir değişkenin değerini değiştirmek istediğimizde hafıza o değişkenin adresine gider ve değerini değiştiririz.

Tam bu noktada bir belirteç olan pointer, hafızadaki değişkenin adresini almak veya bu adresteki değeri değiştirebilmemiz için kullanılmaktadır.

C, C++ tabanlı dillerde kullanılan bu yapı çoğunlukla referans tutmak için kullanılan bir yapıdır.

> & -->  değişken adresi bulmakta kullanılır. (Address Operator)

> \* --> pointer'in işaret ettiği değeri temsil eder.

**Not:** Ayrıca adresteki değeri geri döndürme işlemine de *dereferencing* denilir.

```
func main() {
	x := 19

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*(&x))
}
```

Çıktı:
```
19
0xc000012088
19
```

**Example:**

```
func main() {
	x1 := 10
	x2 := &x1

	fmt.Println(x1, x2)  
	fmt.Println(x1, *x2) 

	*x2 = 3 

	fmt.Println("------")

	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)

	fmt.Println("------")

	x3 := &x1

	*x3 = 5

	fmt.Println(x1, *x2, *x3)
}
```

Çıktı:
```
10 0xc000012088
10 10
------        
3 0xc000012088
3 3
------        
5 5 5
```

## Pass by Value

```
func main() {
	x1 := 10 
	x2 := x1 
	fmt.Println(x1, x2)
	x1 = 5 
	fmt.Println(x1, x2)
}
```

Çıktı:
```
10 10
5 10
```

## Pass by Reference

```
func main() {
	x1 := 10 
	x2 := &x1 
	fmt.Println(x1, x2)
	fmt.Println(x1, *x2)

	*x2 = 3
	fmt.Println(x1, *x2)
}
```

Çıktı:
```
10 0xc000012088
10 10
3 3  
```

## Fonksiyon Parametresi Olarak Pointer Kullanımı

Normalde fonksiyon parametreleri değer kopyalaması yöntemi ile kullanılır.
Yani bir fonksiyona dışardan geçilen parametre değişkeni fonksiyon bloğu içerisine kopyalanarak değerlendirilir.

Bu nedenle fonksiyon çağrımı yapılandan farklı bir değişken üzerinde işlemler yapılır. Pointer kullanıldığında ise fonksiyona geçen parametrenin adresi taşınır.

```
func main() {
	x := 5  
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}

func double(num *int) {
	fmt.Println(num) 
	*num *= 2
	fmt.Println(*num)
}  
```

Çıktı:
```
5
0xc000012088
10
10  
```

## New Kullanımı ve Pointer

New fonksiyonu memory'de yer ayırır ancak make fonksiyonu gibi initialize etmez. Geriye bir pointer adres döner.

```
func main() {
	a := new(int)

	fmt.Printf("%T, %v\n", a, a)

	*a = 3
	fmt.Printf("%T, %v\n", *a, *a)
}
```

Çıktı:
```
*int, 0xc000012088
int, 3
```