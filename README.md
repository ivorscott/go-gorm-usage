# go-gorm-usage

### Requirements

1. [Install postgres](https://www.postgresql.org/download/)
2. Create a user named "postgres" (no password required)

   ```
   createuser postgres
   ```

3. Create a database named "gorm-usage"

   ```
   createdb gorm-usage
   ```

4. Install [pgcli](https://www.pgcli.com/) for better tooling.

## Examples

Use the repository's Makefile to run examples.

Commit messages use Angular commit messages enforced by [commitizen-go](https://github.com/lintingzhen/commitizen-go). So make sure to read them for documentation.

To see the changes for a given an example run:

`pgcli gorm-usage -U postgres`

![pgcli screenshot](./docs/pgcli.png)

### Basics

<details>
  <summary>See commands</summary>

```makefile
make connection # connect to database
make table # create database table
make model # create gorm.Model
make record # create database record
make fields # create table fields with specified type and size
make increment # create auto incrementing field
make transient # create temporary field
make query # query "first" or "last" name (e.g., make query query=last)
make update # perform record update
make delete # delete record
make unique # create unique field
make index # create unique index
make required # disable null fields
make default # provide default field value
make primary # create primary key
make rename_column # rename column
make embedding # embedd child objects
make index_call # call index and remove index functions
```

<img src="./docs/tags.png" height=300/>

</details>

### Relationship

<details>
  <summary>See commands</summary>

```

```

</details>

### CRUD

<details>
  <summary>See commands</summary>

```

```

</details>

### Querying the database

<details>
  <summary>See commands</summary>

```

```

</details>

### Modifing Schemas

<details>
  <summary>See commands</summary>

```

```

</details>

### Advanced Topics

<details>
  <summary>See commands</summary>

```

```

</details>
