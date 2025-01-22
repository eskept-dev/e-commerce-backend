### Local Development

# 1. Set up database (PostgreSQL)

```bash
docker run \
--name eskept-postgres \
-p 10001:5432 \
-e POSTGRES_USER=eskept \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=eskept \
-d postgres
```

*Note*
- Enable pgcrypto extension: `CREATE EXTENSION IF NOT EXISTS "pgcrypto";`
- Enable uuid-ossp extension: `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

# 2. Set up cache (Redis)

```bash
docker run \
--name eskept-redis \
-p 10002:6379 \
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
