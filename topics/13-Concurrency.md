# Eş Zamanlılık (Concurrency)

Pek çok durumda, eş zamanlı olarak gelen birden çok işlem talebi söz konusu olabilir. Aslında eş zamnlılıktan kasıt cpu'da bir proccess'in tam olarak bitmeden diğer proccess ile ilgili işlerin de cpu tarafından yapılmasıdır. 

Bir program içerisinde (n) sayıda görevin aynı anda çalışması olarak akla gelmelidir (Burada fonksiyonlar sırayla değil aynı anda çalıştırıldığı anlaşılmalıdır).

Eş zamanlılığa örnek olarak web sunucuları verilebilir. Sunuculara gelen ziyaretçiler eş zamanlı olarak web sitelerine ulaşabilmektedir.

Parallelism ile çokca karıştırılan bir durumdur. Bir kodun concurrently (aynı anda) çalışıyor olduğunu söylemek o kodun paralel olarak çalıştığını göstermiyor.

![con_and_par](/images/Untitled.jpg)

> Parallelism ise aynı anda bir fazla isteğimiz mevcutsa ve bu işlemleri aynı anda yapmak istiyorsak ("aynı andan yapmaktan" daha çok minimum talebi minimum bekletme ile gerçekleştirmek.) bu noktada parallelism kavramından bahsedebiliriz. Parallellik, gelen işlem taleplerini birden fazla işlem ünitesi ile aynı anda çalıştırılması prensibine odaklanır.

Bu iki konunun farkını en iyi şekilde anlamak için Joe Armstrong'un kahve dükkanı örneğine bakılabilir.

![con_and_par](/images/con_and_par.jpg)

## Go Dilinde Eş Zamanlılık

Go dili concurrency sayesinde eş zamanlı fonksiyon çalıştırmayı destekler. 
Go dilinde eş zamanlı görevleri yerine getirebilmek için go routine ve channels kullanılmaktadır. 

### Go Routines 

Go routine denilince birbirinden "bağımsız" ve "eş zamanlı" olarak çalışan ufak iş parçacıkları aklımıza getirebiliriz. Go dilinde bunu gerçekleştirmek için tek yapmamız gereken bağımsız çalışmasını istediğimiz fonksiyonun başına **go** keyword'unu eklemek.

```
func main() {
	fruits := []string{"banana", "apple", "orange", "lemon", "apple", "watermelon", "apple"}
	go findA(fruits)
	go findB(fruits)
	fmt.Scanln()

}

func findA(arr []string) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == "apple" {
			fmt.Println("Find A : " + arr[i])
		}
		time.Sleep(time.Second)
	}
}

func findB(arr []string) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == "apple" {
			fmt.Println("Find B : " + arr[i])
		}
		time.Sleep(time.Second)
	}
}
```

Çıktı:
```
Find A : apple
Find B : apple
Find B : apple
Find A : apple
Find A : apple
Find B : apple
```

Görüldüğü üzere fonksiyonlar sıra ile çalıştırılmadı yani iki fonksiyon birbirini beklemeden ve eş zamanlı olarak çalışmıştır.

Peki go routines kullanılmasaydı ne olacaktı.. İlk olarak **findA** fonksiyonu ardından **findB** fonksiyonu çalışacaktı. Çıktımız ise;

Çıktı:
```
Find A : apple
Find A : apple
Find A : apple
Find B : apple
Find B : apple
Find B : apple
```

**Anonymous Go Routine Nedir?**

Dışarıdan fonksiyon çağırmadan, pratik bir biçimde go routine kullanmamıza olanak tanır. Oldukça yaygın bir kullanımdır.

```
func main() {
	fruits := []string{"banana", "apple", "orange", "lemon", "apple", "watermelon", "apple"}

	go func(arr []string) {
		for i := 0; i < len(arr); i++ {
			fmt.Printf("anon go routine : %v\n", arr[i])
			time.Sleep(time.Second)
		}
	}(fruits)

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
		time.Sleep(time.Second)
	}
	fmt.Scanln()
}
```

Çıktı:
```
banana
anon go routine : banana
anon go routine : apple
apple
orange
anon go routine : orange
anon go routine : lemon
lemon
apple
anon go routine : apple
anon go routine : watermelon
watermelon
anon go routine : apple
apple
```

Burada anonim bir go routine içerisindeki for loop ile aynı işlevi yapan for loop eş zamanlı olarak çalışmıştır.

### Channels (Kanallar)

Go routines birbirinden bağımsız ve iletişim kurmadan çalışırlar. Sessizce çalışır işleri bittiğinde kullandıkları kaynağı iade ederler. Peki channel'larda durum ne?

Channels sayesinde go routine'ler birbiri ile iletişim kurabilirler ve birbirlerine sinyaller göndererek senkronize çalışabilirler. Bir go routine herhangi bir channel'a veri gönderebilirveya channel'dan veri alabilir. Bunu yaparken de verinin akış yönünü gösteren ok (<-) kullanılır.

```
firstChannel <- "we good!" // veri gönderme
variable := <- firstChannel 
```

Channel'lar kullanılmaya başlanmadan önce tanımlanmalıdır.

`firstChannel := make(chan string)`

Go routine'lerin haberleşmesini bir örnek üzerinden inceleyelim..

```
func main() {
	colors := []string{"red", "blue", "orange", "pink", "purple"}

	firstChannel := make(chan string)

	go func(arr []string) {
		for _, color := range arr {
			firstChannel <- color // Kanala veriyi gönderiyoruz.
			time.Sleep(time.Second)
		}
	}(colors)

	go func() {
		for i := 0; i < 5; i++ {
			data := <-firstChannel // Kanaldan alınan veri data değişkenine atanıyor.
			fmt.Println(data)
		}
	}()
	<-time.After(time.Second * 7) 
    // Main go routine tamamlandığında diğer tüm go routine'lerin           sonlandırıldığını biliyoruz. Uygulamamızın sağlıklı çalışabilmesi için main go routine'nin diğer go routine'leri beklemesi için ilave önlem almış olacağız.
}
```

Çıktı:
```
red
blue
orange
pink
purple
```

Buranın çalışma mantığını anlatacak olursak, kanallar arasında gönderme ve alma işlemleri yaparken birbirlerini bloke ederler. 
Kanala bir veri gönderdiğimizde mevcut kontrol mekanizması kanala veri girişinin sağlandığı satırda bloklanır, ta ki başka go routine o kanaldaki veriyi okuyana kadar.
Aynı işlem kanaldaki veri okunduktan sonra  okuma işleminin yapıldığı satırdaki kontrol bloklanır, ta ki başka bir go routine tarafından o kanala veri girişi yapılana kadar.
Yani `firstChannel <- color` burada kanala veri girdiği anda kontrol mekanizması (burada for) durur ve `data := <-firstChannel` bu satırdaki okuma işlemi devreye girer. 
Veri okunduktan sonra da buradaki kontrol durur ve yeniden veri girişi yapılması beklenilir.

**Main Go Routine Nedir?**

Her go prejesinde main fonksiyonu ve main go routine mevcuttur. Main fonksiyonu başlı başına bir go routine'dir. 
Hatta en büyük go routine desek yanlış olmaz. Main fonksiyonu tamamlandığı anda diğer bütün go routine'ler de sonlandırılmış olur. 

`<-time.After(time.Second * 7)` 

main go routine'imizi bekletmek için bir channel oluşturup 7 saniye sonra veri gönderilmesini sağladık bir diğer şekilde 7 saniye beklettik.


