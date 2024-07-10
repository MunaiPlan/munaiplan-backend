data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./infrastructure/database/postgres/models",  // Adjust path to your models
    "--dialect", "postgres",  // Replace with your database dialect
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable&TimeZone=UTC"  // Direct connection string
  migration {
    dir = "file://../../infrastructure/database/postgres/migrations"  // Adjust path to your migrations
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
