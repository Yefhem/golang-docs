# For Döngüsü (For Loop)

Belirli koşullar altında ve belirli sayıda üst üste çalıştırmak için kullandığımız bir koşullu ifade çeşitidir.

Go dili içerisinde for loop'u farklı şekillerde kullanabiliyoruz. Bilindiği üzere go dilinde **while** döngüsü bulunmamaktadır.  Daha doğrusu bir keyword olarak bulunmamaktadır. While döngüsü ihtiyacı for loop ile giderilir. Konu içerisinde buna da değineceğiz.

## Üç Bileşenli Temel For Loop Döngüsü (Three-component Loop)

For döngüsü denilince temel ve ilk akla gelen for loop çeşiti bu olsa gerek. Yapısına bakacak olursak:

`for <init statement> ; <condition> ; <post statement> {}`

İlk olarak bir başlangıç değeri belirliyoruz. Daha sonra şart ve son olarak da döngünün nasıl devam edeceğini belirleme kalıyor.

```
func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}
```

Çıktı:
```
1
2
3
4
5
6
7
8
9
```


**Not:** Başlangıç değerini for döngüsü içerisinde initialize etme gibi bir zorunluluğumuz bulunmamaktadır. 

```
func main() {
	a := 5

	for ; a < 9; a++ {
		fmt.Println(a)
	}
}
```

Çıktı:
```
5
6
7
8
```

## While Loop

Bir for döngüsünde init statement ve post statement bölümlerini atlayıp sadece condition kullanarak bir while döngüsü elde edilebilir. Yapı olarak bakacak olursak:

`for <condition> {}`

```
func main() {
	n := 2
	for n < 9 {
		n *= 2
	}
	fmt.Println(n)
}
```

Çıktı:
```
16
```

## Sonsuz Döngü (Infinite Loop)

Koşul ifadesini de çıkarttığımız zaman sonsuz döngü elde etmiş oluruz. Yapı olarak bakacak olursak:

`for {}`

```
func main() {
	number := 0
	for {
		number++
		fmt.Println(number)
	}
}

```

Çıktı:
```
1
2
3
.
.
.
```

## For-each Range Loop

Slices, arrays, maps, channels or strings bu elementlerin üzerinde döngü yapmak for-each ile daha iyi yapılır.

```
func main() {
	fruits := []string{"apple", "banana", "cherry"}
	for i, f := range fruits {
		fmt.Println(i, f)
	}
}
```

Çıktı:
```
0 apple
1 banana
2 cherry
```

## Break and Continue

Go içerisinde **break** ve **continue** anahtar kelimeleri mevcuttur. Bir continue ifadesi for loop içerisindeki  bir sonraki yinelemeye gönderir. Bir bakıma (i++) işlemi yapmış olur. Break ise  döngüyü tamamen keser.

```
func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			continue
		}
		if i == 14 {
			break
		}
		fmt.Println(i)
	}
}
```

Çıktı:
```
1
2 
4 
5 
7 
8 
10
11
13
```

## Örnekler 

`1 ile 50 arasındaki asal sayıları gösteren kod bloğu:`

```

```

Çıktı:
```

```

