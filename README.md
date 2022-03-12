# canopas-website

Source code repository of [Canopas Software LLP Website](https://canopas.com).

We intend to keep this open source. Plan is to keep the repository up to date with latest technologies and
best practices.

## Requirements

- Node16
- Go 1.17

## To run front-end

- Install dependencies

  ```
  cd vue-frontend && npm install
  ```

- Run vue server
  ```
  npm run serve
  ```

You can access the page by pointing a web browser at http://localhost:8080.

## To run backend

```
go run main.go
```

You can access the APIs at http://localhost:8080.

## To enable pre-commit hook

```
git config core.hooksPath .githooks
```

## LICENSE

Canopas is released under the [GNU V3](https://github.com/canopas/canopas-website/blob/master/LICENSE.md).
