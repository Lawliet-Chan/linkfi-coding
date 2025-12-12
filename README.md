# LinkFi Coding

## How to Test

### Start the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

### Test Endpoints

#### POST /process-people

```bash
curl -X POST http://localhost:8080/process-people \
  -H "Content-Type: application/json" \
  -d '[
    { "no": 1, "name": "Alden", "age": 24, "birthday": "1999.12.12" },
    { "no": 2, "name": "Briony", "age": 32, "birthday": "1990.05.10" },
    { "no": 3, "name": "Cedric", "age": 28, "birthday": "1995.08.20" }
  ]'
```

Expected response:

```json
{
  "Alden": { "age": 24, "birthday": "1999.12.12" },
  "Briony": { "age": 32, "birthday": "1990.05.10" },
  "Cedric": { "age": 28, "birthday": "1995.08.20" }
}
```
