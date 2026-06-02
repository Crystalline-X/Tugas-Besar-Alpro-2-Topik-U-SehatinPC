### Tugas-Besar-Alpro-2-Topik-U-SehatinPC
SehatinPC adalah aplikasi untuk mengawasi kondisi operasional perangkat keras komputer
secara real-time. Data utama yang digunakan adalah data komponen, data suhu sensor, dan data
penggunaan beban kerja. Pengguna aplikasi adalah teknisi komputer atau pemilik sistem PC.

## PENJELASAN SPESIFIKASI
a. Pengguna dapat menambahkan, mengubah, dan menghapus data komponen PC yang terpasang.
b. Sistem dapat mencatat status kondisi perangkat terutama saat mengalami lag atau panas berlebih (overheat).
c. Pengguna dapat mencari data komponen berdasarkan nama perangkat atau status kesehatan menggunakan Sequential dan Binary Search.
d. Pengguna dapat mengurutkan data perangkat berdasarkan nomor seri komponen menggunakan Selection dan Insertion Sort.
e. Sistem dapat menampilkan statistik jumlah komponen yang bermasalah dan rata-rata suhu kerja perangkat.

## MENU UTAMA 
1. Tambah Komponen: Menambahkan data hardware baru (SN, Nama, Suhu) dengan penentuan status otomatis (Normal/Lag/Overheat).
2. Ubah Komponen: Memperbarui data Nama dan Suhu komponen berdasarkan Serial Number (SN).
3. Hapus Komponen: Menghapus data komponen dari sistem berdasarkan pencarian SN.
4. Tampilkan Semua Data: Menampilkan daftar lengkap seluruh komponen beserta kondisinya ke layar terminal.
5. Cari Komponen (Sequential Search): Mencari data spesifik berdasarkan Nama secara linear.
6. Cari Komponen (Binary Search): Pencarian data super cepat berdasarkan Nama (program otomatis mengurutkan abjad A-Z di latar belakang).
7. Urutkan Data (Selection Sort): Mengurutkan daftar komponen dari SN terkecil ke terbesar secara ascending.
8. Urutkan Data (Insertion Sort): Alternatif algoritma pengurutan untuk menyusun data dari SN terkecil ke terbesar.
9. Statistik Kesehatan: Menampilkan laporan ringkas berisi jumlah komponen yang bermasalah dan rata-rata suhu sistem.
0. Keluar: Menutup aplikasi dengan aman (data in-memory akan otomatis dibersihkan).
