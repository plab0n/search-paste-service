package model

type EmbeddingRequestBody struct {
	Input string
}
type Data struct {
	Embedding []float64 `json:"embedding"`
}

type Usage struct {
	Prompt_Tokens int `json:"prompt_tokens"`
	Total_Tokens  int `json:"total_tokens"`
}

type EmbeddingResponse struct {
	Data  Data   `json:"data"`
	Model string `json:"model"`
	Usage Usage  `json:"usage"`
}
