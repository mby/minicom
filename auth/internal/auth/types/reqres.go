package types

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		Status string `json:"status"`
	}
)

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}
)

type (
	VerifyRequest struct {
		Token string `json:"token"`
	}

	VerifyResponse struct {
		Username string `json:"username"`
	}
)
