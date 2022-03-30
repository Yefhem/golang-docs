# Koşullu İfadeler (Conditional Statement)

Koşullu ifadelere geçmeden önce "statement" ve "expression" konularına kısaca değinelim.
İkisinin de sözlük anlamı "ifade" olarak geçer. Peki bu iki terim arasındaki fark nedir?

### Statement
Her bir satır kod *statement* olarak ifade edilir. Statement'lar programdaki komutları, eylemleri temsil etmektedir (Örneğin atama işlemleri, yazdırma işlemleri, değişken tanımlama vs..).
 
### Expression
Değer döndüren satırlar için kullanılır. Her expression aynı zamanda eylem gerçekleştirdiğinden birer statement'tır. Ancak her statement değer döndüremeyeceği için expression değildir.

**Not** Bu konu detaylı bir şekilde ilerleyen bölümlerde anlatılacaktır.

Koşullu ifadelere tekrardan dönüş yapacak olursak...

Koşullu ifadeler bir programla dilinin en önemli parçalarından biridir. Programlamada koşullu ifadeleri pek çok yerde kullanabiliriz. Bunlardan birkaçına örnek verecek olursak;

> Username ve password doğru ise homepage'e redirect işlemi yapılması.
> Maksimum parola denemesi var ise belirli bir süre login işleminin engellenmesi.
> Kullanıcı belirli miktarda ürün satın aldıysa indirim yapılması etc..

Go dilinde aşağıdaki koşullu ifadeler vardır:

- if statement
- if...else statement
- if...else if...else statement
- switch...case statement 

## if statement

`if <boolean expression> { koşul doğru ise buradaki kod yürütülür.. }`

```
func main() {
	x := true
	if x {
		fmt.Println(x)
	}
}
```

Çıktı:
```
true
```

## if...else statement

```
if  <boolean expression> { 
    koşul doğru ise buradaki kod yürütülür..
} else {
    koşul doğru değil ise buradaki kod yürütülür..
}
```

```
func main() {
	a := 28
	if a == 28 {
		fmt.Printf("a = %v", a)
	} else {
		fmt.Printf("a != %v", a)
	}
}
```

Çıktı:
```
a = 28
```

## if...else if...else statement

```
if  <boolean expression-1> { 
    1. koşul doğru ise buradaki kod yürütülür..
} else if <boolean expression-2> {
    2. koşul doğru ise buradaki kod yürütülür..
} else {
    yukarıdaki iki koşul da doğru değil ise buradaki kod yürütülür..
}
```

```
func main() {
	a := 15
	if a < 15 {
		fmt.Printf("a %v'ten küçüktür.", a)
	} else if a > 15 {
		fmt.Printf("a %v'ten büyüktür.", a)
	} else {
		fmt.Printf("a %v'e eşittir.", a)
	}
}
```

Çıktı:
```
a 15'e eşittir.
```

## switch...case statement 

Switch statement, yürütülebilecek birçok kod bloğundan birini seçmek için kullanılır.

```
    switch <expression> {
    case condition:
        /* Code */
    default:
        /* Code */     
    }
```

```
func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a = 2")
	case 3:
		fmt.Println("a = 3")
	case 4:
		fmt.Println("a = 4")
	default:
		fmt.Println("a herhangi bir değer olabilir.")
	}
}
```

Çıktı:
```
a = 3
```