Applying migrations:

```bash
$ ~ go install github.com/pressly/goose/v3/cmd/goose@latest
$ ~ cd migrations
$ ~ GOOSE_DRIVER=postgres GOOSE_DBSTRING="host=localhost user=postgres password=password dbname=list" goose up
```
### Done:
- JWT for authentication (with Middleware)
- Basic CRUD functionality
- Db migrations
- User CRUD functionality

### Todo:
- CSRF
- Generate OpenAPI spec
- Automate build and deployment using Github Actions