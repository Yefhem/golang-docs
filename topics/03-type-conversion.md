# Tip Dönüşümleri (Type Conversion)

Programla sırasında pek çok nedenden dolayı tip dönüşümü yapmak zorunda kalabiliriz. Bu json'dan gelen string verinin integer'a dönüştürülmesi, uint8 bir değer ile uint16 olan değerin toplanması vs.. tip dönüşümü yapma sebeplerimizdendir...

```
func main() {
	var x uint8 = 50
	var y uint16 = 500

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", x+y, x+y)
}
```

Çıktı:

```
mismatched types uint8 and uint16
```

Böyle bir durumda hata alırız çünkü uint8 ile uint16 temelde ikisi de int olmasına rağmen farklı veri tipleridir.

Go dilinde data tiplerini birbirine dönüştürmenin çeşitli yolları mevcuttur. Bunlardan bazılarını örneklerle birlikte inceleyelim.

İlk olarak alışılmış yöntemlerden birini kullanarak başlayalım. Explicit type conversion olarak da bilinen bu yöntemle tip dönüşümü yapmak istersek, dönüştürmek istediğimiz veri tipini belirtip ardından parentez içinde değişkenimizi veririz.

**type(value)**

```
func main() {
	a := 20
	b := 20.1

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)

	// type(value) > int(b) = 20.1 > 20

	fmt.Printf("%v %T\n", int(b), int(b))
}
```

Çıktı:

```
20 int
20.1 float64
20 int      
```

```
func main() {
	var x int8 = 99
	var y int16 = 600

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", int16(x)+y, int16(x)+y)
}
``` 

Çıktı:

```
99 int8
600 int16
699 int16    
```

**Not:** Tip dönüşümü yaparken dönüşüm genelde daha geniş veri kapsayan veri tipine yapılır. Örneğin int16 ile int32 arasında bir dönüşüm yapılacaksa veri kaybına yol açmamak için int16'yı int32'ye dönüştürmek daha sağlıklı bir yöntemdir.

Bu tarz tip dönüşümlerinde string veriyi integer'a dönüştüremiyoruz. Bunun için başka yöntemler mevcut.

```
func main() {
	x := 10
	y := "10"

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Println(x + int(y))
}
```

Çıktı:

```
cannot convert y (type string) to type int  
```

Şöyle harici durumlar da olabilir..

```
func main() {
	number := 101
	string1 := string(number)

	fmt.Printf("%v %T\n", number, number)
	fmt.Printf("%v %T\n", string1, string1)
}
```

Çıktı:

```
101 int
e string
```