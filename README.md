# 🐾 Animal API (Go Interview Project)

## Overview
Build a minimal Go web server that allows a user to view a list of animals from a JSON file using a single API endpoint. The endpoint must be protected with a fixed bearer token.

## ✅ Requirements

#### 🔐 Authentication
- The API must require a static Bearer token in the `Authorization` header.
- Example token: `Bearer secret123`
- If missing or incorrect, return:
  - 401 Unauthorized for missing or invalid token

#### Endpoint
| Method | Route | Description |
|----------|----------|----------|
| GET |	`/animals` |	Return list of animals |
| GET |	`/animals/:id` |	Return one animal (Bonus) |
| PATCH |	`/animals/:id` |	Update one animal (Bonus) |

- Returns a JSON array of animal objects.
- Each object includes: `id`, `name`, and `species`.

#### Sample Data (`animal.json`)
```json
[
  { "id": 1, "name": "Otter", "species": "Lutra lutra" },
  { "id": 2, "name": "Ferret", "species": "Mustela putorius furo" }
]
```
#### Bonus
-  Simple test
-  Good Practices / Structure


## Example Request
```bash
curl -H "Authorization: Bearer secret123" http://localhost:8080/animals
```

Response:
```json
[
  { "id": 1, "name": "Otter", "species": "Lutra lutra" },
  { "id": 2, "name": "Ferret", "species": "Mustela putorius furo" }
]
```
