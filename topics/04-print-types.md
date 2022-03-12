# Print Tipleri (Print, Println and Printf)

Öncelikle Print, Println ve Printf fonksiyonları arasındaki farkları inceleyelim.

**Print** bize normal bir ekrana yazdırma işlemi sunar.

```
func main() {
	fmt.Print("x")
	fmt.Print("y")
}
```

Çıktı:
```
xy
```

**Println** yazdırma işlemini gerçekleştirdikten sonra bir alt satıra geçer.

```
func main() {
	fmt.Println("x")
	fmt.Println("y")
}
```

Çıktı:
```
x
y
```

**Printf** yazdırma işlemini formatlanmış bir şekilde kullanmamıza olanak tanır.

```
func main() {
	world := "world"
	fmt.Printf("Hello %v", world)
}
```

Çıktı:
```
Hello world
```

Printf iki farklı parametre alır (format string, a ...interface{}). Bunlardan ilki format'tır. Bizim örneğimizde *%v*, ikinci parametre olarak da o format'a karşılık gelen veriyi/obj alır. Interface tipinde tanımlıdır. Şimdi gelin bu format tiplerini yakından inceleyelim...

**Genel:**

- **%v**  : varsayılan format'ı temsil eder. Yukarıda bunu kullanmıştık.
- **%+v** : bu flag ise struct'ları print ederken field'ları da görmemezi olanak tanır. 
- **%#v** : bu flag ise verilen değerin go-syntax temsilini bize verir.
- **%T**  : Veri tipini temsil eder.
- **%%**  : Printf kullanırken gerçek bir yüzde ifadesi için bu yapıyı kullanırız. 

Yukarıda öğrendiklerimizi örneklerle pekiştirelim:

```
func main() {
	var apple string = "apple"
	type employee struct {
		name string
		age  int
	}
	fmt.Printf("%v %T\n", apple, apple)
	fmt.Printf("%#v\n", apple)
	fmt.Printf("%T\n", employee{})
	fmt.Printf("%+v %T\n", employee{}, employee{})
	fmt.Printf("100 %% 6 : %v\n", 100 % 6)
}
```

Çıktı:
```
apple string
"apple"
main.employee
{name: age:0} main.employee
100 % 6 : 4
```

**İnteger:**

- **%b**  : (base 2)verilen değerin ikilik tabanda karşılığını verir.
- **%c**  : (unicode)verilen değerin unicode kod karşılığını verir.
- **%d**  : (base 10)verilen değerin onluk tabanda karşılığını verir.
- **%o**  : (base 8)verilen değerin sekizlik tabanda karşılığını verir.
- **%x**  : (base 16)verilen değerin onaltılık tabanda karşılığını verir.(with lower-case)
- **%X**  : (base 16)verilen değerin onaltılık tabanda karşılığını verir.(with upper-case)
- **%U**  : (unicode)verilen değeri unicode formatında verir.

Yukarıda öğrendiklerimizi örneklerle pekiştirelim:

```
func main() {
	var number int = 105
	
	fmt.Printf("%b \n", number)
	fmt.Printf("%c \n", number)
	fmt.Printf("%d \n", number)
	fmt.Printf("%o \n", number)
	fmt.Printf("%x \n", number)
	fmt.Printf("%X \n", number)
	fmt.Printf("%U \n", number)
}
```

Çıktı:
```
1101001 
i      
105    
151    
69     
69     
U+0069 
```

**Slice**
- **%p**  : (base 16)Sıfırıncı elemanın adresini onaltılık tabanda gösterir.

```
func main() {
	myslc := []int{1,2,3}
	
	fmt.Printf("%v \n", myslc)
	fmt.Printf("%p \n", myslc)
}
```

Çıktı:
```
[1 2 3] 
0xc000014150 
```