# go-orm-challange

## sql migrate setup
```bash
# https://github.com/rubenv/sql-migrate
go get -v github.com/rubenv/sql-migrate/...
```

```yaml
development:
  dialect: mysql
  dir: db/gorm/migrations
  datasource: root:password@tcp(localhost)/gormdb?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true
  table: migrations
```

```sql
-- db/gorm/migrations/20200402185223-create_users.sql
-- +migrate Up
CREATE TABLE users (
                       id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       uid BIGINT NOT NULL,
                       created_at DATETIME  DEFAULT CURRENT_TIMESTAMP NOT NULL ,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL,
                       UNIQUE unique_users_on_uid (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE users;
```

```bash
sql-migrate up -config=./db/gorm/dbconfig.yml
```

## sqlboiler
```
# https://github.com/volatiletech/sqlboiler
# https://zenn.dev/sagae/articles/c6b2e460201d31
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
# go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
sql-migrate up -config=./db/sqlboiler/dbconfig.yml
sqlboiler mysql
```

```toml:sqlboiler.toml
pkgname="dao"
output="internal/infrastructure/sqlboiler/dao"
add-global-variants=true
add-panic-variants=true
[mysql]
  dbname="sqlboilerdb"
  host="127.0.0.1"
  port=3394
  user="root"
  pass="password"
  sslmode="false"
  skipSQLCmd="false"
  blacklist = ["migrations"]
```

```bash
sqlboiler mysql
go mod tidy
```
