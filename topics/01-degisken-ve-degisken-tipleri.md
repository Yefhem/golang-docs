# Değişken ve Değişken Tipleri
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



