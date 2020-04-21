# MIGRATION GUIDES

### Some Notes

- Migration in go is different than python (django), ruby (ruby on rails) or php (laravel)

- To run database migration in go, we using both `gorm` and `migrate`. 

- `gorm` is only migrate new model or new column, `gorm` cannot modify existing column to protect your data

- `migrate` cli can help us run migrate via cli, but we need write plain SQL syntax to migration

- So, we'll use `gorm` to migrate new table/column and use `migrate` to modify on existing data


### Prerequisites

- `gorm` : we already have in our source code

- `migrate`: https://github.com/golang-migrate/migrate - Download and install on your machine before we go to the next step


### How to migrate

1. `gorm`: We using we just need define model, then we invoke function db.Automigrate() on the main.go. Gorm will run migration automatically. And don't need do anything. But, please note that, `gorm` only migrate new table/column and ignore existing columns

2. `migrate`: We use `migate` cli to modify existing table/column. 

- Step 1: Create new migration file:  `migrate create -ext .sql -dir db/migrations <name of migration>` . 2 file: `up` and `down` migration will be create in folder `db/migrations`

- Step 2: Edit `up` file and add SQL statement like this:

```sql 
    CREATE TABLE IF NOT EXISTS test (
      firstname VARCHAR(16)
    );
```

- step 3: Edit `down` file and SQL statement like this:

```sql 
    DROP TABLE IF EXISTS test;
```

- step 4: There're 2 way to make migration: using cli or using code. 
* Using cli: 
```sh 
migrate -source file://db/migrations -database "mysql://pave:Pave2020@tcp(localhost:3306)/client_api_dev" up <version>
```

* Using code: Currently, we already use the `db.Migrate()`, so, the migration will be automatically apply and we don't need do anything.

* To get help, please ru: `migrate --help`


Happy Coding! 👨‍💻


