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

# CRUD

create_task:
	@echo "running ./examples/postgres/_crud/create"
	@go run ./examples/postgres/_crud/create

update_task:
	@echo "running ./examples/postgres/_crud/update"
	@go run ./examples/postgres/_crud/update

update_columns:
	@echo "running ./examples/postgres/_crud/updateColumns"
	@go run ./examples/postgres/_crud/updateColumns

batch_update:
	@echo "running ./examples/postgres/_crud/batchUpdate"
	@go run ./examples/postgres/_crud/batchUpdate

delete_task:
	@echo "running ./examples/postgres/_crud/delete"
	@go run ./examples/postgres/_crud/delete

transactions:
	@echo "running ./examples/postgres/_crud/transactions"
	@go run ./examples/postgres/_crud/transactions

# Querying the database

find:
	@echo "running ./examples/postgres/_querying/find"
	@go run ./examples/postgres/_querying/find

where_clauses:
	@echo "running ./examples/postgres/_querying/whereClauses"
	@go run ./examples/postgres/_querying/whereClauses

eager_loading:
	@echo "running ./examples/postgres/_querying/eagerLoading"
	@go run ./examples/postgres/_querying/eagerLoading

limits_orderby_offset:
	@echo "running ./examples/postgres/_querying/limitsOrderbyOffset"
	@go run ./examples/postgres/_querying/limitsOrderbyOffset

subsets:
	@echo "running ./examples/postgres/_querying/subsets"
	@go run ./examples/postgres/_querying/subsets

attribute_assignments:
	@echo "running ./examples/postgres/_querying/attributeAssignments"
	@go run ./examples/postgres/_querying/attributeAssignments

joins:
	@echo "running ./examples/postgres/_querying/joins"
	@go run ./examples/postgres/_querying/joins

rows:
	@echo "running ./examples/postgres/_querying/rows"
	@go run ./examples/postgres/_querying/rows

grouping:
	@echo "running ./examples/postgres/_querying/grouping"
	@go run ./examples/postgres/_querying/grouping

raw_sql:
	@echo "running ./examples/postgres/_querying/rawSql"
	@go run ./examples/postgres/_querying/rawSql

# Modifying

add_column:
	@echo "running ./examples/postgres/_modifying/addColumn"
	@go run ./examples/postgres/_modifying/addColumn

add_index:
	@echo "running ./examples/postgres/_modifying/addIndex"
	@go run ./examples/postgres/_modifying/addIndex

drop_column:
	@echo "running ./examples/postgres/_modifying/dropColumn"
	@go run ./examples/postgres/_modifying/dropColumn

modify_column:
	@echo "running ./examples/postgres/_query_modifyinging/modifyColumn"
	@go run ./examples/postgres/_modifying/modifyColumn

remove_entity:
	@echo "running ./examples/postgres/_modifying/removeEntity"
	@go run ./examples/postgres/_modifying/removeEntity

# Advanced

callbacks:
	@echo "running ./examples/postgres/_advanced/callbacks"
	@go run ./examples/postgres/_advanced/callbacks

scopes:
	@echo "running ./examples/postgres/_advanced/scopes"
	@go run ./examples/postgres/_advanced/scopes

logging:
	@echo "running ./examples/postgres/_advanced/logging"
	@go run ./examples/postgres/_advanced/logging