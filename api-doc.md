# Canopas website API documentation

### 1. Get Careers Lists

```
- Method : GET
- Endpoint : /api/careers
- Description : API for get list of job/careers
- Request:
    - Headers : none
    - Body : none

- Response :

    - If API is not found or career data not found then,
    - Status Code: 404 not found
    - Data: no response data

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Data :
        [
            {
                "id": 1,
                "title": "IOS Developer",
                "summary": "IOS Developer Summary",
                "description": "<h2><span style=\"font-size:18px\"></span><span></span><span style=\"color:#3d3d3d;\"><span style=\"font-size:22px;\"><strong>IOS Developer Description</strong></span></span></h2>",
                "button_name": "Apply",
                "qualification": "graduation degree",
                "employment_type": "full_time",
                "base_salary": 15000,
                "experience": "0-3 years",
                "is_active": true,
                "skills": "",
                "total_openings": 1,
                "responsibilities": "",
                "icon_name": "fab fa-icon",
                "unique_id": "ios-developer-123",
                "seo_title": "",
                "seo_description": "",
                "apply_seo_title": "",
                "apply_seo_description": ""
            },
            {
                "id": 2,
                "title": "Android Developer",
                "summary": "Android Developer Summary",
                "description": "<h2><span style=\"font-size:18px\"></span><span></span><span style=\"color:#3d3d3d;\"><span style=\"font-size:22px;\"><strong>Android Developer Description</strong></span></span></h2>",
                "button_name": "Apply",
                "qualification": "graduation degree",
                "employment_type": "full_time",
                "base_salary": 15000,
                "experience": "0-3 years",
                "is_active": true,
                "skills": "",
                "total_openings": 1,
                "responsibilities": "",
                "icon_name": "fab fa-icon",
                "unique_id": "android-developer-123",
                "seo_title": "",
                "seo_description": "",
                "apply_seo_title": "",
                "apply_seo_description": ""
            },
        ]

```

### 2. Get Career By ID

```
- Method : GET
- Endpoint : /api/careers/:uniqueId
- Description : API for get careers by provided id
- Request:
    - Headers : none
    - Body : none

- Response :

    - If API is not found or career data not found then,
    - Status Code: 404 not found
    - Data: no response data

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Data :
        {
            "id": 1,
            "title": "IOS Developer",
            "summary": "IOS Developer Summary",
            "description": "<h2><span style=\"font-size:18px\"></span><span></span><span style=\"color:#3d3d3d;\"><span style=\"font-size:22px;\"><strong>IOS Developer Description</strong></span></span></h2>",
            "button_name": "Apply",
            "qualification": "graduation degree",
            "employment_type": "full_time",
            "base_salary": 15000,
            "experience": "0-3 years",
            "is_active": true,
            "skills": "",
            "total_openings": 1,
            "responsibilities": "",
            "icon_name": "fab fa-icon",
            "unique_id": "ios-developer-123",
            "seo_title": "",
            "seo_description": "",
            "apply_seo_title": "",
            "apply_seo_description": ""
        }

```

### 3. Send career mail

```
- Method : POST
- Endpoint : /api/send-career-mail
- Description : API for sending mail to the company after applying for jobs
- Request:
    - Headers : none
    - Body :
        {
            "job_title": "Web Developer",
            "name": "New Web Developer",
            "email": "developer@canopas.com",
            "phone": "9099999999",
            "place": "surat",
            "references" : "From canopas",
            "message"  : I'm a very good programer,
            "file" : resume.pdf,
        }


- Response :

    - If API is not found,
    - Status Code: 404 not found
    - Data: no response data

    - If request data(input) is not valid then,
    - Status Code: 400 Bad Request
    - Data: {"message" : "Invalid data"}

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Data :
        {
            "job_title": "Web Developer",
            "name": "New Web Developer",
            "email": "developer@gmail.com",
            "phone": "9099999999",
            "place": "surat",
            "references": "From canopas",
            "message": "I'm a very good programer"
        }
```

### 4. Send contact mail

```
- Method : POST
- Endpoint : /api/send-contact-mail
- Description : API for sending mail to the company after submitting contact form
- Request:
    - Headers : none
    - Body :
        {
            "email": "shruti@gmail.com",
            "name": "shruti",
            "designation": "I am individual entrepreneur running my own business.",
            "reason": "I have a rough design for the product I want to develop.",
            "message": "just testing purpose",
            "designation_info": "I am an owner of the business and I run canopas business. For more information about my business, you can visit the following links.",
            "idea": "I have an idea for my business that I want to implement with your help.",
            "social_media_links":
                {
                    "website": "https://canopas.com/",
                    "facebook": "https://www.facebook.com/canopassoftware","twitter": "https://x.com/canopassoftware" ,"instagram": "https://www.instagram.com/canopassoftware/"
                },
            "contact_type": "Chat or Email"

        }


- Response :

    - If API is not found,
    - Status Code: 404 not found
    - Data: no response data

    - If request data(input) is not valid then,
    - Status Code: 400 Bad Request
    - Data: {"message" : "Invalid data"}

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Data :
        {
            "name": "shruti",
            "designation": "I am individual entrepreneur running my own business.",
            "designation_info": "I am an owner of the business and I run canopas business. For more information about my business, you can visit the following links.",
            "social_media_links": {
                "facebook": "https://www.facebook.com/canopassoftware",
                "instagram": "https://www.instagram.com/canopassoftware/",
                "twitter": "https://x.com/canopassoftware",
                "website": "https://canopas.com/"
            },
            "idea": "I have an idea for my business that I want to implement with your help.",
            "reason": "I have a rough design for the product I want to develop.",
            "email": "shruti@gmail.com",
            "message": "just testing purpose",
            "contact_type": "Chat or Email"
}
```

