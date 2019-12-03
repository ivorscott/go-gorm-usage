#!make

# Basics

connection:
	@echo "running ./examples/postgres/connection"

ifndef dsn
	@go run ./examples/postgres/connection 
else
	@go run ./examples/postgres/connection -dsn $(dsn)
endif

table:
	@echo "running ./examples/postgres/createTable"
	@go run ./examples/postgres/createTable

model:
	@echo "running ./examples/postgres/createModel"
	@go run ./examples/postgres/createModel

record:
	@echo "running ./examples/postgres/createRecord"
	@go run ./examples/postgres/createRecord

fields:
	@echo "running ./examples/postgres/createModelFields"
	@go run ./examples/postgres/createModelFields

increment:
	@echo "running ./examples/postgres/increment"
	@go run ./examples/postgres/increment

transient:
	@echo "running ./examples/postgres/tempField"
	@go run ./examples/postgres/tempField

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

unique:
	@echo "running ./examples/postgres/unique"
	@go run ./examples/postgres/unique

index:
	@echo "running ./examples/postgres/index"
	@go run ./examples/postgres/index

required:
	@echo "running ./examples/postgres/required"
	@go run ./examples/postgres/required

default:
	@echo "running ./examples/postgres/default"
	@go run ./examples/postgres/default

primary:
	@echo "running ./examples/postgres/primary"
	@go run ./examples/postgres/primary

rename_column:
	@echo "running ./examples/postgres/renameColumn"
	@go run ./examples/postgres/renameColumn

embedding:
	@echo "running ./examples/postgres/embedding"
	@go run ./examples/postgres/embedding

index_call:
	@echo "running ./examples/postgres/indexCall"
	@go run ./examples/postgres/indexCall

# Relationships

one_to_one:
	@echo "running ./examples/postgres/_relationships/oneToOne"
	@go run ./examples/postgres/_relationships/oneToOne

foreign_key:
	@echo "running ./examples/postgres/_relationships/foreignKey"
	@go run ./examples/postgres/_relationships/foreignKey

one_to_many:
	@echo "running ./examples/postgres/_relationships/oneToMany"
	@go run ./examples/postgres/_relationships/oneToMany

many_to_many:
	@echo "running ./examples/postgres/_relationships/manyToMany"
	@go run ./examples/postgres/_relationships/manyToMany

polymorphic:
	@echo "running ./examples/postgres/_relationships/polymorphism"
	@go run ./examples/postgres/_relationships/polymorphism