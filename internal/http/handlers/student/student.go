package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/ritam-g/students-api/internal/types"
	"github.com/ritam-g/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var stu types.Stuendt

		err := json.NewDecoder(r.Body).Decode(& stu)

		if errors.Is(err, io.EOF) {
			response.WriteJson( w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		slog.Info("creating a student details ")

		response.WriteJson(w,http.StatusCreated,map[string]string{"success":"ok"})
	}
}
