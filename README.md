nomer 1

Fungsi fibonacci:
- Definisi Rekursif: Fungsi ini memanggil dirinya sendiri untuk menghitung nilai Fibonacci. Ini adalah cara yang umum untuk mendefinisikan barisan Fibonacci secara matematis.
- Kasus Dasar: Jika nilai n yang diberikan adalah 0 atau 1, fungsi langsung mengembalikan nilai tersebut. Ini adalah titik akhir dari rekursi.
- Kasus Rekursif: Jika n lebih dari 1, fungsi akan menghitung nilai Fibonacci dengan menjumlahkan hasil dari memanggil dirinya sendiri dengan n-1 dan n-2. Ini sesuai dengan definisi barisan Fibonacci: setiap bilangan adalah jumlah dari dua bilangan sebelumnya.

Fungsi main:
- Perulangan: Program akan mengulang proses perhitungan sebanyak 11 kali, mulai dari indeks 0 hingga 10.
- Pemanggilan Fungsi: Pada setiap iterasi, fungsi fibonacci dipanggil dengan nilai indeks saat ini (i).
- Penampilkan Hasil: Hasil perhitungan (nilai Fibonacci) kemudian dicetak ke layar bersama dengan indeksnya.

nomer 2

Cara Kerja:

Input Nilai N:
- Program meminta pengguna untuk memasukkan nilai N yang menentukan jumlah baris segitiga.

Fungsi printStar:
- Rekursi: Fungsi ini menggunakan rekursi untuk mencetak setiap baris segitiga.
- Kasus Dasar: Jika n sama dengan 0, fungsi berhenti. Ini adalah kondisi akhir rekursi.
- Kasus Rekursif:
   - Fungsi memanggil dirinya sendiri dengan n-1 untuk mencetak baris berikutnya.
   - Kemudian, menggunakan loop for, fungsi mencetak n buah bintang pada baris tersebut.
   - Setelah mencetak bintang, fungsi mencetak karakter newline untuk pindah ke baris berikutnya.
 
nomer 3

Fungsi faktor:
- Rekursif: Fungsi ini menggunakan rekursi untuk mencari faktor secara efisien.
- Parameter: Fungsi menerima dua parameter:
   - n: Bilangan yang akan dicari faktornya.
   - i: Bilangan pembagi yang akan dicoba. Awalnya, i dimulai dari 1.
- Logika:
   - Kondisi Berhenti: Jika i lebih besar dari n, berarti semua kemungkinan pembagi sudah dicoba, dan fungsi mengembalikan nil (slice kosong).
   - Kondisi Faktor: Jika n habis dibagi i, maka i adalah faktor dari n. Fungsi kemudian menambahkan i ke dalam slice hasil dan melanjutkan pencarian faktor dengan i+1.
   - Lanjutkan Pencarian: Jika i bukan faktor dari n, fungsi melanjutkan pencarian dengan i+1.

Fungsi main:
- Input: Program meminta pengguna memasukkan bilangan bulat positif.
- Pemanggilan Fungsi: Fungsi faktor dipanggil dengan bilangan yang dimasukkan oleh pengguna dan nilai awal i sebesar 1.
- Output: Hasil dari fungsi faktor (yaitu, slice yang berisi semua faktor) dicetak ke layar.

nomer 4
Fungsi printSequence:
- Basis rekursif: Jika nilai n adalah 1, maka fungsi langsung mencetak 1 dan berhenti (rekursi berakhir).
- Kasus rekursif: Jika nilai n lebih dari 1, maka:
  - Mencetak nilai n saat itu.
  - Memanggil dirinya sendiri (rekursif) dengan nilai n-1, sehingga proses mencetak berulang untuk bilangan yang lebih kecil.
  - Kembali mencetak nilai n saat itu.
  
Fungsi main:
- Meminta pengguna memasukkan nilai N.
- Memanggil fungsi printSequence dengan nilai N yang telah diinputkan.

nomer 5
Fungsi main:
- Meminta pengguna untuk memasukkan nilai N.
- Memanggil fungsi printGanjil dengan nilai N yang diinputkan.
- Mencetak hasil dari fungsi printGanjil beserta label "Barisan bilangan ganjil:".

Fungsi printGanjil:
- Basis rekursif: Jika n kurang dari atau sama dengan 0, maka fungsi langsung mengembalikan string kosong. Ini menandai akhir dari rekursi.
- Kasus rekursif: Jika n lebih dari 0, maka:
  - Membentuk sebuah string dengan menggabungkan:
    - Nilai n (yang diasumsikan ganjil).
    - Spasi.
    - Hasil rekursif dari printGanjil(n-2). Artinya, fungsi akan dipanggil lagi dengan nilai n dikurangi 2 untuk mendapatkan bilangan ganjil berikutnya.

nomer 6
Fungsi main:
- Meminta pengguna untuk memasukkan nilai x dan y.
- Memanggil fungsi power dengan nilai x dan y yang diinputkan.
- Mencetak hasil dari fungsi power beserta label yang menjelaskan perhitungan.

Fungsi power:
- Basis rekursif: Jika y adalah 0, maka fungsi langsung mengembalikan 1. Ini adalah kasus dasar, karena bilangan apa pun pangkat 0 adalah 1.
- Kasus rekursif: Jika y lebih dari 0, maka:
  - Mengembalikan hasil perkalian antara x dan hasil rekursif power(x, y-1). Artinya, fungsi akan dipanggil lagi dengan nilai y dikurangi 1, sehingga proses perhitungan berulang hingga mencapai kasus dasar.
 
 
