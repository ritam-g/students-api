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

		err := json.NewDecoder(r.Body).Decode(&stu)
		slog.Info("creating a student details ")
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return 
		}

		

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
