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

**Done Channel Nedir?**

Done channel(tamamlandı kanalı) program bitene kadar main go routine'ni bekletmenin uygun yollarından biridir.
Projenin başında bir done channel tanımlanır ve işlemlerimiz tamamlandığında bu kanala veri göndeririz. 
Bu channel main fonksiyonu okuyacak şekilde yerleştirilirse biz tamamlandı komutu veresiye kadar main fonksiyonu/go routine'i çalışmaya devam edecektir.


```
func main() {
	doneChan := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hey ", i)
			time.Sleep(time.Second)
		}
		doneChan <- "done.."
	}()
	<-doneChan
}
```

Çıktı:
```
hey  0
hey  1
hey  2
hey  3
hey  4
hey  5
hey  6
hey  7
hey  8
hey  9
```

Burada go routine içerisindeki for döngüsü tamamlandıktan sonra main go routine'e bir veri göndererek işimizin bittiğini belirtiyoruz.

**Sender, Receiver ve Direction**

Bir kanala veri gönderip alma işlemini kolaylaştırmak için bu tip fonksiyonlar kullanabiliriz. Bu oluşturduğumuz fonksiyonlarda tek bir işi yapmaya zorlamış oluruz yani sadece alıcı veya verici olması garanti edilir. Burada kanallar fonksiyon parametresi olarak kullanılır.

Bununla birlikte aşağıda göreceğiniz direction fonksiyonu gibi tek işi bir kanaldaki veriyi alıp başka bir kanala aktaran fonksiyonlar da yazabiliriz.

```
func main() {
	firstChan := make(chan string, 1)

	sender(firstChan, "data")

	message := receiver(firstChan)

	fmt.Println(message)
}

func sender(channel chan<- string, message string) {
	channel <- message
}

func receiver(channel <-chan string) string {
	message := <-channel 
	return message
}
```

Çıktı:
```
data
```

> Burada sender fonksiyonunun yaptığı tek iş mesajı kanala göndermektir.

> Receiver fonksiyonunun yaptığı tek iş kanaldaki veriyi almaktır. 

```
func main() {
	firstChan := make(chan string, 1)
	secondChan := make(chan string, 1)

	sender(firstChan, "hello world")

	direction(firstChan, secondChan)

	fmt.Println(<-secondChan)
}

func sender(channel chan<- string, message string) {
	channel <- message 
}

func direction(receiver <-chan string, sender chan<- string) {
	message := <-receiver 
	sender <- message  
}
```

Çıktı:
```
hello world
```

> Direction fonksiyonunun yaptığı tek iş ise bir kanaldaki veriyi alıp başka bir kanala aktarmaktır.

> Biraz daha detaya inersek firstChan kanalına sender fonksiyonu ile "hello world" stringi aktarıldı. Direction fonksiyonunda firstChan kanalındaki veri message değişkenine atandı. 
Message değişkeni içerisindeki veriyi SecondChan kanalına gönderdik ve kanaldaki veriyi yazdırdık.

**Not:** Eğer direkt olarak `fmt.Println(secondChan)` bu şekilde yazdıracak olsaydık secondChan kanalının adresini elde edecektik.

**Buffered ve Unbuffered Channel**

*Buffered: Kanallarda birden fazla veri tutabilir. Kanal oluştururken buffered veya unbuffered olduğu belirtilir*

> FIFO(first in first out) yaklaşımı burada da geçerlidir. Kanala ilk giren veri okuma tarafında ilk çıkan veridir.

> Kanala veri girişi ve çıkışı esnasında bloklama işlemi gerçekleşir. Unbuffered ile fark yoktur.

> Kanala veri girişi esnasında kanalın tamamı dolu ise yalnızca yeni veri girişi bloklanır. Ve kanal boşsa yalnızca veri çıkışı işlemi bloklanır.

Örnek Buffered Channel Tanımlanması;

`myChannel := make(chan string, 50)`

*Unbuffered: Bu tür kanallar sadece bir adet veri muhafaza edebilir.*

`myChannel := make(chan string, 1)`

```
func main() {
	myBufferedChan := make(chan string, 2)
	go func() {
		myBufferedChan <- "first"
		fmt.Println("first sent")
		myBufferedChan <- "second"
		fmt.Println("second sent")
	}()
	<-time.After(time.Second * 2)
	go func() {
		firstRead := <-myBufferedChan
		fmt.Println("...")
		fmt.Println(firstRead)
		secondRead := <-myBufferedChan
		fmt.Println(secondRead)
	}()
	<-time.After(time.Second * 2)
}
```

Çıktı:
```
first sent
second sent
...
first 
second
```

**Deadlock Problemi**

En yaygın hatalardan biri olan deadlock bir channel'a gönderen kadar okuyucu atanmazsa yani kapasiteden fazla veri gönderilirse programımız deadlock problemi ile karşılaşır.

```
func main() {
	myBufferedChan := make(chan string, 2)

	myBufferedChan <- "1"
	myBufferedChan <- "2"
	myBufferedChan <- "3"

	one := <-myBufferedChan
	two := <-myBufferedChan
	three := <-myBufferedChan

	fmt.Println(one, two, three)

	<-time.After(time.Second * 2)
}
```

Çıktı: 
```
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        C:/*******/******/go/src/github.com/Yefhem/golang-docs/main.go:13 +0x6a
exit status 2
```

Aslında bu problem kanalın kapasitesi ile ilgili bir mesele. Bu gibi durumlarda len() ve cap() fonksiyonları kullanılabilir. 

> Len() : kanal içi mevcut veri miktarını gösterir. 

> Cap() : kanalın ilk tanımlandığı andaki belirtilen maksimum veri miktarını gösterir.

```
func main() {
	myChannel := make(chan string, 3)

	myChannel <- "first data"
	myChannel <- "second data"

	fmt.Println("Capasity: ", cap(myChannel))
	fmt.Println("Length: ", len(myChannel))
	fmt.Println("Reading Data: ", <-myChannel)
	fmt.Println("New Data Length: ", len(myChannel))
}
```

Çıktı: 
```
Capasity:  3
Length:  2
Reading Data:  first data
New Data Length:  1
```
