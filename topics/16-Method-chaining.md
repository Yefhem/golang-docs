# Method Zincirleme (Method Chaining)

Aynı nesne referansı ile farklı metodları ayrı ayrı çağırmak yerine, tek bir expression içinde farklı methodları çağırma uygulamasıdır.

Method zincirlemenin mümkün olması için, zincirdeki methodların nesne referansını döndürmesi gerekir. Zincirin son halkası için referansın döndürülmesi isteğe bağlıdır.

## Golang'de Zincirleme Metodu Uygulama (Implementing Method Chaining in Golang) 

```
type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) printName() employee {
	fmt.Printf("Name: %s\n", e.name)
	return e
}

func (e employee) printAge() employee {
	fmt.Printf("Age: %d\n", e.age)
	return e
}

func (e employee) printSalary() {
	fmt.Printf("Age: %d\n", e.salary)
}

func main() {
	emp := employee{name: "sam", age: 31, salary: 2000}
	emp.printName().printAge().printSalary()
}
```

Çıktı:
```
Name: sam
Age: 31  
Age: 2000
```

**Example 2**

```
type Movie struct {
	Name      string
	Year      int
	Creator   string
	MovieType string
}

func (u Movie) SetName(name string) Movie {
	u.Name = name
	return u 
}

func (u Movie) SetYears(year int) Movie {
	u.Year = year
	return u 
}

func (u Movie) SetCreators(creator string) Movie {
	u.Creator = creator
	return u 
}

func (u Movie) SetMovieType(movieType string) Movie {
	u.MovieType = movieType
	return u 
}

func (u Movie) Print() {
	fmt.Printf("Movie Name:%v\nYear:%v\nCreator:%v\nType:%v", u.Name, u.Year, u.Creator, u.MovieType)
}

func main() {
	movie := Movie{}
	movie.SetName("In Time").
	SetYears(2011). 
	SetCreators("Andrew Niccol"). 
	SetMovieType("Action").Print()
}
```

Çıktı:
```
Movie Name:In Time   
Year:2011
Creator:Andrew Niccol
Type:Action
```

**Example 3**
```
type Card struct {
	amount   float64
	currency string
	address  string
}

func (c *Card) Charge(amount float64) *Card {
	c.amount = amount
	return c
}

func (c *Card) WithCurrency(currency string) *Card {
	c.currency = currency
	return c 
}

func (c *Card) WithAddress(address string) *Card {
	c.address = address
	return c 
}

func (c *Card) Execute() {
	fmt.Printf("Dear Customer %s, \n%s %v is Debited from your account ", c.address, c.currency, c.amount)
}

func main() {
	card := Card{}
		card.Charge(10.4).WithCurrency("USD").WithAddress("Samet Celik").Execute()
}
```

Çıktı:
```
Dear Customer Samet Celik, 
USD 10.4 is Debited from your account
```

## Neden Method Zincirleme Tercih Edilir?

* Gorm gibi go kitaplıkları, daha karmaşık sql sorguları oluşturmak için method zincirlemeye başvurur.
* Bazı geliştiriciler, kaynak kodun okunabilirliğini arttırmak için method zincirlemeyi teşvik eder.
* Fazlalık oluşturacak geçici değişkenleri, kod tekrarlarını ortadan kaldırır.
* İfadeyi doğal dil metinleri gibi soldan sağa doğru okunmasına olanak tanır.