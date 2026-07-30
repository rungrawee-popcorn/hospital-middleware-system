# Hospital Middleware System ER Diagram

## Entity Relationship

```text
+----------------------+
|      hospitals       |
+----------------------+
| PK id                |
| name                 |
| created_at           |
| updated_at           |
+----------------------+
          |
          |
          | 1
          |
          | N
+----------------------+
|       staffs         |
+----------------------+
| PK id                |
| username             |
| password             |
| FK hospital_id       |
| created_at           |
| updated_at           |
+----------------------+



+----------------------+
|      hospitals       |
+----------------------+
          |
          |
          | 1
          |
          | N
+----------------------+
|      patients        |
+----------------------+
| PK id                |
| FK hospital_id       |
| patient_hn           |
| national_id          |
| passport_id          |
| first_name_th        |
| middle_name_th       |
| last_name_th         |
| first_name_en        |
| middle_name_en       |
| last_name_en         |
| date_of_birth        |
| phone_number         |
| email                |
| gender               |
| created_at           |
| updated_at           |
+----------------------+
```

```md
## Relationship Description

### Hospital - Staff

One Hospital can have many Staff members.

Relationship:

Hospital (1) ---- (Many) Staff

### Hospital - Patient

One Hospital can have many Patients.

Relationship:

Hospital (1) ---- (Many) Patient

### Access Control

A staff member can only search patients belonging to the same hospital.

The relationship is validated by:

staffs.hospital_id = patients.hospital_id
```
