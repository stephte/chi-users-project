package enums


type Enum interface {
	String()			string
	IsValid()			bool
}


func GetParseMethodsMap() map[string]interface{} {
	return map[string]interface{} {
		"UserRole": ParseUserRoleString,
	}
}

func GetNewMethodsMap() map[string]interface{} {
	return map[string]interface{} {
		"UserRole": NewUserRole,
	}
}
