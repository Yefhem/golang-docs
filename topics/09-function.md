# Fonksiyonlar (Functions)

Fonksiyonlar, kod bloğunu yeniden kullanma yeteneği veren, aşırı bellek kullanımından tasarruf sağlayan, zaman kazandıran ve en önemlisi kodun daha iyi okunabilirliğini sağlayan bir programdaki kodlar veya ifadeler bloğudur.

Günün sonunda fonksiyon, belirli bir işlevi yerine getiren ve sonucu döndüren ifadeler topluluğudur. Bir fonksiyon herhangibir şey döndürmeden de belirli bir işlevi yerine getirebilir.

## Function Declaration

![and](https://media.geeksforgeeks.org/wp-content/uploads/Untitled-Diagram-37.jpg)

- **keyword** > Go dilinde bir fonksiyon oluşturmak için kullanılan anahtar kelimedir(func).
- **function_name** > Fonksiyonun ismidir. Fonksiyon isimlendirmede dikkat edilmesi gereken bazı hususlar vardır. 1. İlk karakter harf olmalıdır. 2. Camel case kullanılması önerilir. 3. Paket dışı kullanımlarda ilk harf büyük olmalıdır.
- **parameter_list** > Fonksiyonun parametrelerinin adını ve türünü içerir.
- **return_type** > Fonksiyonumuz içerisinde return işlevi yapıyorsak döndürdüğümüz nesnenin tipini belirtmek zorundayız.
- **function_body** > Fonksiyonumuzda kodlarımızı yazdığımız bölümdür.

```
func Area(len, width int) int {
	return len * width
}

func main() {
	fmt.Printf("Dikdörtgenin Alanı: %v", Area(5, 4))
}
```

Çıktı:
```
Dikdörtgenin Alanı: 20
```

## Variadic Function (Değişken Fonksiyonlar)

Değişken fonksiyonlar, değişken sayıda argümanla çağırılan fonksiyonlardır. Sıfır veya daha fazla argüman alabilirler. **fmt.Printf** variadic fonksiyona bir örnektir. 

Declaration işlemi sırasında son parametrenin türünden önce üç nokta(...) kullanılmalıdır.

```
function function_name(para1, para2...type)type{
    // code...
}
```

Fonksiyon içerisindeki ...type bir slice gibi davranır. Yani (b ...int) = []int ile eşdeğer gibi düşünebiliriz.

```
func a(element ...string) string {
	return strings.Join(element, " + ")
}

func main() {
	// zero argument
	fmt.Println(a())

	// multiple arguments
	fmt.Println(a("apple", "banana"))
	fmt.Println(a("A", "P", "P", "L", "E"))

	element := []string{"orange", "cherry", "strawberry"}

	fmt.Println(a(element...))
}
```

Çıktı:
```

apple + banana
A + P + P + L + E
orange + cherry + strawberry
```

## Anonymous Functions (Anonim Fonksiyonlar)

İsimsizlerdir. Yazıldıkları yerde direkt olarak çalışırlar.

```
func main() {
	func() {
		fmt.Println("Welcome!")
	}()
}
```

Çıktı:
```
Welcome!
```
 
Bir değişkene anonim bir fonksiyon atanabilir.

```
func main() {
	v := func() {
		fmt.Println("Apple!")
	}
	v()
	fmt.Printf("type: %T", v)
}
```

Çıktı:
```
Apple!
type: func()
```

Anonim fonksiyona bağımsız argümanlar iletebiliriz.

```
func main() {
	func(s string, y int) {
		fmt.Printf("Apple %v %v", s, y)
	}("Computer", 1976)
}
```

Çıktı:
```
Apple Computer 1976
```

Anonim bir fonksiyonu değişken olarak başka bir fonksiyona iletebiliriz.

```
func y(i func(a, b string) string) {
	fmt.Println(i("hello ", "world"))
}

func main() {
	v := func(a, b string) string {
		return a + b
	}
	y(v)
}
```

Çıktı:
```
hello world
```

## Main ve Init Fonksiyonları

**Main Function:**

Go programlama dilinde main package, çalıştırılabilir programlarla birlikte kullanılan özel bir pakettir ve içerisinde main fonksiyonu içerir. Main fonksiyonu özel bir fonksiyon olmakla birlikte yürütülebilir programların giriş noktasıdır. Herhangi bir argüman almaz ve herhangi bir argüman döndürmez. Go otomatik olarak main fonksiyonunu çağırır. Sonuç olarak her yürütülebilir program tek bir main package ve tek bir main fonksiyonu içermelidir.

```
func main() {
	// code
}
```

**Init Function:**

Main fonksiyonuna benzerdir. Herhangi bir argüman almaz veya herhangi bir argüman döndürmez. Bu fonksiyon her pakette bulunur ve paket başlatıldığında bu fonksiyon başlatılır. Bu işlev implicitly(üstü kapalı, örtük) olarak declare edilir. Yani bu fonksiyon herhangi bir yerden başvurulamaz.

Aynı programda birden çok init fonksiyonu oluşturabiliriz. Bu oluşturulan init fonksiyonları oluşturuldukları sıraya göre yürütülürler.

Init fonksiyonu her zaman main fonksiyonundan önce yürütülür. Bu nedenle main fonksiyonuna bağlı değildir. Init fonksiyonunun temel amacı global bağlamda başlatılmayan global değişkenleri başlatmaktır.

```
func init() {
	fmt.Println("init c")
}

func init() {
	fmt.Println("init a")
}

func init() {
	fmt.Println("init b")
}

func main() {
	fmt.Println("main")
}
```

Çıktı:
```
init c
init a
init b
main 
```
