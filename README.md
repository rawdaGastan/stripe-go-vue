# STRIPE-GO-VUE

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/cd6e18aac6be404ab89ec160b4b36671)](https://www.codacy.com/gh/rawdaGastan/stripe-go-vue/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=rawdaGastan/stripe-go-vue&amp;utm_campaign=Badge_Grade) [![Testing](https://github.com/rawdaGastan/stripe-go-vue/actions/workflows/golint.yml/badge.svg?branch=development)](https://github.com/rawdaGastan/stripe-go-vue/actions/workflows/golint.yml) [![Testing](https://github.com/rawdaGastan/stripe-go-vue/actions/workflows/vuelint.yml/badge.svg?branch=development)](https://github.com/rawdaGastan/stripe-go-vue/actions/workflows/vuelint.yml)

stripe-go-vue aims to help students deploy their projects on Threefold Grid.

## Requirements

- docker compose

## Build

First create `config.json` check [configuration](#configuration)

To build backend and frontend images

```bash
docker compose build
```

## Run

First create `config.json` check [configuration](#configuration)

To run backend and frontend:

```bash
docker compose up
```

### Configuration

Before building or running backend image create `config.json` in `server` dir.

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

## Manual testing

You can test stripe using the following [cards](https://stripe.com/docs/testing#cards)

## Frontend

check frontend [README](frontend/README.md)

## Backend

check backend [README](backend/README.md)
