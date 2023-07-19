# Contribution API documentation

## 1. Get github repository stars

#### Method : GET

#### Endpoint : /api/github/stars

#### Description : API for getting github repository stars.

#### Request:

- Headers : none
- Body : none

#### Response :

- If API is not found then,

  - Status Code: 404 not found
  - Data: no response data

- If any server error is there then,

  - Status Code: 500 Internal server error
  - Data: no response data

- If request will success ,
  - Status Code: 200 Ok
  - Headers : none
  - Data :
  ```
      [
          {
              "name": "MultiImagePicker",
              "stargazers_count": 0
          },
          {
              "name": "FBNotifications",
              "stargazers_count": 0
          },
      ]
  ```
