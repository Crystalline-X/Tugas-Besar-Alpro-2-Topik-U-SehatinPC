## Tugas-Besar-Alpro-2-Topik-U-SehatinPC
SehatinPC adalah aplikasi untuk mengawasi kondisi operasional perangkat keras komputer
secara real-time. Data utama yang digunakan adalah data komponen, data suhu sensor, dan data
penggunaan beban kerja. Pengguna aplikasi adalah teknisi komputer atau pemilik sistem PC.

## PENJELASAN SPESIFIKASI
a. Pengguna dapat menambahkan, mengubah, dan menghapus data komponen PC yang
terpasang.
b. Sistem dapat mencatat status kondisi perangkat terutama saat mengalami lag atau panas
berlebih (overheat).
c. Pengguna dapat mencari data komponen berdasarkan nama perangkat atau status kesehatan
menggunakan Sequential dan Binary Search.
d. Pengguna dapat mengurutkan data perangkat berdasarkan nomor seri komponen
menggunakan Selection dan Insertion Sort.
e. Sistem dapat menampilkan statistik jumlah komponen yang bermasalah dan rata-rata suhu
kerja perangkat. Tampilkan

## MENU UTAMA 
1. Tambah Komponen
2. Ubah Komponen
3. Hapus Komponen
4. Tampilkan Semua Data
5. Cari Komponen (Sequential Search - By Name)
6. Cari Komponen (Binary Search - By Name)
7. Urutkan Data (Selection Sort - By Serial)
8. Urutkan Data (Insertion Sort - By Serial)
9. Statistik Kesehatan
0. Keluar

## PENJELASAN TIAP MENU
1. Tambah Komponen
Opsi ini berfungsi untuk memasukkan data perangkat keras baru ke dalam sistem pengawasan. Anda akan diminta untuk mengetikkan Serial Number, Nama, dan Suhu komponen. Sistem kemudian akan mengevaluasi suhu tersebut dan secara otomatis memberikan label "Status" (Normal, Lag, atau Overheat) sebelum menyimpannya ke dalam memori.

3. Ubah Komponen
Opsi ini digunakan ketika Anda ingin memperbarui data komponen yang sudah terdaftar. Anda harus memasukkan Serial Number dari komponen target. Jika ditemukan, sistem akan meminta Anda memasukkan Nama dan Suhu yang baru, lalu menghitung ulang status kesehatannya berdasarkan suhu terbaru tersebut.

4. Hapus Komponen
Berfungsi untuk membuang atau menghapus data komponen dari memori program. Anda hanya perlu memberikan Serial Number target. Program akan mencari nomor tersebut dan menghapusnya secara permanen dari daftar pengawasan.

5. Tampilkan Semua Data
Opsi ini akan mencetak seluruh daftar komponen yang saat ini tersimpan di dalam memori ke layar terminal Anda. Informasi yang ditampilkan mencakup Serial Number, Nama, Status, dan Suhu dengan format yang rapi dan mudah dibaca.

6. Cari Komponen (Sequential Search - By Name)
Fitur pencarian dasar berdasarkan "Nama" komponen. Program akan mengecek daftar komponen satu per satu dari urutan paling awal hingga akhir secara linear sampai menemukan nama yang persis sama dengan yang Anda ketikkan. Metode ini cocok digunakan meskipun data masih dalam keadaan acak atau belum diurutkan.

7. Cari Komponen (Binary Search - By Name)
Fitur pencarian tingkat lanjut yang juga berdasarkan "Nama", namun jauh lebih cepat untuk jumlah data yang besar. Karena algoritma Binary Search mewajibkan data harus terurut, program akan secara otomatis mengurutkan nama komponen sesuai abjad (A-Z) terlebih dahulu di belakang layar sebelum melakukan pencarian dengan metode belah tengah (divide and conquer).

8. Urutkan Data (Selection Sort - By Serial)
Opsi ini akan merapikan dan menyusun ulang seluruh daftar komponen dari Serial Number terkecil hingga terbesar (ascending). Ia menggunakan algoritma Selection Sort, yaitu dengan cara mencari nilai paling kecil di dalam daftar yang belum terurut, lalu memindahkannya ke posisi paling depan secara berulang-ulang.

9. Urutkan Data (Insertion Sort - By Serial)
Sama seperti opsi nomor 7, opsi ini bertujuan mengurutkan daftar dari Serial Number terkecil ke terbesar, namun menggunakan mesin algoritma yang berbeda (Insertion Sort). Cara kerjanya mirip dengan menyusun kartu remi di tangan, yaitu mengambil satu per satu data dan menyisipkannya ke posisi yang tepat di antara data lain yang sudah terurut.

10. Statistik Kesehatan
Menampilkan ringkasan atau rangkuman kondisi sistem secara keseluruhan. Program akan menghitung dan menampilkan berapa banyak jumlah komponen yang sedang bermasalah (berstatus "Lag" atau "Overheat"), serta menghitung rata-rata suhu dari seluruh komponen yang terdaftar.

0. Keluar
Opsi untuk menghentikan perulangan menu dan menutup aplikasi secara aman. Perlu diingat bahwa karena program ini menggunakan penyimpanan sementara (In-Memory), memilih opsi keluar akan membuat semua data komponen yang sudah diinputkan menghilang secara otomatis.
