data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "mysql://root:Sunil@513@:3306/practice_db"
  migration {
    dir = "file://migrations" 
    format = golang-migrate // this is optional
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

