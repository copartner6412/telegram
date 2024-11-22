package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
)

// send sends an HTTP request to the Telegram API.
func (b *Bot) send(httpMethod, telegramMethod, contentTypeHeader string, requestPayload io.ReadWriter) (*http.Response, error) {
	url := fmt.Sprintf(telegramEndpoint, b.Token, telegramMethod)

	httpRequest, err := http.NewRequest(httpMethod, url, requestPayload)
	if err != nil {
		return nil, fmt.Errorf("error creating new request: %w", err)
	}

	httpRequest.Header.Set("Content-Type", contentTypeHeader)

	httpResponse, err := b.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Telegram API error: HTTP status %s", httpResponse.Status)
	}

	return httpResponse, nil
}

func (b *Bot) sendJson(httpMethod, telegramMethod string, request any) (*http.Response, error) {
	requestPayload := new(bytes.Buffer)
	if err := json.NewEncoder(requestPayload).Encode(request); err != nil {
		return nil, fmt.Errorf("error encoding request payload: %w", err)
	}

	httpResponse, err := b.send(httpMethod, telegramMethod, "application/json", requestPayload)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

func (b *Bot) sendJsonForBool(httpMethod, telegramMethod string, request any) error {
	httpResponse, err := b.sendJson(httpMethod, telegramMethod, request)
	if err != nil {
		return fmt.Errorf("error sending \"%s\" request to method \"%s\" in JSON format: %w", httpMethod, telegramMethod, err)
	}
	defer httpResponse.Body.Close()

	var response struct {
		Ok          bool    `json:"ok"`
		Result      bool    `json:"result"`
		Description *string `json:"description"`
		ErrorCode   *int    `json:"error_code"`
		// Parameters  ResponseParameters `json:"parameters"`
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if !response.Ok || !response.Result || response.ErrorCode != nil || response.Description != nil {
		errorMessage := fmt.Sprintf("Telegram API error: HTTP status \"%s\"", httpResponse.Status)
		if !response.Ok || !response.Result {
			errorMessage += ", result \"false\" (request unsuccessful)"
		} else {
			errorMessage += ", result \"true\" (request successful)"
		}
		if response.ErrorCode != nil {
			errorMessage += fmt.Sprintf(", Telegram error code \"%d\"", *response.ErrorCode)
		}
		if response.Description != nil {
			errorMessage += fmt.Sprintf(", description: %s", *response.Description)
		}
		return errors.New(errorMessage)
	}

	return nil
}

func (b *Bot) sendJsonForResult(result any, httpMethod, telegramMethod string, request any) error {
	httpResponse, err := b.sendJson(httpMethod, telegramMethod, request)
	if err != nil {
		return fmt.Errorf("error sending \"%s\" request to method \"%s\" in JSON format: %w", httpMethod, telegramMethod, err)
	}
	defer httpResponse.Body.Close()

	var response struct {
		Ok          bool            `json:"ok"`
		Result      json.RawMessage `json:"result,omitempty"`
		ErrorCode   *int            `json:"error_code,omitempty"`
		Description *string         `json:"description,omitempty"`
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if httpResponse.StatusCode != http.StatusOK || !response.Ok || response.ErrorCode != nil || response.Description != nil {
		errorMessage := fmt.Sprintf("Telegram API error: HTTP status \"%s\"", httpResponse.Status)
		if !response.Ok {
			errorMessage += ", request unsuccessful"
		} else {
			errorMessage += ", request successful"
		}
		if response.ErrorCode != nil {
			errorMessage += fmt.Sprintf(", Telegram error code \"%d\"", *response.ErrorCode)
		}
		if response.Description != nil {
			errorMessage += fmt.Sprintf(", description: %s", *response.Description)
		}
		return errors.New(errorMessage)
	}

	if result == nil || len(response.Result) == 0 {
		return nil
	}

	if err := json.Unmarshal(response.Result, result); err != nil {
		return fmt.Errorf("error decoding result into specified type: %w", err)
	}

	return nil
}

// func (b *Bot) sendJsonForResult(result any, httpMethod, telegramMethod string, request any) error {
// 	httpResponse, err := b.sendJson(httpMethod, telegramMethod, request)
// 	if err != nil {
// 		return fmt.Errorf("error sending \"%s\" request to method \"%s\" in JSON format: %w", httpMethod, telegramMethod, err)
// 	}
// 	defer httpResponse.Body.Close()

// 	var intermediateResponse map[string]interface{}
// 	if err := json.NewDecoder(httpResponse.Body).Decode(&intermediateResponse); err != nil {
// 		return fmt.Errorf("error decoding response: %w", err)
// 	}

// 	ok, _ := intermediateResponse["ok"].(bool)
// 	errorCode, errorCodeExists := intermediateResponse["error_code"].(float64)
// 	description, descriptionExists := intermediateResponse["description"].(string)
// 	if !ok || errorCodeExists || descriptionExists {
// 		errorMessage := fmt.Sprintf("Telegram API error: HTTP status \"%s\"", httpResponse.Status)
// 		if !ok {
// 			errorMessage += ", request unsuccessful"
// 		} else {
// 			errorMessage += ", request successful"
// 		}
// 		if errorCodeExists {
// 			errorMessage += fmt.Sprintf(", Telegram error code \"%d\"", int(math.Round(errorCode)))
// 		}
// 		if descriptionExists {
// 			errorMessage += fmt.Sprintf(", description: %s", description)
// 		}
// 		return errors.New(errorMessage)
// 	}

// 	resultData, err := json.Marshal(intermediateResponse["result"])
// 	if err != nil {
// 		return fmt.Errorf("error re-encoding result: %w", err)
// 	}

// 	if err := json.Unmarshal(resultData, result); err != nil {
// 		return fmt.Errorf("error decoding result into specified type: %w", err)
// 	}

// 	return nil
// }

func StructToMapStringByJSONTag(fields map[string]string, structVariable any) error {
	structValue := reflect.ValueOf(structVariable)
	structType := reflect.TypeOf(structVariable)

	if structType.Kind() == reflect.Pointer {
		if structValue.IsNil() {
			return nil
		}
		structValue = structValue.Elem()
		structType = structType.Elem()
	}

	if structType.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct but got %s", structType.Kind())
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldType := field.Type
		fieldValue := structValue.Field(i)

		if field.Type.Kind() == reflect.Pointer {
			if fieldValue.IsNil() {
				continue
			}
			fieldType = fieldType.Elem()
			fieldValue = fieldValue.Elem()
		}

		if field.Anonymous && fieldType.Kind() == reflect.Struct {
			if err := StructToMapStringByJSONTag(fields, fieldValue.Interface()); err != nil {
				return fmt.Errorf("error processing embedded field %s: %w", field.Name, err)
			}
			continue
		}

		jsonTag, ok := field.Tag.Lookup("json")
		if !ok || jsonTag == "" || jsonTag == "-" {
			continue
		}
		jsonTagValue := strings.Split(jsonTag, ",")[0]

		switch fieldType.Kind() {
		case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice, reflect.Interface:
			if fieldValue.IsZero() {
				continue
			}
			jsonData, err := json.Marshal(fieldValue.Interface())
			if err != nil {
				return fmt.Errorf("error JSON marshaling field %s: %w", field.Name, err)
			}
			fields[jsonTagValue] = string(jsonData)
			continue
		case reflect.String:
			fields[jsonTagValue] = fieldValue.String()
			continue

		}

		fields[jsonTagValue] = fmt.Sprintf("%v", fieldValue.Interface())
	}

	return nil
}

// func structToMapByJSONTag(fields map[string]any, structVariable any) error {
// 	structValue := reflect.ValueOf(structVariable)
// 	structType := reflect.TypeOf(structVariable)

// 	if structType.Kind() == reflect.Pointer {
// 		if structValue.IsNil() {
// 			return nil
// 		}
// 		structValue = structValue.Elem()
// 		structType = structType.Elem()
// 	}

// 	if structType.Kind() != reflect.Struct {
// 		return fmt.Errorf("expected struct but got %s", structType.Kind())
// 	}

// 	for i := 0; i < structType.NumField(); i++ {
// 		field := structType.Field(i)
// 		if !field.IsExported() {
// 			continue
// 		}
// 		fieldType := field.Type
// 		fieldValue := structValue.Field(i)

// 		if field.Type.Kind() == reflect.Pointer {
// 			if fieldValue.IsNil() {
// 				continue
// 			}
// 			fieldType = fieldType.Elem()
// 			fieldValue = fieldValue.Elem()
// 		}

// 		if field.Anonymous && fieldType.Kind() == reflect.Struct {
// 			if err := structToMapByJSONTag(fields, fieldValue.Interface()); err != nil {
// 				return fmt.Errorf("error processing embedded field %s: %w", field.Name, err)
// 			}
// 			continue
// 		}

// 		jsonTag, ok := field.Tag.Lookup("json")
// 		if !ok || jsonTag == "" || jsonTag == "-" {
// 			continue
// 		}
// 		jsonTagParts := strings.Split(jsonTag, ",")

// 		if fieldValue.IsZero() {
// 			continue
// 		}

// 		fields[jsonTagParts[0]] = fieldValue.Interface()
// 	}

// 	return nil
// }

func (b *Bot) sendMultipart(httpMethod, telegramMethod string, request any, fileFields, attachFileFields []string) (*http.Response, error) {
	fields := make(map[string]string)
	if err := StructToMapStringByJSONTag(fields, request); err != nil {
		return nil, fmt.Errorf("error converting request struct to map by JSON tags: %w", err)
	}

	requestPayload := new(bytes.Buffer)
	writer := multipart.NewWriter(requestPayload)

	for field, value := range fields {
		var isFile bool
		for _, fileField := range fileFields {
			if field == fileField {
				isFile = true
			}
		}
		if isFile {
			var needsAttach bool
			for _, attachFileField := range attachFileFields {
				if field == attachFileField {
					needsAttach = true
				}
			}
			if strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") || strings.HasPrefix(value, "file_id:") {
				if err := writer.WriteField(field, strings.TrimPrefix(value, "file_id:")); err != nil {
					return nil, fmt.Errorf("error adding %s field: %w", field, err)
				}
			} else if path.IsAbs(value) {
				file, err := os.Open(value)
				if err != nil {
					return nil, fmt.Errorf("error opening %s file: %w", field, err)
				}
				defer file.Close()

				filename := strings.ReplaceAll(strings.TrimSpace(file.Name()), " ", "_")
				if needsAttach {
					filename = "attach://" + filename
				}
				fileWriter, err := writer.CreateFormFile(field, filename)
				if err != nil {
					return nil, fmt.Errorf("error creating form file for %s: %w", field, err)
				}
				if _, err := io.Copy(fileWriter, file); err != nil {
					return nil, fmt.Errorf("error copying %s file: %w", field, err)
				}
			} else {
				return nil, fmt.Errorf("value %s of field %s is neither URL (with \"http://\" or \"https://\" prefix), file ID (with \"file_id:\" prefix), nor an absolute local path", value, field)
			}
			continue
		}

		err := writer.WriteField(field, value)
		if err != nil {
			return nil, fmt.Errorf("error adding %s field: %w", field, err)
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing multipart message: %w", err)
	}

	httpResponse, err := b.send(httpMethod, telegramMethod, writer.FormDataContentType(), requestPayload)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

func (b *Bot) sendMultipartForBool(httpMethod, telegramMethod string, request any, fileFields, attachFileFields []string) error {
	httpResponse, err := b.sendMultipart(httpMethod, telegramMethod, request, fileFields, attachFileFields)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	var response struct {
		Ok          bool    `json:"ok"`
		Result      bool    `json:"result"`
		Description *string `json:"description"`
		ErrorCode   *int    `json:"error_code"`
		// Parameters  ResponseParameters `json:"parameters"`
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if !response.Ok || !response.Result || response.ErrorCode != nil || response.Description != nil {
		errorMessage := fmt.Sprintf("Telegram API error: HTTP status \"%s\"", httpResponse.Status)
		if !response.Ok || !response.Result {
			errorMessage += ", result \"false\" (request unsuccessful)"
		} else {
			errorMessage += ", result \"true\" (request successful)"
		}
		if response.ErrorCode != nil {
			errorMessage += fmt.Sprintf(", Telegram error code \"%d\"", *response.ErrorCode)
		}
		if response.Description != nil {
			errorMessage += fmt.Sprintf(", description: %s", *response.Description)
		}
		return errors.New(errorMessage)
	}
	return nil
}

func (b *Bot) sendMultipartForResult(result any, httpMethod, telegramMethod string, request any, fileFields, attachFileFields []string) error {
	httpResponse, err := b.sendMultipart(httpMethod, telegramMethod, request, fileFields, attachFileFields)
	if err != nil {
		return fmt.Errorf("error sending \"%s\" request to Telegram method endpoint \"%s\" in multipart format: %w", httpMethod, telegramMethod, err)
	}
	defer httpResponse.Body.Close()

	var intermediateResponse map[string]interface{}
	if err := json.NewDecoder(httpResponse.Body).Decode(&intermediateResponse); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	ok, _ := intermediateResponse["ok"].(bool)
	errorCode, errorCodeExists := intermediateResponse["error_code"].(float64)
	description, descriptionExists := intermediateResponse["description"].(string)
	if !ok || errorCodeExists || descriptionExists {
		errorMessage := fmt.Sprintf("Telegram API error: HTTP status \"%s\"", httpResponse.Status)
		if !ok {
			errorMessage += ", request unsuccessful"
		} else {
			errorMessage += ", request successful"
		}
		if errorCodeExists {
			errorMessage += fmt.Sprintf(", Telegram error code \"%d\"", int(math.Round(errorCode)))
		}
		if descriptionExists {
			errorMessage += fmt.Sprintf(", description: %s", description)
		}
		return errors.New(errorMessage)
	}

	resultData, err := json.Marshal(intermediateResponse["result"])
	if err != nil {
		return fmt.Errorf("error re-encoding result: %w", err)
	}

	if err := json.Unmarshal(resultData, result); err != nil {
		return fmt.Errorf("error decoding result into specified type: %w", err)
	}

	return nil
}

func writeNestedMultipart(nestedWriter io.Writer, request any, fileFields, attachFileFields []string) error {
	writer := multipart.NewWriter(nestedWriter)
	fields := make(map[string]string)
	if err := StructToMapStringByJSONTag(fields, request); err != nil {
		return fmt.Errorf("error converting request struct to map by JSON tags: %w", err)
	}

	for field, value := range fields {
		var isFile bool
		for _, fileField := range fileFields {
			if field == fileField {
				isFile = true
			}
		}
		if isFile {
			var needsAttach bool
			for _, attachFileField := range attachFileFields {
				if field == attachFileField {
					needsAttach = true
				}
			}
			if strings.HasPrefix(value, "http://") || strings.HasPrefix(value, "https://") || strings.HasPrefix(value, "file_id:") {
				if err := writer.WriteField(field, strings.TrimPrefix(value, "file_id:")); err != nil {
					return fmt.Errorf("error adding %s field: %w", field, err)
				}
			} else if path.IsAbs(value) {
				file, err := os.Open(value)
				if err != nil {
					return fmt.Errorf("error opening %s file: %w", field, err)
				}
				defer file.Close()

				filename := strings.ReplaceAll(strings.TrimSpace(file.Name()), " ", "_")
				if needsAttach {
					filename = "attach://" + filename
				}
				fileWriter, err := writer.CreateFormFile(field, filename)
				if err != nil {
					return fmt.Errorf("error creating form file for %s: %w", field, err)
				}
				if _, err := io.Copy(fileWriter, file); err != nil {
					return fmt.Errorf("error copying %s file: %w", field, err)
				}
			} else {
				return fmt.Errorf("value %s of field %s is neither URL (with \"http://\" or \"https://\" prefix), file ID (with \"file_id:\" prefix), nor an absolute local path", value, field)
			}
			continue
		}

		err := writer.WriteField(field, value)
		if err != nil {
			return fmt.Errorf("error adding %s field: %w", field, err)
		}
	}

	return nil
}
