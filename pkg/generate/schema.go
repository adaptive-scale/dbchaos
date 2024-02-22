package generate

var socialMediaGenerator = map[SchemaType][]FieldType{
	"_users": []FieldType{
		{"id", TypeString},
		{"username", TypeString},
		{"password", TypePassword},
		{"name", TypeName},
		{"email", TypeString},
		{"phone", TypeString},
		{"address", TypeString},
		{"city", TypeString},
	},
}
