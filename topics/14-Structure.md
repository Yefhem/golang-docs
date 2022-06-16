# Structure (Struct)

Go'da nesne tabanlı bir yapı yoktur. Struct kavramı genellikle nesne yönelimli programlama dillerindeki class'larla karşılaştırılır. 
Fakat struct'lar tam anlamıyla class'ları karşılamaz.

Go'da struct nedir? Farklı türlerdeki yapıları tek bir türde gruplandırmaya/birleştirmeye izin veren bir veri tipidir. 

Yeni bir struct oluşturmak istersek;

```
type Example struct {
    name    string 
    surname string
    age     int
}
```

Aslında burada kendi veri tipimiz struct veri tipi üzerine kurulu. Burada yeni bir veri tipi tanımlamış olduk diyebiliriz.

Bir struct oluşturup veri ataması nasıl yapılır?

```
type Example struct {
	name string
}

func main() {
	n := Example{
		name: "steve",
	}
	fmt.Println(n)
	fmt.Println(n.name)
}
```

Çıktı:
```
{steve}
steve
```

veya 

```
type Example struct {
	name string
}

func main() {
	n := Example{}

	n.name = "jobs"

	fmt.Println(n)
	fmt.Println(n.name)
}
```

Çıktı:
```
{jobs}
jobs
```

Go'da belli bir ölçüde çok biçimliliği(Polymorphism) de uygulayabiliriz.

```
type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

func main() {
	employee := employee{
		name:      "steve",
		age:       20,
		isMarried: true,
	}

	manager := manager{
		employee:  employee,
		hasDegree: true,
	}

	fmt.Println(employee)
	fmt.Println(manager)
}
```

Çıktı:
```
{steve 20 true}
{{steve 20 true} true}
```