### 5. Get sitemap

```
- Method : GET
- Endpoint : /sitemap
- Description : API for generate sitemap
- Request:
    - Headers : none
    - Params :
        key: "baseUrl"
        value: "http://canopas.com"


- Response :

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
        [
            {
                "loc": "/services",
                "changefreq": "monthly",
                "lastmod": "2024-02-01T00:00:00.000Z",
                "priority": "0.9"
            },
            {
                "loc": "/portfolio",
                "changefreq": "monthly",
                "lastmod": "2024-02-01T00:00:00.000Z",
                "priority": "0.9"
            },
            {
                "loc": "/contributions",
                "changefreq": "monthly",
                "lastmod": "2024-02-01T00:00:00.000Z",
                "priority": "0.9"
            },
            {
                "loc": "/resources",
                "changefreq": "monthly",
                "lastmod": "2024-02-01T00:00:00.000Z",
                "priority": "0.9"
            },
            {
                "loc": "/blog",
                "changefreq": "monthly",
                "lastmod": "2024-02-01T00:00:00.000Z",
                "priority": "0.9"
            },
        ]
```

### 6. Send New Leave Request mail

```
- Method : POST
- Endpoint : /api/leave/new
- Description : API for sending mail to HR for New leave request
- Request:
    - Headers : none
    - Body :
        {
            "name":     "mansi dhameliya",
            "date":     "3 jan 2023",  // OR if you want to add multiple dates, you can add like this : "3 jan 2023 to 5 jan 2023"
            "duration": "first-half" // You can add first-half or second-half or 3 days as you want
            "status":   1,  // you can use 1 for request
            "reason":   "Casual Leave",
            "receiver": "hr@canopas.com",
        }
- Response :
    - If API is not found,
    - Status Code: 404 not found
    - Data: no response data

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Message : Leave request has been sent successfully
```

### 7. Send Update Leave Request mail

```
- Method : POST
- Endpoint : /api/leave/update
- Description : API for sending Rejection/Approval mail of leave request to the employee
- Request:
    - Headers : none
    - Body :
         {
            "name": "mansi dhameliya",
            "date": "3 jan 2023", // OR if you want to add multiple dates, you can add like this : "3 jan 2023 to 5 jan 2023"
            "status": 2 OR 3 , // you can use 2 for approved and 3 for rejected
            "receiver": "mansi@canopas.com",
        }
- Response :
    - If API is not found,
    - Status Code: 404 not found
    - Data: no response data

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Message : Update leave request has been sent successfully
```

### 8. Send Invitation Mail (Notification)

```
- Method : POST
- Endpoint : /api/invitation
- Description : API for sending invitation mail to the Company employee
- Request:
    - Headers : none
    - Body :
         {
            "receiver": "mansi@canopas.com",
            "companyname" : "canopas"
            "spacelink": "https://unity.canopas.com/home",
        }
- Response :
    - If API is not found,
    - Status Code: 404 not found
    - Data: no response data

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Message : Invitation mail has been sent successfully
```

### 8. Send Acceptence Mail (Notification)

```
- Method : POST
- Endpoint : /api/acceptence
- Description : API for sending Accepted mail to the Company HR manager
- Request:
    - Headers : none
    - Body :
         {
            "receiver": "hr@canopas.com",
            "sender": "mansi@canopas.com", // accepter employee mail
        }
- Response :
    - If API is not found,
    - Status Code: 404 not found
    - Data: no response data

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Message : Acceptence mail has been sent successfully
```

### 9. Get Career Page LifeatCanopas Section Images

```
- Method : GET
- Endpoint : /api/lifeimages
- Description : API for get images of lifeat canopas section of career page
- Request:
    - Headers : none
    - Body : none

- Response :

    - If API is not found or image data not found then,
    - Status Code: 404 not found
    - Data: no response data

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Data :
        [
            {
                 "id": 1,
                 "image_urls": "https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-400.webp,https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-800.webp,https://canopas-website.s3.ap-south-1.amazonaws.com/lifeCanopas/1-1600.webp",
                 "index": 0
            },
        ]

```

### 10. Get Career Page Perks and Benefit Section Images

```
- Method : GET
- Endpoint : /api/perksimages
- Description : API for get images of perks and benefit section of career page
- Request:
    - Headers : none
    - Body : none

- Response :

    - If API is not found or image data not found then,
    - Status Code: 404 not found
    - Data: no response data

    - If any server error is there then,
    - Status Code: 500 Internal server error
    - Data: no response data

    - If request will success ,
    - Status Code: 200 Ok
    - Headers : none
    - Data :
        [
            {
                 "id": 1,
                 "image_urls": "https://canopas-website.s3.ap-south-1.amazonaws.com/perksBenefits/1-400.webp,https://canopas-website.s3.ap-south-1.amazonaws.com//perksBenefits/1-800.webp",
                 "index": 0
            },
        ]

```
