// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package misc

import (
	"encoding/json"
	"strings"

	"github.com/haproxytech/dataplaneapi/haproxy"

	"github.com/haproxytech/client-native/configuration"
	"github.com/haproxytech/models"
)

const (
	// ErrHTTPNotFound HTTP status code 404
	ErrHTTPNotFound = int64(404)
	// ErrHTTPConflict HTTP status code 409
	ErrHTTPConflict = int64(409)
	// ErrHTTPInternalServerError HTTP status code 500
	ErrHTTPInternalServerError = int64(500)
	// ErrHTTPBadRequest HTTP status code 400
	ErrHTTPBadRequest = int64(400)
)

// HandleError translates error codes from client native into models.Error with appropriate http status code
func HandleError(err error) *models.Error {
	switch t := err.(type) {
	case *configuration.ConfError:
		msg := t.Error()
		httpCode := ErrHTTPInternalServerError
		switch t.Code() {
		case configuration.ErrObjectDoesNotExist:
			httpCode = ErrHTTPNotFound
		case configuration.ErrObjectAlreadyExists, configuration.ErrVersionMismatch, configuration.ErrTransactionAlreadyExists:
			httpCode = ErrHTTPConflict
		case configuration.ErrObjectIndexOutOfRange, configuration.ErrValidationError, configuration.ErrBothVersionTransaction,
			configuration.ErrNoVersionTransaction, configuration.ErrNoParentSpecified, configuration.ErrParentDoesNotExist,
			configuration.ErrTransactionDoesNotExist:
			httpCode = ErrHTTPBadRequest
		}
		return &models.Error{Code: &httpCode, Message: &msg}
	case *haproxy.ReloadError:
		httpCode := ErrHTTPBadRequest
		msg := t.Error()
		return &models.Error{Code: &httpCode, Message: &msg}
	default:
		msg := t.Error()
		code := ErrHTTPInternalServerError
		return &models.Error{Code: &code, Message: &msg}
	}
}

// DiscoverChildPaths return children models.Endpoints given path
func DiscoverChildPaths(path string, spec json.RawMessage) (models.Endpoints, error) {
	var m map[string]interface{}
	err := json.Unmarshal(spec, &m)
	if err != nil {
		return nil, err
	}
	es := make(models.Endpoints, 0, 1)
	paths := m["paths"].(map[string]interface{})
	for key, value := range paths {
		v := value.(map[string]interface{})
		g := v["get"].(map[string]interface{})
		title := g["summary"].(string)
		description := g["description"].(string)

		if strings.HasPrefix(key, path) && key != path {
			if len(strings.Split(key[len(path)+1:], "/")) == 1 {
				e := models.Endpoint{
					URL:         key,
					Title:       title,
					Description: description,
				}
				es = append(es, &e)
			}
		}
	}
	return es, nil
}
