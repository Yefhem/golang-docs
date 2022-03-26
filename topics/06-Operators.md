# Operatörler (Operators)

Operatör, derleyiciye belirli eylemleri gerçekleştirmesini söyleyen bir semboldür.

## Aritmetik Operatörler (Arithmetic Operators)

Aritmetik operatörler, toplama, çıkarma, çarpma, bölme vs. gibi temel matematik işlemlerini gerçekleştirmek için kullanılır.

| Aritmetik Operatörler | Açıklama | Örnek Gösterim | Sonuç
| ----- | ----- | ----- | -----
| + | toplama | x + y | iki değeri toplar
| - | çıkarma | x - y | bir değeri diğerinden çıkarır
| * | çarpma | x * y | iki değeri çarpar
| / | bölme | x / y | bir değeri diğerine böler
| % | mod alma | x % y | iki değerin bölümünden kalanı verir
| ++ | arttırma | x++ | değeri bir arttırır
| -- | azaltma | x-- | değeri bir azaltır
 

```
func main() {
	var x, y int = 15, 10

	fmt.Printf("x + y = %d\n", x + y)
	fmt.Printf("x - y = %d\n", x - y)
	fmt.Printf("x * y = %d\n", x * y)
	fmt.Printf("x / y = %d\n", x / y)
	fmt.Printf("x %% y = %d\n", x % y)

	x++
	y--
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)
}
```

Çıktı:
```
x + y = 25
x - y = 5  
x * y = 150
x / y = 1  
x % y = 5  
x = 16     
y = 9
```

## Atama Operatörleri (Assignment Operators)

Atama operatörleri, bir değişkene değer atamak için kullanılan operatörlerdir.

| Atama Operatörleri | Açıklama | Örnek Gösterim |
| ----- | ----- | ----- | 
| = | atama | x = y | 
| += | ekleme ve atama | x = x + y | 
| -= | çıkartma ve atama | x = x - y | 
| *= | çarpma ve atama | x = x * y | 
| /= | bölme ve atama | x = x / y | 
| %= | bölme ve modu atama | x = x % y | 

```
func main() {
	var x, y int = 15, 10

	x = y 
	fmt.Printf("1. sonuç = %d\n", x)

	x = 15
	x += y
	fmt.Printf("2. sonuç = %d\n", x)

	x = 15
	x -= y
	fmt.Printf("3. sonuç = %d\n", x)

	x = 15
	x *= y
	fmt.Printf("4. sonuç = %d\n", x)

	x = 15
	x /= y
	fmt.Printf("5. sonuç = %d\n", x)
}
```

Çıktı:
```
1. sonuç = 10
2. sonuç = 25 
3. sonuç = 5  
4. sonuç = 150
5. sonuç = 1  
```

## Karşılaştırma Operatörleri (Comparison Operators)

Karşılaştırma operatörleri, iki değeri karşılaştırmak için kullanılan operatörlerdir.

| Karşılaştırma Operatörleri | Açıklama | Örnek Gösterim |
| ----- | ----- | ----- | 
| == | eşit | x == y | 
| != | eşit değildir | x != y | 
| < | küçüktür | x < y | 
| <= | küçük eşittir | x <= y | 
| > | büyüktür | x > y | 
| >= | büyük eşittir | x >= y | 

```
func main() {
	var x, y int = 15, 10

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)
}
```

Çıktı:
```
false
true 
false
false
true 
true 
```

## Mantıksal Operatörler (Logical  Operators)

Mantıksal operatörler, değişkenler veya değerler arasında True ve False karşılaştırması yapması için kullanılan operatörlerdir.

| Mantıksal Operatörler | Açıklama | Örnek Gösterim |
| ----- | ----- | ----- | 
| && | mantıksal ve  | x < y && y > z | 
| &#124;&#124; | mantıksal veya | x < y &#124;&#124; y > z | 
| ! | mantıksal değildir | !(x < y && y > z) | 

```
func main() {
	var x, y, z int = 15, 10, 5

	fmt.Println(x < y && y > z)
	fmt.Println(x < y || y > z)
	fmt.Println(!(x < y && y > z))
}
```

Çıktı:
```
false
true
true
```

## Bitsel Operatörler (Bitwise Operators)

Bitsel operatörler, binary düzeyde sayıları karşılaştırırken kullanılan operatörlerdir.

| Bitsel Operatörler | Açıklama | Örnek Gösterim |
| ----- | ----- | ----- | 
| & | ve  | **x & y =** Her iki bit de 1 ise = 1 aksi durumlarda = 0 | 
| &#124; | veya | **x &#124; y =** Her iki bit'ten herhangi biri 1 ise = 1 aksi durumlarda = 0| 
| ^ | xor | **x ^ y =** İki bit'ten sadece biri 1 ise = 1 aksi durumlarda = 0 | 
| << | mantıksal değildir | **x << 1 =** Sola doğru bitleri kaydırır | 
| >> | mantıksal değildir | **x >> 1 =** Sağa doğru bitleri kaydırır | 


```
func main() {
	var x uint = 50 // 00110010
	var y uint = 30 // 00011110
	var z uint

	z = x & y
	fmt.Printf("x & y = %d binary = %b\n", z, z) // 00010010

	z = x | y
	fmt.Printf("x | y = %d binary = %b\n", z, z) // 00111110

	z = x ^ y
	fmt.Printf("x ^ y = %d binary = %b\n", z, z) // 00101100

	z = x << 1
	fmt.Printf("x << y = %d binary = %b\n", z, z) // 01100100

	z = x >> 1
	fmt.Printf("x >> y = %d binary = %b\n", z, z) // 00011001
}
```

Çıktı:
```
x & y = 18 binary = 00010010 
x | y = 62 binary = 00111110
x ^ y = 44 binary = 00101100   
x << y = 100 binary = 01100100
x >> y = 25 binary = 00011001
```