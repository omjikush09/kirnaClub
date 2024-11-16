## Description

The service is to process the job. It add perimeter to Images that was given in the job.

## Assumptions

1. Image type will be always jpeg.
2. visit_time will be always UTC in **RFC3339** format .

## Installing

1. Clone from github
2. ```bash
      cp ".env.example" ".env"
   ```
3. Add the environment variables to `.env`
4. ```bash
    go mod download
   ```
5. ```bash
     go run main.go
   ```

## Testing

[![Run In Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/14294787-23da7464-3320-497f-bc1a-b75c82affc99?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D14294787-23da7464-3320-497f-bc1a-b75c82affc99%26entityType%3Dcollection%26workspaceId%3Df54489f0-e8d7-44dc-a9d1-5c91cdb059c3#?env%5Blocal%5D=W3sia2V5IjoidXJsIiwidmFsdWUiOiJodHRwOi8vbG9jYWxob3N0OjMwMDAiLCJlbmFibGVkIjp0cnVlLCJ0eXBlIjoiZGVmYXVsdCJ9LHsia2V5IjoibG9jYWwiLCJ2YWx1ZSI6Imh0dHA6Ly9sb2NhbGhvc3Q6MzAwMCIsInR5cGUiOiJkZWZhdWx0In1d)

Example- create job api-
`api/submit/`

Body--

```
{
   "count":1,
   "visits":[
      {
         "store_id":"RP00001",
         "image_url":[
            "https://www.gstatic.com/webp/gallery/2.jpg",
            "https://www.gstatic.com/webp/gallery/3.jpg"
         ],
         "visit_time": "2024-11-17T12:34:56Z"
      }
   ]
}
```

## Work Environment

Laptop- iMac, OS- MacOS, Text Editor- VS Code

## Libraries

1. [Gorm](https://gorm.io/) - for database ORM
2. [Fiber](https://gofiber.io/) - for Building Server
3. [Postgres Driver](https://pkg.go.dev/gorm.io/driver/postgres@v1.5.9) - For handling postgres database
4. [Cron](https://pkg.go.dev/github.com/robfig/cron/v3@v3.0.1) - for handling cron job
5. [Validator](https://pkg.go.dev/github.com/go-playground/validator)- for validating data

## Improvement

1. Would use queue (AWS SQS) to handle jobs
2. Would add more code abstraction
