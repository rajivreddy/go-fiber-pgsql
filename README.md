# Sample Go REST API

### Frameworks
1. gorm
2. fiber/v2

# Start your service

```bash
docker-compose pull
docker-compose build
docker-compose up
+] Running 2/2
 ✔ Container go-fiber-pgsql-postgres-1  Created                                                                                                                                                                       0.0s
 ✔ Container go-fiber-pgsql-api-1       Recreated                                                                                                                                                                     0.0s
Attaching to api-1, postgres-1
```

you can access endpints on http://localhost:3000/api/books
