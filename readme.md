# macam macam pendeklarasian variabel

untuk mendeklarasikan variabel di GO, maka ketika di deklarasi, tipe data yang digunakan harus dituliskan juga, istilah lain dari konsep ini adalah **mainfest typing**.

untuk mendeklarasikan sebuah variabel adalah dengan menggunakan keyword `var` diikuti nama variabel dan tipe data.

1. skema penggunaan keyword `var`

```go
var <nama-variabel> <tipe-data>
var <nama-variabel> <tipe-data> = <nilai>
```

contoh pembuatan variabel yang tipe datanya harus ditulis.

```go
package main

import "fmt"

func main(){
    // deklarasi variabel beserta tipe datanya
    var firstName string = "fani"
    var lastName string
    lastName = "alfirdaus"

    fmt.Printf("Hello %s %s\n",firstName,lastName)
}
```

ketika dijalankan program di atas akan mengeluarkan output

> Hello fani alfirdaus

- penggunaan fungsi `fmt.Printf()`

digunakan untuk menampilkan output dalam bentuk tertentu, kegunaannya sama seperti `fmt.Println()` hanya saja struktur outputnya didefinisikan di awal.

selain mendeklarasikan variabel dengan konsep **mainfest typing** GO juga mengadopsi konsep **type inference**, yaitu deklarasi variabel yang tipe datanya ditentukan oleh tipe data nilai-nya. Dengan metode jenis ini keyword `var` dan _tipe data_ tidak perlu lagi di tulis.

```go
var namaDepan string = "fani" // deklarasi variabel dengan konsep mainfest typing
namaBelakang := "alfirdaus" // deklarasi variabel dengan konsep type inference
```

pada penggunaan konsep **type inference** operand `=` harus di ganti dengan `:=` dan keyword `var` dihilangkan.

jika pendeklarasian variabel dengan menggunakan konsep **type inference** juga diperbolehkan menggunakan keyword `var` meskipun tanpa menuliskan tipe data-nya, dengan ketentuan tidak menggunakan operand `:=` tapi dengan tetap menggunakan operand `=`.

namun penggunaan `:=` hanya digunakan sekali saat pertama kali pendeklarasian variabel

```go
nama := "fani"
nama = "alfi"
nama = "fanialfi"
```

GO juga mendukung metode deklarasi banyak variabel sekaligus, caranya dengan menuliskan nama variabel lalu diikuti dengan tanda `,` yang berguna sebagai pemisah antara satu variabel dengan variabel yang lain. Selain itu juga dapat langsung di isi value dari variabel nya, dengan menggunakan tanda `,` sebagai pemisah antara value satu dengan value yang lain.

pendeklarasian multiple variabel dengan konsep **mainfest typing**

```go
var firstName,lastName string = "fani","alfirdaus"
```

pendeklarasian multiple variabel dengan konsep **type inference**

```go
firstName,lastName := "fani","alfirdaus"
```

dengan menggunakan konsep **type inference** kita juga bisa mendeklarasikan multiple variabel dan multiple tipe data

```go
firstName,lastName,age := "fani","alfirdaus",17
```

GO memiliki aturan yang unik dimana tidak dimiliki oleh bahasa pemrograman lain, yaitu tidak boleh ada variabel yang menganggur (semua variabel yang di deklarasikan harus digunakan), jika ada variabel yang didefinisikan tapi tidak digunakan maka akan muncul error saat proses kompilasi program, dan program tidak akan bisa di run.

_underscore_ (`_`) adalah _reserved variabel_ yang bisa digunakan untuk menampung nilai yang tidak di pakai, atau bisa di bilang variabel ini merupakan keranjang sampah.

```go
_ = "Belajar Golang"
_ = "Apa Belajar Golang Itu Susah ?"
name,_ := "fani","Golang"
```

pada contoh di atas variabel `name` akan berisi _fani_ sedangkan nilai "Golang" ditampung oleh variabel _underscore_, yang menandakan bahwa nilai tersebut tidak akan digunakan.

Biasanya variabel _underscore_ sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan, dan isi dari variabel _underscore_ tidak bisa ditampilkan, data yang sudah masuk kedalam variabel tersebut akan hilang seperti blackhole.

- deklarasi variabel dengan keyword `new`

digunakan untuk membuat variabel **pointer** dengan tipe data tertentu. Nilai data default-nya akan menyesuaikan tipe data-nya.

```go
var name = new(string)
fmt.Println(name) // 0xc000014320
fmt.Println(*name) // ""
```

variabel `name` diatas menampung data bertipe **pointer string**, jika di tampilkan yang muncul bukanlah nilai-nya, melainkan alamat memori nilai tersebut (dalam bentuk notasi hexadecimal). Untuk menampilkan nilai aslinya, variabel tersebut perlu di-**dereferensi** terlebih dahulu dengan menggunakan tanda asterik (`*`).

- deklarasi variabel dengan keyword `make`

keyword ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu :

1. channel
2. slice
3. map

2. konstanta

merupakan jenis variabel yang tidak bisa diubah nilainya, inisialisasi nilai hanya bisa dilakukan di awal saat pendeklarasian variabel.

penggunaan atau penerapa variabel konstanta sama seperti penerapan variabel dengan menggunakan keyword `var` yang membedakan ialah hanya perlu mengganti keyword `var` dengan `const`.