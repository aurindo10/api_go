package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddRoutes(
	mux *http.ServeMux,
) {
	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("/ola", HandleOla())
}

type Validator interface {
	Valid(ctx context.Context) (problems map[string]string)
}

type ress struct {
	Name string
}

func (r ress) Valid(ctx context.Context) (problems map[string]string) {
	println("validado com sucesso")
	return nil
}
func HandleOla() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, _, _ := decodeValid[ress](r)
		encode(w, r, 200, res)
	})
}

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decodeValid[T Validator](r *http.Request) (T, map[string]string, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, nil, fmt.Errorf("decode json: %w", err)
	}
	if problems := v.Valid(r.Context()); len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T: %d problems", v, len(problems))
	}
	return v, nil, nil
}
