# Go'da map[string]interface{} Kullanımı 

Türünü bilmeden bir değere atıfta bulunmanın en iyi yollarından biri interface kullanmaktır. 

Buradaki interface değeri bir değişken, bir dizi, bir tam sayı, os.File için bir pointer gibi aklımıza gelebilecek çoğu şeyi tutabilir.

**Bir map[string]interface{} örneği;**

```
func main() {
	foods := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	fmt.Printf("Value: %v\nType: %T", foods, foods)
}
```

Çıktı:
```
Value: map[bacon:delicious eggs:{chicken 1.75} steak:true]
Type: map[string]interface {}
```

Kendisine verilen değeri yazıdıran bir fonksiyon yazmamız gerektiğini varsayalım ancak değerin ne tür olacağını bilmiyoruz. Bu gibi durumlarda interface kullanırız.

```
func main() {
	writeAnything("apple")
	writeAnything(true)
	writeAnything(1.54)
}

func writeAnything(v interface{}) {
	fmt.Println(v)
}
```

Çıktı:
```
apple
true
1.54
```

Aynı şekilde fmt.Println methodu da benzer şekilde tanımlanır.

```
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

Aynı şekilde her biri bir string ile tanımlanan herhangi bir türden herhangi bir şey koleksiyonu istiyorsak bu rastgele verileri düzenlemenin uygun bir yolu olarak **map[string]interface{}** kullanabiliriz.

Aslında bunu yaparak JSON nesnelerin şemasını tanımlamış oluyoruz.

Aşağıdaki ham JSON verilerine bakıcak olursak;

```
{
   "name":"John",
   "age":29,
   "hobbies":[
      "martial arts",
      "breakfast foods",
      "piano"
   ]
}
```

Bu ifadeyi eğer bir go yapısına çevirmek istersek, şöyle bir yapı oluşturmamız gerekecek...

```
type Person struct {
    Name    string
    Age     int
    Hobbies []string
}
```

Bu durum için tamam fakat nesnenin şemasını önceden bilmemiz gerekiyor.

Ya biri bize keyfi JSON verisi verirse ve bizim bunu bir Go değerine ayırmamız gerekirse?

## JSON data'yı *map[string]interface{}* Decode etme

JSON verilere sahip olduğumuzu düşünelim ve bu verileri data isimli değişkende saklağımızı varsayalım. Bu veriyi bir go değişkenine nasıl ayarlayabiliriz. Bunun için ne tür bir değişkene ihtiyacımız var?

Tam bu noktada **map[string]interface{}** istediğimiz bir yapıdır. 

```
var p map[string]interface{}
err = json.Unmarshal(data, &p)
```

Özetle dışarıdan gelen verilerle uğraşmamız gerektiğinde map[string]interface{} yapısı çok kullanışlıdır. Örneğin şeması bilinmeyen rastgele JSON verileri.


