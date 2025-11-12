#!/bin/bash

ulang="y"

while [ $ulang = "y" ] || [ $ulang = "Y" ]
do
  echo -e "\e[95mMasukkan Nama Produk: \e[96m\c"
  read nama
  echo -e "\e[95mMasukkan Harga Produk: \e[96m\c"
  read harga
  echo -e "\e[95mMasukkan Diskon (%): \e[96m\c"
  read diskon

  if [ -z "$diskon" ] || [ "$diskon" -eq 0 ]; then
    potongan=0
    harga_akhir=$harga
  else
    potongan=$((harga * diskon / 100))
    harga_akhir=$((harga - potongan))
  fi

  echo -e "\n\e[92m=== Detail Diskon ===\e[0m"
  echo -e "\e[96mNama Produk   : \e[0m$nama"
  echo -e "\e[96mHarga Awal    : \e[0mRp $harga"
  echo -e "\e[96mDiskon        : \e[0m$diskon%"
  echo -e "\e[96mPotongan Harga: \e[0mRp $potongan"
  echo -e "\e[96mHarga Akhir   : \e[0mRp $harga_akhir"

  echo -e "\n\e[92mApakah ingin menghitung produk lain? (y/n): \e[0m\c"
  read ulang
done

echo -e "\n\e[93mTerima kasih telah menggunakan program ini!\e[0m"
