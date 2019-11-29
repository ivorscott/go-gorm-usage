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

### Examples

Use the repository's Makefile to run examples.

Commit messages use Angular commit messages enforced by [commitizen-go](https://github.com/lintingzhen/commitizen-go). So make sure to read them for documentation or comments in the example files.

```
make connection
make table
make model
make fields
make increment
make transient
make record
make query
make update
make delete
```

### Debugging

To see the changes for a given an example run:

`pgcli gorm-usage -U postgres`

![pgcli screenshot](./docs/pgcli.png)
