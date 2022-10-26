package model

func ResponseError (statusCode int, err error) map[string]interface{} {
	return map[string]interface{}{
		"status_code": statusCode,
		"message":     err,
	}
}
