package main

import "fmt"

func deklarasiVar(){
  // jika menggunakan operand `:=` maka hanya digunakan sekali saat pendeklarasian variabel 
  nama := "fani"
  fmt.Println(nama)
  nama = "alfi"
  fmt.Println(nama)
  nama = "fanialfi"
  fmt.Println(nama)
}

func multiVar(){
  // deklarasi multiple variabel
  var first,middle,last int 
  first = 1
  middle = 2
  last = 3
  fmt.Println(first,middle,last)


  var satu,dua,tiga = "satu","dua","tiga"
  fmt.Println(satu,dua,tiga)

  var string,boolean,number = "ini string",true,20 // deklarasi multiple variabel & multiple tipe data
  fmt.Println(string,boolean,number)
}

func underscore(){
  var nama = "fani alfirdaus"
  _ = "belajar GoLang"

  fmt.Printf("Hi My Name Is %s\n",nama)
}

func varNew(){
  // deklarasi variabel dengan keyword `new` topik ini akan dibahas di project lain
  var name = new(string);
  fmt.Println(name) // 0xc000014320
  fmt.Println(*name) // ""
}

func main(){
  // deklarasi variabel dengan konsep mainfest typing
  var firstName string = "fani"
  var lastName string 
  lastName = "alfirdaus"

  umur := 17 // deklarasi variabel dengan konsep type inference
  var alamat = "Indonesia" // deklarasi variabel dengan konsep type inference tapi tetap dengan menggunakan keyword var

  fmt.Printf("Hallo saya %s %s, umur saya %d dan alamat saya di %s\n",firstName,lastName,umur,alamat)

  deklarasiVar()
  multiVar()
  underscore()
  varNew()
}
