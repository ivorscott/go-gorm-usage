#!make

table:
	@echo "running ./examples/postgres/createTable"
	@go run ./examples/postgres/createTable

record:
	@echo "running ./examples/postgres/createRecord"
	@go run ./examples/postgres/createRecord

query:
	@echo "running ./examples/postgres/query"

    ifndef query
		  @go run ./examples/postgres/query 
    else
		  @go run ./examples/postgres/query -query $(query)
    endif

update:
	@echo "running ./examples/postgres/update"
	@go run ./examples/postgres/update 

delete:
	@echo "running ./examples/postgres/delete"
	@go run ./examples/postgres/delete 

connection:
	@echo "running ./examples/postgres/connection"

    ifndef dsn
		  @go run ./examples/postgres/connection 
    else
		  @go run ./examples/postgres/connection -dsn $(dsn)
    endif
