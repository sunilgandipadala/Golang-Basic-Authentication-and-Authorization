schema "practice_db" {}
table "employees" {
  schema = schema.practice_db
  column "id" {
    null = false
    type = varchar(191)      
  }
  column "name" {
    null = true
    type = longtext
  }
  column "gender" {
    null = true
    type = longtext
  }
  column "role" {
    null = true
    type = longtext
  }
  column "district" {
    null = true
    type = longtext
  }
  column "pincode" {
    null = true
    type = bigint
  }
  column "age" {
    null = true
    type = int
  }
  primary_key {
    columns = [column.id]
  }
  index "id" {
    unique  = true
    columns = [column.id]
  }
}