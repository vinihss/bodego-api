package database

type Model interface {
}

func NewModelByTypeName(typeName string, pkg interface{}) (interface{}, error) {
	// Obter o tipo do pacote
	pkgType := reflect.TypeOf(pkg).Elem()

	// Procurar o tipo pelo nome no pacote
	for i := 0; i < pkgType.NumField(); i++ {
		field := pkgType.Field(i)
		if field.Name == typeName {
			return reflect.New(field.Type).Interface(), nil
		}
	}

	return nil, fmt.Errorf("type %s not found in package", typeName)
}

// HydrateModel copia os campos de um tipo de origem para um tipo de destino
func HydrateModel(source interface{}, target interface{}) error {
	sourceValue := reflect.ValueOf(source)
	targetValue := reflect.ValueOf(target)

	// Garantir que ambos sejam structs ou ponteiros para structs
	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}
	if targetValue.Kind() == reflect.Ptr {
		targetValue = targetValue.Elem()
	}

	if sourceValue.Kind() != reflect.Struct || targetValue.Kind() != reflect.Struct {
		return errors.New("source and target must be structs or pointers to structs")
	}

	// Iterar sobre os campos do tipo de origem
	for i := 0; i < sourceValue.NumField(); i++ {
		sourceField := sourceValue.Type().Field(i)
		sourceFieldValue := sourceValue.Field(i)

		// Procurar o campo correspondente no tipo de destino
		targetField := targetValue.FieldByName(sourceField.Name)
		if targetField.IsValid() && targetField.CanSet() && targetField.Type() == sourceFieldValue.Type() {
			targetField.Set(sourceFieldValue)
		}
	}

	return nil
}
