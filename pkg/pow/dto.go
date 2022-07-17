package pow

type (
	request struct {
		Complexity uint8  `json:"complexity"`
		Resource   []byte `json:"resource"`
	}

	response struct {
		Result []byte `json:"result"`
	}

	accept struct {
	}
)
