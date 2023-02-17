# Section 4 – Introduction to Algorithm and Golang

## Video – Introduction Algorithm 1 ( 5:23 menit )

**ALGORTIMA** adalah prosedur komputasi yang didefinisikan dengan baik yang mengambil beberapa nilai sebagai input dan output dan menghasilkan beberapa nilai sebagai output.

**Algoritma ex.** check bilangan prima, sorting, searching

**Karakteristik algoritma :**
1. memiliki batas (awal dan akhir)
2. instruksi terdefinisi dengan baik
3. efektif & efisien

**Jenis algortima :**
1. sequential = flow nya berurutan dari awal sampe akhir atau atas sampe bawah, semua bagian tereksekusi
2. branching = flow nya bercabang dari awal sampe akhir atau atas sampe bawah, ada bagian/cabang yang ditinggalkan karena tidak sesuai
3. looping = flow nya berulang pada baris tertentu hingga kondisinya memenuhi untuk menghentikan perulangan dan melanjutkan ke baris kode setelahnya
--- 
<br/>

## Video – Introduction Algorithm 2 ( 8:31 menit )

**PSEUDOCODE** adalah sebuah deskripsi penulisan algoritma dengan bahasa yang mudah dipahami dan bersifat umum


**FLOWCHART** adalah suatu bagan dengan simbol tertentu yang menggambar urutan dan hubungan antar baris proses secara mendetail.

<p align="center">
<b>Flowchart Symbol</b>
</p>
<p align="center">
<img title="flowchart-symbol" alt="flowchart-symbol"  width= "300px" src="images\flowchart_symbol.png">
</p>
<br/>

**Kasus 1** Menghitung luas segitiga :
1. INPUT `Alas` dan `Tinggi`
2. CALCULATE `Luas = (Alas x Tinggi) / 2`
3. PRINT `Luas`
   
<p align="center">
<img title="flowchart-symbol" alt="flowchart-symbol"  width= "100px" src="images\flowchart_luas-segitiga.png">
</p>
<br/>

**Kasus 2** Menentukan Bilangan Ganjil :

### bilanganGanjil.go
```golang
package functions

import (
	"fmt"
)

func BilanganGanjil(N []int) {
	for _, value := range N {
		if value%2 != 0 {
			fmt.Println(value)
		}
	}
}

```
<br/>

**Kasus 3** Mencetak Faktor Bilangan :

### faktorBilangan.go
```golang
package functions

import ("fmt")

func FaktorBilangan(N int){
	for i:=1; i<=N; i++{
		if N%i==0{
			fmt.Println(i)
		}
	}
}
```

---
<br/>

