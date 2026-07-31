# API Specification

## Base URL

```
http://localhost:8080
```

---

# Authentication

Protected APIs require a JWT access token.

Example Header

```http
Authorization: Bearer <jwt_token>
```

---

# 1. Create Staff

## Endpoint

```http
POST /staff/create
```

## Description

Create a new hospital staff account.

## Request Body

| Field       | Type   | Required | Description    |
| ----------- | ------ | -------- | -------------- |
| username    | string | Yes      | Staff username |
| password    | string | Yes      | Staff password |
| hospital_id | uint   | Yes      | Hospital ID    |

### Example Request

```json
{
  "username": "john",
  "password": "123456",
  "hospital_id": 1
}
```

## Success Response

**201 Created**

```json
{
  "message": "staff created successfully",
  "staff": {
    "id": 1,
    "username": "john",
    "hospital_id": 1
  }
}
```

## Error Responses

### 400 Bad Request

```json
{
  "message": "username, password and hospital_id are required"
}
```

### 409 Conflict

```json
{
  "message": "username already exists in this hospital"
}
```

---

# 2. Staff Login

## Endpoint

```http
POST /staff/login
```

## Description

Authenticate a staff member and generate a JWT token.

## Request Body

| Field       | Type   | Required |
| ----------- | ------ | -------- |
| username    | string | Yes      |
| password    | string | Yes      |
| hospital_id | uint   | Yes      |

### Example Request

```json
{
  "username": "john",
  "password": "123456",
  "hospital_id": 1
}
```

## Success Response

**200 OK**

```json
{
  "message": "login successful",
  "token": "<jwt_token>",
  "staff": {
    "id": 1,
    "username": "john",
    "hospital_id": 1,
    "hospital": "Hospital A"
  }
}
```

## Error Responses

### 400 Bad Request

```json
{
  "message": "username, password and hospital_id are required"
}
```

### 401 Unauthorized

```json
{
  "message": "invalid username, password or hospital"
}
```

---

# 3. Search Patient

## Endpoint

```http
GET /api/patient/search
```

## Description

Search patients belonging to the authenticated staff's hospital.

Authentication is required.

## Query Parameters

| Parameter     | Type   | Required |
| ------------- | ------ | -------- |
| national_id   | string | No       |
| passport_id   | string | No       |
| first_name    | string | No       |
| middle_name   | string | No       |
| last_name     | string | No       |
| date_of_birth | string | No       |
| phone_number  | string | No       |
| email         | string | No       |

### Example

```http
GET /api/patient/search?national_id=9876543210987
```

## Success Response

**200 OK**

```json
{
  "message": "patients found",
  "data": [
    {
      "ID": 2,
      "HospitalID": 2,
      "Hospital": {
        "ID": 2,
        "Name": "Hospital B",
        "CreatedAt": "2026-07-31T13:58:49.438314Z",
        "UpdatedAt": "2026-07-31T13:58:49.438314Z"
      },
      "PatientHN": "HN000002",
      "NationalID": "9876543210987",
      "PassportID": "BB7654321",
      "FirstNameTH": "สุดา",
      "MiddleNameTH": "",
      "LastNameTH": "สุขใจ",
      "FirstNameEN": "Suda",
      "MiddleNameEN": "",
      "LastNameEN": "Sukjai",
      "DateOfBirth": "1995-05-10T00:00:00Z",
      "PhoneNumber": "0898765432",
      "Email": "suda@example.com",
      "Gender": "F",
      "CreatedAt": "2026-07-31T14:02:04.265206Z",
      "UpdatedAt": "2026-07-31T14:02:04.265206Z"
    }
  ]
}
```

## Error Responses

### 401 Unauthorized

```json
{
  "message": "missing authorization header"
}
```

### 500 Internal Server Error

Returned when an unexpected server-side error occurs, such as a database or external API failure.

Example:

```json
{
  "message": "<error message>"
}
```

---

# Authentication Test Endpoint

## Endpoint

```http
GET /api/profile
```

## Description

Simple protected endpoint used to verify JWT authentication.

## Success Response

```json
{
  "message": "protected route success"
}
```

---

# Design Notes

- JWT is used for authentication.
- Passwords are securely hashed before storage.
- Only authenticated staff members can access protected APIs.
- Staff members can only search patients within their assigned hospital.
- Patient records are retrieved from the local database first. If no matching patient is found and a national ID or passport ID is provided, the system retrieves data from Hospital A API and stores it in PostgreSQL.
- The Staff APIs use `hospital_id` instead of a hospital name to maintain referential integrity and normalize the database schema through foreign key relationships.
