<p align="center"><a href="https://canopas.com/contact"><img src="./assets/banner.png"></a></p>

<h1 align="left"><strong>Website with intuitive UI designs</strong></h1>

We intend to keep this open source. The plan is to keep the repository up to date with the latest technologies and best practices.

## Showcase

This repository contains the working code of [Canopas website](https://canopas.com).
You can also check out [website](https://canopas.com/) to view a live example of this repository.

### [Services](https://canopas.com/services)

<p align="center"><a href="https://canopas.com/services"><img src="./assets/services.gif" /></a></p>

### [portfolio](https://canopas.com/portfolio)

<p align="center"><a href="https://canopas.com/services"><img src="./assets/portfolio.gif" /></a></p>

---

# Table of contents

- [Key features](https://github.com/canopas/canopas-website#key-features)
- [Canopas website -- Backend](https://github.com/canopas/canopas-website#backend)
  - [Requirements](https://github.com/canopas/canopas-website#requirements)
  - [Setup](https://github.com/canopas/canopas-website#setup)
- [Canopas website -- Frontend](https://github.com/canopas/canopas-website#frontend)
  - [Requirements](https://github.com/canopas/canopas-website#requirements-1)
  - [Setup](https://github.com/canopas/canopas-website#setup-1)
- [Formatting and Linting](https://github.com/canopas/canopas-website#formatting-and-linting)
- [Dependencies](https://github.com/canopas/canopas-website#dependencies)
  - [Backend](https://github.com/canopas/canopas-website#requirements)
- [Licence](https://github.com/canopas/canopas-website#licence)

---

# Key Features

1. **Nuxt 3 as Frontend:** Canopas website is built using Nuxt 3, a progressive JavaScript framework that offers a flexible and efficient way to build user interfaces. Nuxt is renowned for its simplicity and ease of integration, making it an ideal choice for creating dynamic and interactive web applications.

2. **Go as Backend:** Backend is developed using Golang. Go is a statically typed, compiled language designed for simplicity, efficiency, and concurrency. With its focus on performance and readability, Go offers a robust foundation for building scalable applications.

3. **Server Side Rendering:** We have used Nuxt's different rendering modes for the website. For static pages, it is SSG and for dynamic pages, we used SSR/ISR.

4. **SEO-friendly URLs and Metadata:** We prioritize SEO best practices to ensure that the website receives the visibility it deserves in search engines. From meta tags to URL structure and dynamic sitemaps, canopas website is equipped with all the essential elements for optimal SEO.

5. **Responsive and Mobile-friendly Design:** We understand the importance of a mobile-friendly design in modern SEO. Canopas' website is fully responsive and adapts seamlessly to various screen sizes, providing a positive user experience across devices.

6. **CI/CD and Deployment:** With CI in place, every code change is automatically tested and integrated into the main codebase. This ensures that the main branch always remains stable, reducing the chances of bugs and allowing for rapid deployment. This repository contains AWS Lambda, S3, and CloudFront stack for smooth deployment using a cloudformation template.

7. **Code formatting and linting:** Clean, readable, and consistent code is the foundation of any successful project. This website follows strict code formatting and linting rules, which are enforced through automated tools. This ensures that the codebase remains maintainable and adheres to industry best practices.

8. **reCAPTCHA Integration:** Security is an important aspect, especially when it comes to user-generated content. Canopas website integrates reCAPTCHA, a widely trusted CAPTCHA service, to protect your site from spam and abuse while maintaining a user-friendly experience.

---

# Canopas website -- Backend

## Requirements

- Go 1.21

## Setup

- Start to Go Server

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

- By running the above command go.mod file created in the project directory. It includes all packages required in the project.

- Add package to project

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

### For writing unit tests in Golang:

- cd to package for which, you want to write test.

  ```
  cd package_name
  ```

- Create a file with a suffix test like packageName_test.go.

- Create a function with the prefix Test like TestFunctionName.

- Run a test using

  ```
  go test .
  ```

  For cleaning the test cache

  ```
  go clean -testcache
  ```

  [Here](https://github.com/canopas/canopas-website/blob/master/api-doc.md) is an APIs reference used in the website.

### For generating sitemap :

- Add your pages to templates/path.txt

# Canopas website -- Frontend

## Requirements

- Node22

## Setup

### Install dependencies

```
cd nuxt-frontend && yarn install
```

### Start website in dev mode using,

```
yarn dev
```

### Compiles and minifies for production

```
yarn build
```

You can access the page by pointing a web browser at http://localhost:8080.

# Formatting and Linting

The pre-commit hook will automatically lint and format your code before committing.

## To enable pre-commit hook

```
git config core.hooksPath .githooks
```

# Dependencies

## Backend

### [gin](https://github.com/gin-gonic/gin)

- The core package of the Gin web framework, providing routing, middleware, and other essential features for building web applications.

### [sqlx](https://github.com/jmoiron/sqlx)

- A library that provides a set of extensions on top of the standard Go database/sql package. It enhances database access with features like named parameters and better scanning into structs.

### [logrus](https://github.com/sirupsen/logrus)

- A structured logger for Go. It provides a flexible way to log messages with different severity levels and customizable output formats.

### [testify](https://github.com/stretchr/testify)

- A popular testing toolkit for Go. It provides various assertion functions and testing utilities to enhance the readability and maintainability of your test code.

### [cors](https://github.com/gin-contrib/cors)

- A middleware for Gin that provides Cross-Origin Resource Sharing (CORS) support, allowing controlled access to resources hosted on different domains.

### [gzip](https://github.com/gin-contrib/gzip)

- A middleware for Gin that enables Gzip compression of HTTP responses, reducing data transfer size and improving performance.

## Frontend

### [VueJs](https://github.com/vuejs/core)

- It is used to build web interfaces and one-page applications.

### [Vite](https://github.com/vitejs/vite)

- It is frontend build tool that significantly improves the frontend development experience. We have used it for server side rendering.

### [Tailwind CSS](https://tailwindcss.com/)

- A utility-first CSS framework packed with classes like flex, pt-4, my-4 and text-center that can be composed to build any design, directly in your markup.

### [Vue collapse transition](https://github.com/ivanvermeyen/vue-collapse-transition)

- It is used for collapsing elements horizontally or vertically.

### [FontAwesome](https://github.com/FortAwesome/Font-Awesome)

- Used for icons.

### [Swiper](https://github.com/nolimits4web/swiper)

- Used for Animation.

## Contribution

The Canopas team enthusiastically welcomes contributions and project participation! There are a bunch of things you can do if you want to contribute! The [Contributor Guide](CONTRIBUTING.md) has all the information you need for everything from reporting bugs to contributing entire new features. Please don't hesitate to jump in if you'd like to, or even ask us questions if something isn't clear.

## Credits

This repository is owned and maintained by the [Canopas team](https://canopas.com/). Please let us know if you are interested in building web apps or designing products. We'd love to hear from you!

<a href="https://canopas.com/contact"><img src="./assets/cta.png" width=300></a>


## LICENSE

The Canopas website has been released under the [GNU V3](https://github.com/canopas/canopas-website/blob/master/LICENSE.md).
