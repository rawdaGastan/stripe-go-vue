# Backend stripe go

Go backend server using stripe and sqlite3 DB.

## Requirements

- Go >= 1.18
- make
- docker

## Configuration

Before building or running backend create `config.json` in `backend` dir.

example `config.json`:

```json
{
    "port": ":3000",
    "database": {
        "file": "./database.db"
    },
    "stripe": {
        "publisher": "<publisher-key>",
        "secret": "<secret-key>"
    },
    "version": "v1"
}
```

## Build

```bash
make build
```

## Run

```bash
make run
```

### Build Using Docker

```bash
docker build -t stripe_go .
```

### Run Using Docker

```bash
docker run -p 3000:3000 stripe_go
```
