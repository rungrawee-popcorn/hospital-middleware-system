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
