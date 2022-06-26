package helper

func ResponseFailed(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "404",
		"message": msg,
	}
}

func ResponseSuccessNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "200",
		"message": msg,
	}
}

func ResponseSuccessWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "200",
		"message": msg,
		"data":    data,
	}
}
