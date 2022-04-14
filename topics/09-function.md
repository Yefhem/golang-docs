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






