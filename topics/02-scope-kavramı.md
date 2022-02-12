# Go'da Scope Kavramı

Scope Türkçe karşılığı olarak **kapsam** veya **alan** diyebiliriz.

Bir değişkenin kapsamı yani scope'u o değişkenin görülebilir olduğu komutların alanıdır desek yanlış olmaz.

Go'da scope kavramını **blog scope** ve **package scope** olarak ikiye bölersek; blog scope, belirli bir blog içerisinde tanımlı ve sadece o blog içerisinde erişilebilir değişkenleri bize çağrıştıracaktır Blog'dan kastımız aslında süslü parentezler **{}** içerisinde kalan kısımlar diyebiliriz. Package scop ise tüm paket içerisinde erişebildiğimiz değişkenleri bize çağrıştıracaktır. Bu değişkenler global veriable olarak da bilinir. Bu değişkenlere istediğimiz alandan ulaşabiliriz(package içerisinde). 

Şöyle de ufak bir not geçmekte fayda var. Package değişkenleri oluştururken sort hand declaration yapamıyoruz. 

Şimdi örneklerle yukarıdaki anlattıklarımızı pekiştirelim.

```
var a string = "a"
var b string = "b"

func main() {
	var b string = "b2"

	fmt.Println(a) // a
	fmt.Println(b) // b2
}
```

Çıktı:
```
a
b2
```

görüldüğü üzere blog scope içerisinde tanımlanan b değişkeni global değişken olarak tanımlanan b değişkenine göre daha baskın geliyor.

```
var a string = "a"
var b string = "b"

func main() {
	var b string = "b2"

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("--")

	myLittleFunc()
}

func myLittleFunc() {
	fmt.Println(a)
	fmt.Println(b)
}
```

Çıktı:
```
a
b2
--
a
b
```

Bu durumda ise görüldüğü üzere myLittleFunc fonksiyonu sadece global değişkenleri görebiliyor. main fonksiyonu içerisindeki b değişkeni myLittleFunc fonksiyonu için bir anlam ifade etmiyor. 

```
var b string = "b"

func main() {
	var b string = "b2"

	if true {
		var b string = "b3"
		fmt.Println(b)
		if true {
			var b string = "b4"
			fmt.Println(b)
		}
	}

	fmt.Println(b)

	myLittleFunc()
}

func myLittleFunc() {
	fmt.Println(b)
}
```

Çıktı:
```
b3
b4
b2
b 
```

görüldüğü üzere blog scope her zaman daha baskın geliyor. 

**Not:** Aynı isimli değişkenler oluşturmak her zaman için bir risk oluşturuyor fakat böyle bir kullanım yapma şansımız var. Değişkenlerin kapsamını belirlerken şu unutulmamalıdır. Package değişkenleri yani package scope içerisinde tanımlı olanlar, programın tüm çalışma süresi boyunca hafızada yer kaplarlar. Bu dikkat edilmesi gereken bir husustur. 