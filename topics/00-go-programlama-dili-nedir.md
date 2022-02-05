## Go Programlama Dili Nedir?

Go, bilinen diğer ismiyle Golang, Google mühendisleri tarafından geliştirilmiş ve hala geliştirilmekte olan açık kaynaklı programlama dilidir. Golang şuanda hali hazırda Google'daki yazılım geliştirme sürecinde ve birçok açık kaynaklı projede kullanılmaktadır.

Tarihçesine bakacak olursak; Golang 2007 yılında Robert Griesemar, Rob Pike ve Ken Thompson tarafından bir deney olarak geliştirilmeye başlanmış , 2009 yılında ise açık kaynaklı bir programlama dili olarak piyasaya sürülmüştür.

Google kendi iç sistemlerinde yıllardır kullandığı çeşitli programlama dili ve teknolojilerin avantajlı yönlerini alarak ortaya basit , güvenilir ve verimli yazılım oluşturmayı kolaylaştıran bir programlama dili çıkarmıştır.

## Go Dilinin Özellikleri

- Go, derlenebilir ve statik bir dildir. Hızlı derlenme süresi vardır. Yorumlanabilir diller gibi sanal makineye ihtiyaç duymadığından, direkt olarak doğal makine diline çevrilebilir. Bu özellik ekstra zaman tasarrufu sağlar.
- 25 keyword bulundurması Golang'in sade bir dil olduğunun göstergesidir.
- Go esnek bir dildir, pek çok alanda kullanıma uygundur. Sistem ve ağ programlama, makine öğrenmesi, big data(büyük veri), web, mobil, CLI ve Desktop gibi alanlarda geliştirmeler yapılabilir.
- Go, kendi için ayrılan belleğin yönetimini sağlayarak programların düzgün çalışmasını sağlamasının yanı sıra kendi çöp toplayıcısına sahiptir. (Garbage Collector)
- Go dili kendi içinde gömülü olarak concurrency(eş zamanlılık) destekler. Diğer diller gibi sonradan eklenmediği için yüksek performanslı olarak gerçekleştirir.
- Paket yöneticisi `go get` dir.
- Maskotu **Gopher** adı verilen sevimli bir sincaptır. Go dilini geliştirenlere de Gopher denilmektedir.

## Örnek Go Kodu

```
package main

import "fmt"

func main() {
	fmt.Println("We Good!")
}
```