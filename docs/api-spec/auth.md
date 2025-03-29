## Auth API Spec

### Register User
**Endpoint:** `POST /api/v1/auth/register`

**Request Body:**
```json
{
    "nama": "John Doe",
    "kata_sandi": "123456",
    "tanggal_lahir": "02/01/2000",
    "pekerjaan": "engineer",
    "email": "johndoe@mail.com",
    "id_provinsi": "11",
    "id_kota": "1101"
}
```

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Failed to POST data",
    "errors": null,
    "data": "Register Succeed"
}
```

**Response Body (Failed):**
```json
{
    "status": false,
    "message": "Failed to POST data",
    "errors": [
        "Key: 'RegisterRequest.NoTelp' Error:Field validation for 'NoTelp' failed on the 'required' tag"
    ],
    "data": null
}
```

---

## Login User
**Endpoint:** `POST /api/v1/auth/login`

**Request Body:**
```json
{
    "no_telp": "08961231232",
    "kata_sandi": "123456"
}
```

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to GET data",
    "errors": null,
    "data": {
        "nama": "John Doe",
        "no_telp": "08961231232",
        "tanggal_lahir": "2000-01-02T00:00:00+07:00",
        "tentang": "",
        "pekerjaan": "engineer",
        "email": "johndoe@mail.com",
        "id_provinsi": {
            "id": "11",
            "name": "ACEH"
        },
        "id_kota": {
            "id": "1101",
            "province_id": "11",
            "name": "KABUPATEN SIMEULUE"
        },
        "is_admin": false,
        "token": "<JWT_TOKEN>"
    }
}
```

**Response Body (Failed):**
```json
{
    "status": false,
    "message": "Failed to POST data",
    "errors": ["No Telp atau kata sandi salah"],
    "data": null
}
```

---