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

## Strconv Paketi Kullanımı 

**string - integer**

- Atoi > String veriyi int tipine dönüştürür ve dönüş olarak (int, error) döndürür.

- Itoa > İnt veriyi string tipine dönüştürür.

```
func main() {
	if n, err := strconv.Atoi("90"); err == nil {
		fmt.Printf("%v %T\n", n, n)
	}
	k := strconv.Itoa(90)
	fmt.Printf("%v %T\n", k, k)
}
```

Çıktı:

```
90 int
90 string
```

**String to other type**

String verileri diğer değişken tiplerine dönüştürmek için;

- ParseInt/ParseUint > string olarak gelen veriyi int/uint tipine dönüştürür. 3 parametre alır (s string, base int, bitSize int). "s" string olarak gönderdiğimiz veri, "base" sayı tabanını temsil eder ve "bitSize" hangi bit değerinde dönüştürme istiyorsak onu belirleriz. Dönüş olarak (int64/uint64, error) döndürür.

**Note:** 
- **s** string olarak gönderdiğimiz veriyi, 
- **base** sayı tabanını,
- **bitSize** bit değerini temsil eder.

```
func main() {
	i, err := strconv.ParseInt("-54", 10, 64)
	if err != nil {
		fmt.Printf("Error: ", err)
	}
	fmt.Printf("%v %T\n", i, i)

	u, err := strconv.ParseUint("40", 10, 64)
	if err != nil {
		fmt.Printf("Error: ", err)
	}
	fmt.Printf("%v %T\n", u, u)
}
```

Çıktı:

```
-54 int64
40 uint64
```

- ParseFloat > string olarak gelen veriyi float tipine dönüştürür. 2 paremetre alır (s string, bitSize int). Dönüş olarak (float64, error) döndürür. 

```
func main() {
	f, err := strconv.ParseFloat("3.31", 64)
		if err != nil {
			fmt.Printf("Error: ", err)
		}
		fmt.Printf("%v %T\n", f, f)
}
```

Çıktı:

```
3.31 float64
```

- ParseBool > string olarak verilen veriyi bool tipine dönüştürür. Tek parametre alır (str string). Dönüş olarak (bool, error) döndürür.

```
func main() {
	b, err := strconv.ParseBool("true")
		if err != nil {
			fmt.Printf("Error: ", err)
		}
		fmt.Printf("%v %T\n", b, b)
}
```

Çıktı:

```
true bool
```

**Other type to string**

- FormatInt/FormatUint > Int/Uint olarak verilen veriyi string tipine dönüştürür. 2 parametre alır (i int64, base int). Dönüş olarak string döndürür.

```
func main() {
	i := strconv.FormatInt(-13, 10)
	fmt.Printf("%v %T\n", i, i)

	u := strconv.FormatUint(88, 10)
	fmt.Printf("%v %T\n", u, u)
}
```

Çıktı:

```
-13 string
88 string
```

- FormatFloat > Float olarak verilen veriyi string tipine dönüştürür. 4 paremetre alır (f float64, fmt byte, prec, bitSize int).  Dönüş olarak string döndürür.

**Note:** 
- **f** float64 olarak gönderdiğimiz veriyi, 
- **fmt** 'E' (-d.ddddE±dd, a decimal exponent),
- **prec** precision,
- **bitSize** bit değerini temsil eder.

```
func main() {
	f := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("%v %T\n", f, f)
}
```

Çıktı:

```
3.1415E+00 string
```

- FormatBool > Bool olarak verilen veriyi string tipine dönüştürür. Tek parametre alır (b bool). Dönüş olarak string döndürür.

```
func main() {
	b := strconv.FormatBool(true)
	fmt.Printf("%v %T\n", b, b)
}
```

Çıktı:

```
true string
```