# Sabitler (Constants)

Constant kavramına geçmeden önce "run-time" ve "compiler-time" kavramlarına değinmekte fayda var.

## Çalışma-zamanı (run-time) Nedir?

run-time veya execution-time olarak bilinen bu kavram kısaca, bir programın veya kod bloğunun calıştırıldığı andan sonlandığı zamana kadar geçen zaman aralığını ifade eder.

Bu kavram derleme-zamanı(compiler-time), bağlantı-zamanı(connection-time), yükleme-zamanı gibi kavramlara nazaran daha sık görülmektedir. Örneğin çalışma-zamanı hatası (run-time error) programın çalışması sırasında rastlanır.

Error handling (hata yönetimi) veya hata yakalama, çalışma zamanı hatalarını yakalamak için tasarlanan bir dil özelliğidir. Tahmin edilen hataları, beklenmedik durumlar vb.. için yapısal bir yol sağlar.

## Derleme-zamanı (compiler-time) Nedir?

Bilgisayar bilindiği üzere sadece makine dilini anlayabilir.

> **Makine Dili Nedir?** 
Mikroişlemci ya da mikrodenetleyici benzeri komut işleme geleneğine sahip entegrelerin işleyebileceği komutlardan ve buna uygun söz diziminden oluşan dile verilen addır. Makine dili donanıma en yakın, en alt seviye programa dilidir. Bu dil sadece 0 ve 1 'lerden oluşur.

Herhangi bir kod bloğunu çalıştırdığımızda ilk olarak makine diline çevrilir. Tam bu noktada, kaynak kodun makine koduna dönüştürülmesi derleme-zamanı(compiler-time) olarak bilinir. Derleme-zamanında oluşturulan yürütülebilir dosyaların çalıştırılması ise yukarıda bahsettiğimiz run-time ile ilgilidir.

Const (sabit) kavramına gelecek olursak, uygulama akışı boyunca değişmeyen ifadelerdir. 

### Nasıl Declare Edilir?

Değişkenler gibi declare edilir fakat **const** anahtar sözcüğü ön ek olarak kullanılmalıdır. Short declaration **:=** kullanılamaz.

```
const PI float64 = 3.14

func main() {
	const HW string = "World"
	fmt.Println("Hello", HW)

	fmt.Println("Pi Number:", PI)

	const False bool = false 
	fmt.Println(False)

	const x, y, z int = 10, 20 ,30
	fmt.Println(x, y, z)
}

```

Çıktı:
```
Hello World
Pi Number: 3.14
false
10 20 30 
```

## Typeless/Untyped Constants (Tipsiz Sabitler) Nedir?

Sabitlerde tür belirtilebildiği gibi (bakınız yukarıdaki örnekler) tür belirtilmeden de kullanılabilir. 

```
func main() {
	const a = 50
	const b = 50.5
	const c = "apple"
	const d = true

	fmt.Printf("%v - %T\n", a, a)
	fmt.Printf("%v - %T\n", b, b)
	fmt.Printf("%v - %T\n", c, c)
	fmt.Printf("%v - %T\n", d, d)
}
```

Çıktı:
```
50 - int
50.5 - float64
apple - string
true - bool  
```

## Dikkat Edilmesi Gerekenler

- Sabitler **compiler-time** yani derleme-zamanına, değişkenler ise **run-time** yani çalışma-zamanına aittir.
- Sabitler oluşturulduğu andan itibaren değeri verilmelidir.
- Sabitlere değişken **atamanmaz** (değişkenler run-time'de çalışır). Fakat bir sabitten bir değişkene değer ataması yapılabilir.