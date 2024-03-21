package custom_errors

import "errors"

var (
	// When no api key is provided
	ErrNoAPI = errors.New("API Key is required")

	// when invalid model is passed
	ErrInvalidModel = errors.New("invalid model")

	// when no message is provided
	ErrNoMessages = errors.New("no messages provided")

	// when invalid role is used
	ErrInvalidRole = errors.New("invalid role. Only `user`, `system` and `assistant` are supported")

	// when invalid temperature value is sent in the request
	ErrInvalidTemp = errors.New("invalid temperature. 0<= temp <= 2")

	// when invalid presence penalty is provided
	ErrInvalidPresencePenalty = errors.New("invalid presence penalty. -2<= presence penalty <= 2")

	// when invalid frequency penalty is provided
	ErrInvalidFrequencyPenalty = errors.New("invalid frequency penalty. -2<= frequency penalty <= 2")
)
