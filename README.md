<h1 align="left"><strong>Repository of canopas website includes intuitive UI designs</strong></h1>

Source code repository of [Canopas Software Website](https://canopas.com).

We intend to keep this open source. Plan is to keep the repository up to date with latest technologies and best practices.

## Requirements

- Node18
- Go 1.17

### Note :

- We have used SSR(Server Side Rendering) in website to improve performance. Server will render UI in `<!--app-html-->` part of `vue-frontend/index.html`. Due to this, reformating of `index.html` can cause UI breaking.

## To run front-end

- Install dependencies

  ```
  cd vue-frontend && yarn install
  ```

- Run vue server

  ```
  yarn dev
  ```

- Compiles and minifies for production

  ```
  yarn build
  ```

- Server side rendering

  ```
  yarn build
  ```

  ```
  yarn serve
  ```

You can access the page by pointing a web browser at http://localhost:8080.

## To run backend

- Start Go Server

  ```
  go run main.go
  ```

  If any dependency error is there, run

  ```
  go mod tidy
  ```

  You can access the APIs at `http://localhost:8080`.

### We are using go modules for go package dependency.

- Initialize go modules

  ```
  go mod init
  ```

- By running above command go.mod file created in project directory. It includes all packages required in project.

- Add package in project

  ```
  go get package_name
  ```

- Update package

  ```
  go get -u package_name
  ```

- sync dependency with go.mod

  ```
  go mod tidy
  ```

### For writing unit tests in golang :

- cd to package for which, you want to write test.

  ```
  cd package_name
  ```

- Create file with suffix test like packageName_test.go.

- Create function with prefix Test like TestFunctionName.

- Run test using

  ```
  go test .
  ```

  For cleaning test cache

  ```
  go clean -testcache
  ```

  [Here](https://github.com/canopas/canopas-website/blob/master/api-doc.md) is APIs reference used in the website.

## To enable pre-commit hook

```
git config core.hooksPath .githooks
```

## Dependencies

Following are dependencies used by the project

#### [VueJs](https://github.com/vuejs/core)

- It is used to build web interfaces and one-page applications.

#### [Vite](https://github.com/vitejs/vite)

- It is frontend build tool that significantly improves the frontend development experience. We have used it for server side rendering.

#### [Tailwind CSS](https://tailwindcss.com/)

- A utility-first CSS framework packed with classes like flex, pt-4, my-4 and text-center that can be composed to build any design, directly in your markup.

#### [Vue collapse transition](https://github.com/ivanvermeyen/vue-collapse-transition)

- It is used for collapses elements horizontally or vertically.

#### [FontAwesome](https://github.com/FortAwesome/Font-Awesome)

- Used for icons.

## LICENSE

Canopas is released under the [GNU V3](https://github.com/canopas/canopas-website/blob/master/LICENSE.md).
