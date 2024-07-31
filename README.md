# todo_list

# backend
- Need go (i'm running on 1.22.4), mysql 8.0
- Clone config.tmp.yml to config.yml and set your own env
- Run:
  - go mod tidy
  - go run cmd/main.go server
  - 2 apis:
    - GET /api/v1/task get list tasks
    - POST /api/v1/task upsert tasks

# front end
- You can run via python -m http.server or live server/nginx, any thing you want
- It is not complete yet, you can use postman to fully test back end apis
