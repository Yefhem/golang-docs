# Maps

Key-value çiftleri şeklinde elemanlar barındıran veri yapıları olarak düşünebiliriz. 
Bazı programlama dillerinde bulunan hash table kavramına denk bir veri tipidir.
Map veri tipi, slice ve dizilere göre daha kullanışlı olsa da kaynak açısından biraz daha masraflıdır.
Map veri tipi referans türdür. Birkaç farklı şekilde map oluşturabiliriz.


`var map1 map[int]int`

`map2 := map[int]string{}`

`var map3 = make(map[int]bool)`

Başlangıç değerleri atayarak da map'ler oluşturulabilir.

```
func main() {
	var map1 = map[int]bool{
		1: true,
		2: false,
	}

	fmt.Println("Map1 : ", map1)
}
```

Çıktı:
```
Map1 :  map[1:true 2:false]
```

Map'lerde veriyi alabilir, yeni veri ekleyebilir, güncelleyebilir ve silebiliriz.

```
func main() {
	m := make(map[int]string)

	// Yeni bir key-value çifti ekleme
	m[1] = "steve"
	m[2] = "jobs"
	m[3] = "apple"
	fmt.Println(m)

	// Güncelleme
	m[2] = "wozniak"
	fmt.Println(m)

	// Veri alma
	v := m[3]
	fmt.Println(v)

	// Silme işlemi
	delete(m, 1)

	fmt.Println(m)
}
```

Çıktı:
```
map[1:steve 2:jobs 3:apple]
map[1:steve 2:wozniak 3:apple]
apple                         
map[2:wozniak 3:apple] 
```

Map'ler for range kullanımı için uygun ve idealdir.

```
func main() {
	var myMap = map[string]string{
		"a": "black",
		"b": "blue",
		"c": "orange",
		"d": "purple",
		"e": "red",
		"f": "yellow",
	}

	for i, v := range myMap {
		fmt.Println(i, v)
	}
```

Çıktı:
```
e red
f yellow
a black 
b blue  
c orange
d purple
```