# Değişken Kavramı ve Temel Veri Tipleri
## Değişken(variable) Nedir?
Bilgisayar ve matematik biliminde istediğimiz bir ifade veya niceliği tutmaya/saklamaya/depolamaya olanak sağlayan yapılara değişken denir.

## Tek bir değişken bildirmek (Declaring a single variable)

Go'da tek bir değişken bildirimi yaparken kullanacağımız sözdizimi şu şekildedir:

`var animal string (keyword - variable name - type)`

declaration işlemi ile hafızada bir yer ayrılır. Bu hafızaya herhangi bir assign(atama) işlemi yapılmazsa, değişken tipinin zero value'su yerleştirilir.

> Başlangıç değeri olmadan declare edilen değişkenlere zero value atanır. 
>
> Bu değer sayısal tipler için **0**
>
> Boolean tip için **false**
>
> Stringler için ise **""** boş bir string'dir.

`animal = "snake"`

burada ise declare ettiğimiz değişkene assign işlemi yapıyoruz ve **snake** değerini atıyoruz. 

```
var animal string
animal = "snake"
```

Yukarıdaki işlemlerin tamamına ise **initialization** işlemi diyebiliriz. Yani bir değişken oluşturup, değer atayıp, başlatma işlemini kapsar.

Bu işlem tek bir satırda da yapılabilir:

`var animal string = "snake"`

## Çoklu değişken bildirmek (Multiple variable declaration)

Go'da tek bir ifade ile birden fazla değişken bildirimi yapmak mümkündür.

`var a, b, c int = 5, 10, 15`

veya 

```
var name, surname string 
name = "Joe"
surname = "MacMillan"
```

## Tür çıkarımı (Type inference)

Go'da bir değişkenin başlangıç değeri varsa, Go bu ilk değere bakarak değişken türünü otomatik olarak çıkarabilir.

`var age = 10`

Burada Go otomatik olarak age değişkeninin veri tipini **int** olarak belirler. Başka bir örneğe bakacak olursak:

`var name, age, isMarried, weight, height = "Han", 33, false, 70.5, 180`

**Not:** Go Statically Typed bir dildir. Değişkenlerin tipleri değiştirilemez. Statically Typed dillerin avantajı Dynamic Typed dillerde run-time sırasında sürekli tip kontrolleri programı bir nebzede olsa yavaşlatabilir. Diğer taraftan Statically Typed dillerde derleme esnasında tip kontrolleri yapıldığı için programın çalışması sırasında, ekstra tip kontrolü yapılmaz. Bu da ekstra bir verimlilik ve hız kazandırır. Statically Typed ayrıca çok daha güvenlidir.

## Kısa değişken bildirimi (Short variable declarations)

Go bize değişkenleri declare etmek için kısa bir yol sunar. Bu yok short declarations olarak bilinir ve **:=** operatörü ile kullanılır.

`name := "steve"`

Burada hem name değişkenini declare ediyoruz hem de "steve" string'ini assign ediyoruz. Go tür çıkarımı yaparak name değişkeninin string olduğunu biliyor.

**Not:** Function body dışında yani global değişkenler tanımlanırken short declaration kullanılamaz.

## Veri tipleri (Data Types)

Temel veri tiplerini çok detaya girmeden genel olarak üstünden geçecek olursak;

### İnteger 

Bir tamsayıyı ifade ederken integer kullanılır. Signed(işaretli) ve unsigned(işaretsiz) olarak iki bölüme ayrılır. Signed ifadeler eksi değerleri de kapsar.

| Numeric Types | Bit | s/u |
| ----------- | ----------- | ----------- |
| **uint** | * | unsigned |
| uint8 | 8-bit | unsigned |
| uint16 | 16-bit | unsigned |
| uint32 | 32-bit | unsigned |
| uint64 | 64-bit | unsigned |
| **uintptr** |  any pointer |  |
| * | * | * |
| **int** | *| signed |
| int8 | 8-bit | signed |
| int16 | 16-bit | signed |
| int32 | 32-bit  | signed |
| int64 | 64-bit | signed |

**Not:**
*int*, *uint* ve *uintptr* türleri genellikle 32 bit sistemlerde 32 bit ve 64 bit sistemlerde 64 bit'dir.

**Example**

```
var (
	MaxInt uint64     = 1<<64 - 1
	Number int8		  = 19
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Number, Number)
}
```

Çıktı:

```
Type: uint64 Value: 18446744073709551615
Type: int8 Value: 19
```

### Float 

Noktalı değerleri ifade etmek için float kullanılır. Go'da 32bit ve 64bit'i temsil eden iki farklı float biçimi vardır.

| Numeric Types | Bit |
| ------ | ----- |
| float32 | 32-bit | 
| float64 | 64-bit | 

**Example**

```
var (
	Number1 float32 = 17.31
	Number2	float64 = 99.9999999999
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", Number1, Number1)
	fmt.Printf("Type: %T Value: %v\n", Number2, Number2)
}
```

Çıktı:

```
Type: float32 Value: 17.31
Type: float64 Value: 99.9999999999
```

### Complex Number

Karmaşık sayıları ifade etmek için kullanılır. Matematikte karmaşık sayılar iki bölümden oluşur. Birinci bölüm gerçek sayı, ikinci bölüm ise hayali sayıyı temsil eder. Go'da 64bit ve 128bit'i temsil eden iki farklı Complex Number biçimi vardır.

| Numeric Types | Bit | * |
| ----- | ----- | ----- |
| complex64 | 64-bit | ----- |
| complex128 | 128-bit | ----- |

**Example**

```
var (
	comp complex128 = cmplx.Sqrt(-3 + 45i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", comp, comp)
}
```

Çıktı:

```
Type: complex128 Value: (2+3i)
```

### Boolean

True(1) veya False(0) temsil eder.

**Example**

```
var (
	isMarried bool = true
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", isMarried, isMarried)
}
```

Çıktı:

```
Type: bool Value: true
```

### Byte 

8 bitlik işaretsiz sayıya (uint8) denktir.

```
var (
	num byte
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", num, num)
}
```

Çıktı: 

```
Type: uint8 Value: 0
```

### String 

String'ler bir karakter dizisini temsil eder.

**Example**

```
var (
	name string = "We Good!"
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", name, name)
}  
```

Çıktı: 

```
Type: string Value: We Good!
```

**Not2:** Veri tiplerinin zero value'ları yukarıda belirtilmiştir.