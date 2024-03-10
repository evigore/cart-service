package retry

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type retryRoundTripper struct {
	retries int
}

const httpStatusEnhanceYourCalm = 420

func New(retries int) http.RoundTripper {
	return &retryRoundTripper{
		retries,
	}
}

func (rt retryRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll: %w", err)
	}

	for i := 0; i < rt.retries; i++ {
		req.Body = io.NopCloser(bytes.NewReader(body))
		res, err := http.DefaultTransport.RoundTrip(req)
		if err != nil {
			return nil, fmt.Errorf("failed http.DefaultTransport.RoundTrip: %w", err)
		}

		if res.StatusCode != httpStatusEnhanceYourCalm && res.StatusCode != http.StatusTooManyRequests {
			return res, nil
		}
	}

	return nil, fmt.Errorf("cannot get response, retry later: %w", err)
}
