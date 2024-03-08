package schema

type Application struct {
  CompanyName string  `bson:"company_name"` 
  Role string         `bson:"role"`
  ApplicationURL string `bson:"application_url"`
  Status string `bson:"status"`
}

