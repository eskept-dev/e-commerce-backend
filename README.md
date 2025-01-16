### Local Development

# 1. Set up database (PostgreSQL)

```bash
docker run \
--name eskept-postgres \
-p 5432:10001 \
-e POSTGRES_USER=eskept \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=eskept \
-d postgres
```

# 2. Set up cache (Redis)

```bash
docker run \
--name eskept-redis \
-p 6379:10002 \
-d redis
```

# 3. Run the application

# 3.1. Set up environment variables

```bash
cp .env.example .env
```

# 3.2. Run the application

```bash
go run main.go
```
