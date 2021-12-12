package auth

import (
	"net/http"
	"strings"

	"github.com/mby/minicom/auth/internal/auth/types"
	"github.com/mby/minicom/auth/internal/cfg"
	"github.com/mby/minicom/auth/internal/errors"
	"github.com/mby/minicom/auth/internal/helpers"
)

type Handler struct {
	authRepo IRepo
}

func NewHandler(cfg cfg.Config) Handler {
	h := Handler{
		authRepo: NewRepo(cfg),
	}

	http.HandleFunc("/register", h.Register)
	http.HandleFunc("/login", h.Login)

	http.HandleFunc("/health", h.Health)
	return h
}

func (h Handler) Cleanup() {
	defer h.authRepo.Cleanup()
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write(helpers.Res(errors.MethodNotAllowed))
		return
	}

	req := types.RegisterRequest{}
	if helpers.Body(r, &req) != nil {
		w.Write(helpers.Res(errors.InvalidRequestBody))
		return
	}

	req.Username = strings.ToLower(strings.TrimSpace(req.Username))
	req.Password = strings.TrimSpace(req.Password)
	if req.Username == "" || req.Password == "" {
		w.Write(helpers.Res(errors.InvalidRequestBody))
		return
	}

	if err := h.authRepo.CreateUser(req.Username, req.Password); err != nil {
		w.Write(helpers.Res(err))
		return
	}

	w.Write(helpers.Res(types.RegisterResponse{Status: "ok"}))
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write(helpers.Res(errors.MethodNotAllowed))
		return
	}

	req := types.LoginRequest{}
	if helpers.Body(r, &req) != nil {
		w.Write(helpers.Res(errors.InvalidRequestBody))
		return
	}

	token, err := h.authRepo.Login(req.Username, req.Password)
	if err != nil {
		w.Write(helpers.Res(err))
	}

	w.Write(helpers.Res(types.LoginResponse{Token: token}))
}

func (h Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"status\": \"ok\"}"))
}
