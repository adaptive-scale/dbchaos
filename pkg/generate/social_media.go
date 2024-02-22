package generate

type FieldType struct {
	Name string
	Type Type
}

type Type string

const (
	TypeName         Type = "name"
	TypeString       Type = "string"
	TypeNumber       Type = "number"
	TypeFloat        Type = "float"
	TypeBoolean      Type = "boolean"
	TypeDate         Type = "date"
	TypeEmail        Type = "email"
	TypePassword     Type = "password"
	TypeCompany      Type = "company"
	TypeHackerPhrase Type = "hackerphrase"
)
