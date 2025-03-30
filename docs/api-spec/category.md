## Category API Specification

### Get All Categories

**Endpoint:** `GET /api/v1/category`

**Request Header:**
- `token: <JWT_TOKEN>`

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to GET data",
    "errors": null,
    "data": [
        {
            "id": 1,
            "nama_category": "baju update"
        },
        {
            "id": 2,
            "nama_category": "hapus"
        },
        {
            "id": 3,
            "nama_category": "celana"
        }
    ]
}
```

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Please provide a valid authentication token",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

**Response Body (Failed - Forbidden):**
```json
{
    "status": false,
    "message": "Access denied",
    "errors": [
        "Forbidden"
    ],
    "data": null
}
```

---

### Create Category

**Endpoint:** `POST /api/v1/category`

**Request Header:**
- `token: <JWT_TOKEN>`

**Request Body:**
```json
{
    "nama_category": "sepatu"
}
```

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to CREATE data",
    "errors": null,
    "data": {
        "id": 4,
        "nama_category": "sepatu"
    }
}
```

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Please provide a valid authentication token",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

**Response Body (Failed - Validation Error):**
```json
{
    "status": false,
    "message": "Failed to CREATE data",
    "errors": [
        "nama_category is required"
    ],
    "data": null
}
```

---

### Update Category

**Endpoint:** `PUT /api/v1/category/{id}`

**Request Header:**
- `token: <JWT_TOKEN>`

**Request Body:**
```json
{
    "nama_category": "sepatu olahraga"
}
```

**Response Body (Success):**
```json
{
    "status": true,
    "message": "Succeed to UPDATE data",
    "errors": null,
    "data": {
        "id": 4,
        "nama_category": "sepatu olahraga"
    }
}
```

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Please provide a valid authentication token",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

**Response Body (Failed - Not Found):**
```json
{
    "status": false,
    "message": "Failed to UPDATE data",
    "errors": [
        "record not found"
    ],
    "data": null
}
```

---

### Delete Category

**Endpoint:** `DELETE /api/v1/category/{id}`

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

**Response Body (Failed - Unauthorized):**
```json
{
    "status": false,
    "message": "Please provide a valid authentication token",
    "errors": [
        "Unauthorized"
    ],
    "data": null
}
```

**Response Body (Failed - Not Found):**
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

---