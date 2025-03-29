## User API Spec

### Get My Profile

**Endpoint:** `GET /api/v1/user`

**Request Header:**
- `token: <JWT_TOKEN>`

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to GET data",
    "errors": null,
    "data": {
        "id": 1,
        "nama": "John Doe",
        "no_telp": "08123456789",
        "tanggal_lahir": "1990-01-01T00:00:00+07:00",
        "jenis_kelamin": "Laki-laki",
        "tentang": "",
        "pekerjaan": "Software Engineer",
        "email": "johndoe@mail.com",
        "id_provinsi": "11",
        "id_kota": "1101",
        "is_admin": false
    }
}
```

**Response Body (Failed):**
```json
{
    "status": false,
    "message": "Failed to GET data",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

---

### Update Profile

**Endpoint:** `PUT /api/v1/user`

**Request Header:**
- `token: <JWT_TOKEN>`

**Request Body:**
```json
{
    "nama": "John Doe Updated",
    "kata_sandi": "newpassword",
    "no_telp": "08123456789",
    "pekerjaan": "Senior Software Engineer",
    "email": "johnupdated@mail.com",
    "id_provinsi": "11",
    "id_kota": "1101"
}
```

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to UPDATE data",
    "errors": null,
    "data": {
        "id": 1,
        "nama": "John Doe Updated",
        "no_telp": "08123456789",
        "tanggal_lahir": "1990-01-01T00:00:00Z",
        "jenis_kelamin": "Laki-laki",
        "tentang": "",
        "pekerjaan": "Senior Software Engineer",
        "email": "johnupdated@mail.com",
        "id_provinsi": "11",
        "id_kota": "1101",
        "is_admin": false
    }
}
```

**Response Body (Failed - Validation Error):**
```json
{
    "status": false,
    "message": "Failed to UPDATE data",
    "errors": [
        "Key: 'UpdateProfileRequest.TanggalLahir' Error:Field validation for 'TanggalLahir' failed on the 'required' tag"
    ],
    "data": null
}
```

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Failed to UPDATE data",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

---

### Get My Address

**Endpoint:** `GET /api/v1/user/alamat?judul_alamat=dummy`

**Request Header:**
- `token: <JWT_TOKEN>`

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Failed to GET data",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

---

### Get Address by ID

**Endpoint:** `GET /api/v1/user/alamat/<id>`

**Request Header:**
- `token: <JWT_TOKEN>`

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to GET data",
    "errors": null,
    "data": {
        "id": 2,
        "judul_alamat": "Office",
        "nama_penerima": "Jane Doe",
        "no_telp": "08123456780",
        "detail_alamat": "Jalan Sudirman No. 456, Jakarta"
    }
}
```

**Response Body (Failed - Record Not Found):**
```json
{
    "status": false,
    "message": "Failed to GET data",
    "errors": [
        "record not found"
    ],
    "data": null
}
```

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Failed to GET data",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

---

### Create or Post Address

**Endpoint:** `POST /api/v1/user/alamat`

**Request Header:**
- `token: <JWT_TOKEN>`
- `Content-Type: application/json`

**Request Body:**
```json
{
    "judul_alamat": "Office",
    "nama_penerima": "John Doe",
    "no_telp": "08123456789",
    "detail_alamat": "5678 Oak Street, Springfield"
}
```

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to POST data",
    "errors": null,
    "data": {
        "id": 4,
        "judul_alamat": "Office",
        "nama_penerima": "John Doe",
        "no_telp": "08123456789",
        "detail_alamat": "5678 Oak Street, Springfield"
    }
}
```

**Response Body (Failed - Validation Error):**
```json
{
    "status": false,
    "message": "Failed to POST data",
    "errors": [
        "Key: 'CreateAddressRequest.NoTelp' Error:Field validation for 'NoTelp' failed on the 'required' tag"
    ],
    "data": null
}
```

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Failed to POST data",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

---

### Update Address
**Endpoint:** `PUT /api/v1/user/alamat/{id}`

**Request Header:**
- `token: <JWT_TOKEN>`

**Request Body:**
```json
{
    "nama_penerima": "John Doe",
    "no_telp": "08123456789",
    "detail_alamat": "Jl. Mawar No.10, Jakarta"
}
```

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to UPDATE data",
    "errors": null,
    "data": {
        "id": 1,
        "judul_alamat": "Rumah",
        "nama_penerima": "John Doe",
        "no_telp": "08123456789",
        "detail_alamat": "Jl. Mawar No.10, Jakarta"
    }
}
```

**Response Body (Failed - Validation):**
```json
{
    "status": false,
    "message": "Failed to UPDATE data",
    "errors": [
        "Key: 'UpdateAddressRequest.NoTelp' Error:Field validation for 'NoTelp' failed on the 'required' tag"
    ],
    "data": null
}
```

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Failed to UPDATE data",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

---

### Delete Address

**Endpoint:** `DELETE /api/v1/user/alamat/{id}`

**Request Header:**
- `token: <JWT_TOKEN>`

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to DELETE data",
    "errors": null,
    "data": true
}
```

**Response Body (Failed - Record Not Found):**
```json
{
    "status": false,
    "message": "Failed to DELETE data",
    "errors": [
        "record not found"
    ],
    "data": null
}
```

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Failed to DELETE data",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```