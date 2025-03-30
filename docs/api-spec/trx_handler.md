# Dokumentasi API Transaksi (trx)

## 1. Mendapatkan Semua Transaksi
### Endpoint
```
GET /api/v1/trx?limit={limit}&page={page}
```

### Header
```
token:  <JWT_TOKEN>
```

### Parameter Query
| Nama     | Tipe   | Deskripsi                         |
|----------|--------|---------------------------------|
| limit    | int    | Jumlah transaksi per halaman   |
| page     | int    | Halaman yang ingin diambil     |

### Contoh Request
```
curl --location 'localhost:8000/api/v1/trx?limit=10&page=1' \
tokenheader 'Authorization:  <JWT_TOKEN>'
```

### Contoh Response
```json
{
    "status": true,
    "message": "Succeed to GET data",
    "errors": null,
    "data": {
        "data": [
            {
                "id": 1,
                "harga_total": 1010000,
                "kode_invoice": "INV-1743307505207",
                "method_bayar": "bca",
                "alamat_kirim": {
                    "id": 1,
                    "judul_alamat": "Fort Hulda",
                    "nama_penerima": "Judy Cormier",
                    "no_telp": "(791) 961-0289 x04643",
                    "detail_alamat": "Valentine Trafficway, 3375 Schuster Shoals, North Emilyberg"
                },
                "detail_trx": [
                    {
                        "product": {
                            "id": 1,
                            "nama_produk": "Oriental Fresh Chicken",
                            "slug": "oriental-fresh-chicken-1743307426174",
                            "harga_reseller": "5000",
                            "harga_konsumen": "10000"
                        },
                        "kuantitas": 1,
                        "harga_total": 10000
                    }
                ]
            }
        ],
        "page": 1,
        "limit": 10
    }
}
```

---

## 2. Mendapatkan Transaksi Berdasarkan ID
### Endpoint
```
GET /api/v1/trx/{id}
```

### Header
```
token:  <JWT_TOKEN>
```

### Parameter Path
| Nama | Tipe | Deskripsi               |
|------|------|-------------------------|
| id   | int  | ID transaksi yang dicari |

### Contoh Request
```
curl --location 'localhost:8000/api/v1/trx/1' \
tokenheader 'Authorization:  <JWT_TOKEN>'
```

### Contoh Response
```json
{
    "status": true,
    "message": "Succeed to GET data",
    "errors": null,
    "data": {
        "id": 1,
        "harga_total": 1010000,
        "kode_invoice": "INV-1743307505207",
        "method_bayar": "bca",
        "alamat_kirim": {
            "id": 1,
            "judul_alamat": "Fort Hulda",
            "nama_penerima": "Judy Cormier",
            "no_telp": "(791) 961-0289 x04643",
            "detail_alamat": "Valentine Trafficway, 3375 Schuster Shoals, North Emilyberg"
        },
        "detail_trx": [
            {
                "product": {
                    "id": 1,
                    "nama_produk": "Oriental Fresh Chicken",
                    "harga_reseller": "5000",
                    "harga_konsumen": "10000"
                },
                "kuantitas": 1,
                "harga_total": 10000
            }
        ]
    }
}
```

---

## 3. Menambahkan Transaksi Baru
### Endpoint
```
POST /api/v1/trx
```

### Header
```
token:  <JWT_TOKEN>
Content-Type: application/json
```

### Body Request
```json
{
    "harga_total": 50000,
    "method_bayar": "bca",
    "alamat_kirim": {
        "id": 1
    },
    "detail_trx": [
        {
            "product_id": 1,
            "kuantitas": 5,
            "harga_total": 50000
        }
    ]
}
```

### Contoh Request
```
curl --location 'localhost:8000/api/v1/trx' \
tokenheader 'Authorization:  <JWT_TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "harga_total": 50000,
    "method_bayar": "bca",
    "alamat_kirim": {
        "id": 1
    },
    "detail_trx": [
        {
            "product_id": 1,
            "kuantitas": 5,
            "harga_total": 50000
        }
    ]
}'
```

### Contoh Response
```json
{
    "status": true,
    "message": "Succeed to CREATE data",
    "errors": null,
    "data": {
        "id": 5,
        "harga_total": 50000,
        "kode_invoice": "INV-1743307605259",
        "method_bayar": "bca",
        "alamat_kirim": {
            "id": 1
        },
        "detail_trx": [
            {
                "product_id": 1,
                "kuantitas": 5,
                "harga_total": 50000
            }
        ]
    }
}
```

---
