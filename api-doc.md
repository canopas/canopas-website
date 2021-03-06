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
                    "facebook": "https://www.facebook.com/canopassoftware","twitter": "https://twitter.com/canopassoftware" ,"instagram": "https://www.instagram.com/canopassoftware/"
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
                "twitter": "https://twitter.com/canopassoftware",
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
        <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
            <url>
                <loc>http://canopas.com</loc>
                <changefreq>monthly</changefreq>
                <lastmod>2022-03-01T00:00:00.000Z</lastmod>
                <priority>1</priority>
            </url>
            <url>
                <loc>http://canopas.com/jobs</loc>
                <changefreq>monthly</changefreq>
                <lastmod>2022-03-01T00:00:00.000Z</lastmod>
                <priority>1</priority>
            </url>
            <url>
                <loc>http://canopas.com/contact</loc>
                <changefreq>monthly</changefreq>
                <lastmod>2022-03-01T00:00:00.000Z</lastmod>
                <priority>0.9</priority>
            </url>
            <url>
                <loc>https://blog.canopas.com</loc>
                <changefreq>monthly</changefreq>
                <lastmod>2022-03-01T00:00:00.000Z</lastmod>
                <priority>0.8</priority>
            </url>
            <url>
                <loc>http://canopas.com/jobs/ios-developers-a9b45f34-a1a5-419f-b536-b7c290925d6d</loc>
                <changefreq>monthly</changefreq>
                <lastmod>2022-03-01T00:00:00.000Z</lastmod>
                <priority>0.9</priority>
            </url>
            <url>
                <loc>http://canopas.com/jobs/apply/ios-developers-a9b45f34-a1a5-419f-b536-b7c290925d6d</loc>
                <changefreq>monthly</changefreq>
                <lastmod>2022-03-01T00:00:00.000Z</lastmod>
                <priority>0.9</priority>
            </url>
        </urlset>
```
