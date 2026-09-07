Kontrak API

Endpoint: http://localhost:3000/api/v1/students (semua menggunakan url yang sama)


// Setup Create Data 1

Method: POST
Body: {
    "NIM": 434241001,
    "Name": "udin",
    "Grade": 3.5,
    "IsActive": true
}
Return Status: 201 Created
Response:
{
    "Success": true,
    "Message": "student berhasil dibuat",
    "Data": {
        "ID": 1,
        "NIM": 434241001,
        "Name": "udin",
        "Grade": 3.5,
        "IsActive": true
    },
    "Meta": null,
    "Errors": null
}


// Setup Create Data 2

Method: POST
Body: {
    "NIM": 434241002,
    "Name": "viki",
    "Grade": 4.0,
    "IsActive": true
}
Return Status: 201 Created
Response:
{
    "Success": true,
    "Message": "student berhasil dibuat",
    "Data": {
        "ID": 2,
        "NIM": 434241002,
        "Name": "viki",
        "Grade": 4.0,
        "IsActive": true
    },
    "Meta": null,
    "Errors": null
}


// Setup Create Data 3

Method: POST
Body: {
    "NIM": 434241003,
    "Name": "rizkimok",
    "Grade": 2.75,
    "IsActive": false
}
Return Status: 201 Created
Response:
{
    "Success": true,
    "Message": "student berhasil dibuat",
    "Data": {
        "ID": 3,
        "NIM": 434241003,
        "Name": "rizkimok",
        "Grade": 2.75,
        "IsActive": false
    },
    "Meta": null,
    "Errors": null
}


// GET list dengan pagination

Method: GET
URL: /api/v1/students?page=1&limit=10&sort=id&order=asc
Return Status: 200 OK
Response:
{
    "Success": true,
    "Message": "daftar student berhasil diambil",
    "Data": [
        {
            "ID": 1,
            "NIM": 434241001,
            "Name": "udin",
            "Grade": 3.5,
            "IsActive": true
        },
        {
            "ID": 2,
            "NIM": 434241002,
            "Name": "viki",
            "Grade": 4.0,
            "IsActive": true
        }
    ],
    "Meta": {
        "Page": 1,
        "Limit": 10,
        "Total": 2,
        "TotalPages": 1
    },
    "Errors": null
}


// GET by ID berhasil

Method: GET
URL: /api/v1/students/1
Return Status: 200 OK
Response:
{
    "Success": true,
    "Message": "student ditemukan",
    "Data": {
        "ID": 1,
        "NIM": 434241001,
        "Name": "udin",
        "Grade": 3.5,
        "IsActive": true
    },
    "Meta": null,
    "Errors": null
}


// GET by ID tidak ditemukan

Method: GET
URL: /api/v1/students/999
Return Status: 404 Not Found
Response:
{
    "Success": false,
    "Message": "student tidak ditemukan",
    "Data": null,
    "Meta": null,
    "Errors": null
}


// GET invalid (bukan angka)

Method: GET
URL: /api/v1/students/abc
Return Status: 400 Bad Request
Response:
{
    "Success": false,
    "Message": "id harus berupa angka positif",
    "Data": null,
    "Meta": null,
    "Errors": null
}


// POST tanpa Content-Type

Method: POST
Body: {
    "NIM": 434241099,
    "Name": "fedo",
    "Grade": 3.0,
    "IsActive": true
}
Return Status: 415 Unsupported Media Type
Response:
{
    "Success": false,
    "Message": "Content-Type harus application/json",
    "Data": null,
    "Meta": null,
    "Errors": null
}


// POST body bukan JSON

Method: POST
Body: bukan json valid
Return Status: 400 Bad Request
Response:
{
    "Success": false,
    "Message": "body harus berupa JSON yang valid",
    "Data": null,
    "Meta": null,
    "Errors": null
}


// POST NIM kembar

Method: POST
Body: {
    "NIM": 434241001,
    "Name": "vitosuki",
    "Grade": 3.8,
    "IsActive": true
}
Return Status: 422 Unprocessable Entity
Response:
{
    "Success": false,
    "Message": "validasi gagal",
    "Data": null,
    "Meta": null,
    "Errors": {
        "nim": "NIM sudah dipakai"
    }
}


// PUT ganti total

Method: PUT
URL: /api/v1/students/1
Body: {
    "NIM": 434241010,
    "Name": "fedo",
    "Grade": 3.9,
    "IsActive": false
}
Return Status: 200 OK
Response:
{
    "Success": true,
    "Message": "student berhasil diganti seluruhnya",
    "Data": {
        "ID": 1,
        "NIM": 434241010,
        "Name": "fedo",
        "Grade": 3.9,
        "IsActive": false
    },
    "Meta": null,
    "Errors": null
}


// PATCH ubah sebagian

Method: PATCH
URL: /api/v1/students/2
Body: {
    "Grade": 3.25
}
Return Status: 200 OK
Response:
{
    "Success": true,
    "Message": "student berhasil diperbarui sebagian",
    "Data": {
        "ID": 2,
        "NIM": 434241002,
        "Name": "viki",
        "Grade": 3.25,
        "IsActive": true
    },
    "Meta": null,
    "Errors": null
}


// DELETE

Method: DELETE
URL: /api/v1/students/3
Return Status: 204 No Content
Response: (tidak ada body)
