# Dosya Yükleme (File Upload)

Go'da bir sunucuya dosya yükleme işlemi nasıl gerçekleştirilir bunu inceleyeceğiz. Dosya yükleme işlemi pek çok nedenden ötürü olabilir.

## Maksimum Dosya Boyutunu Ayarlama
İstemci tarafından yanlışlıkla veya kasıtlı olarak büyük boyutlarda dosyalar yüklendiğini ve sunucu kaynaklarının boşa harcandığı bir durumdan kaçınmak için maksimumu dosya yükleme boyutunu kısıtlamak gerekir. Burada farklı yöntemler uygulanabilir.

Max dosya boyutu belirlenip, Content Lenght ile istek başlığı kontrol edilip boyutun aşılıp aşılmadığı tespit edilebilir. Fakat bu yöntem gerçek dosya boyutundan bağımsız herhangi bir değer dönebileceği için güvenli bir yol olarak görülmez.

```
const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

func uploadHandler(rw http.ResponseWriter, r *http.Request) {
    if r.ContentLength > MAX_UPLOAD_SIZE {
        http.Error(w, "The uploaded image is too big. Please use an image less than 1MB in size", http.StatusBadRequest)
        return
    }
}
```

**http.MaxBytesReader** methodu daha güvenli bir yoldur. Bu yöntem gelen istek gövdelerinin boyutunu sınırlandırmak için kullanılır. Tekli dosya yüklemeleri için istek gövdesinin boyutunu sınırlandırmak, dosya boyutunu sınırlandırma konusunda iyi bir yaklaşımdır. 

```
const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

func uploadHandler(rw http.ResponseWriter, r *http.Request) {
    r.Body = http.MaxBytesReader(rw, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(rw, "The uploaded image is too big. Please use an image less than 1MB in size", http.StatusBadRequest)
		return
	}
}
```

## Yüklenen Dosyanın Kaydedilmesi
* Sıra geldi yüklenen dosyanın dosya sistemine kaydedilmesi, FormFile argümanı ön yüzden name attribute ile eşleşmelidir.

```
file, fileHeader, err := r.FormFile("file")
if err != nil {
	http.Error(rw, err.Error(), http.StatusBadRequest)
	return
}
```

* Eğer bir uploads klasörü mevcut değilse bu oluşturulmalıdır. Bunun için *os.MkdirAll* kullanabiliriz.

```
if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
	http.Error(rw, err.Error(), http.StatusInternalServerError)
	return
}
```

> os.ModePerm : Unix permission bits

* Uploads klasöründe yeni bir dosya oluşturma...

```
dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
if err != nil {
	http.Error(rw, err.Error(), http.StatusInternalServerError)
	return
}
```

> Burada dosyayı oluştururken isimlendirmeni uniq olması için time unix kullanır.

> Buradaki  **filepath.Ext()** metodu yüklenen dosyanın sonundaki uzantıyı döndürür (en sonki noktadan itibaren).

* Yüklenen dosyanın belirtilen hedefteki dosya sistemine kopyalanması...

```
_, err = io.Copy(dst, file)
if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
```

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", uploadHandler)

	if err := http.ListenAndServe(":4500", mux); err != nil {
		log.Fatal(err)
	}
}
```

## Yüklenen Dosyanın Türünün Kısıtlanması

Bazı durumlarda servere yüklenen dosyaları sınırlandırmak isteyebiliriz. Bunu yapabilmemiz için yüklenen dosyanın MIME türünü algılamamız ve ardından izin verilen MIME türüyle karşılaştırma yapmamız gerekiyor.

Kabul edilmesi gereken dosya türünü tanımlamak için ön uçtaki input elementinin accept attribute kullanabiliriz. Anacak girişin değiştirilip değiştirilmediği bilinemez. Yine sunucu tarafından kontrol edilmesi gerekir. İlgili kod bloğu;

```
buff := make([]byte, 512)
_, err = file.Read(buff)
if err != nil {
	http.Error(rw, err.Error(), http.StatusInternalServerError)
	return
}

filetype := http.DetectContentType(buff)
if filetype != "image/jpeg" && filetype != "image/png" { 
	http.Error(rw, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
	return
}

_, err = file.Seek(0, io.SeekStart)
if err != nil {
	http.Error(rw, err.Error(), http.StatusInternalServerError)
	return
}
```

* Http paketi tarafından sağlanan DetectContentType() methodu, verilen verinin içerik türünü algılamak amacıyla kullanılır.

* Burada oluşturduğumuz 512 Byte'lık veri `buff := make([]byte, 512)` MIME türünü belirlemek için kullanılır. Yani gelen verinin ilk 512 bayt ile MIME türü belirlenmeye çalışılır. Dosya DetectContentType() methoduna geçirilmeden önce dosyanın ilk 512 baytını boş bir arabelleğe okuruz. Ortaya çıkan dosya türü istenilen dosya türü ile uyuşmaz ise bir hata döndürülür.

* Yukarıda bahsettiğimiz üzere içerik türünü belirlemek için verinin ilk 512 baytını okumuştuk. Bu sebepten ötürü dosya akış işaretçisi (file stream pointer) 512 bayt ilerlemiş oldu. 

* Daha sonra **io.Copy()** methodu çağırıldığında bu konumdan itibaren okumaya devam ederse bozuk bir görüntü dosyası elde ederiz. 

* Tam bu noktada **file.Seek()** methodu dosya akış işaretçisinin dosyanın başına geri dönmesi için kullanılır.

**Kodun Tamamı**

```
const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "text/html")
	http.ServeFile(rw, r, "index.html")
}

func uploadHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(rw, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(rw, "The uploaded image is too big. Please use an image less than 1MB in size", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

## Birden Çok Dosyayı Yükleme

İstemciden aynı anda birden fazla dosya gönderildiği durumda **FormFile()** methodunu kullanmak yerine her dosyayı manuel olarak ayrıştırabiliriz. Dosya açma işleminden sonra yapmamız gereken tek şey dosya tek dosya yükleme ile aynıdır.

```
if err := r.ParseMultipartForm(32 << 20); err != nil {
	http.Error(rw, err.Error(), http.StatusInternalServerError)
	return
}
```

> 32 Mb, **FormFile()** tarafından kullanılan varsayılan değerdir.

**ParseMultipartForm()** çağırıldıktan sonra ön tarafta yüklenen dosyaları alacağız.

`files := r.MultipartForm.File["file"]`

Bu dosyaları bir for range ile dönerek, tek dosya yüklemedeki işlemlerin aynısını yapabiliriz.

**Kodun Tamamı**

```
func uploadHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["file"]
	values := r.MultipartForm.Value["value"]

	fmt.Println(values, files)

	for _, fileHeader := range files {

		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(rw, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(rw, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
	}
	fmt.Fprintf(rw, "Upload successful")
}
```