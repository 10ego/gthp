# Build steps

1. templ generate
2. sqlc generate
3. goose up
4. npm run build-css

# Run steps

To migrate/initialize database:
```
go run cmd/schema/main.go
```

To run the app:
```
go run cmd/server/main.go
```
