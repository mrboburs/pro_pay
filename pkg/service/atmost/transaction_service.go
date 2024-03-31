package atmost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"pro_pay/model"
	"pro_pay/pkg/repository"
	"pro_pay/pkg/store"
	"pro_pay/tools/logger"
	"pro_pay/tools/response"

	"google.golang.org/grpc/codes"
)

type TransactionService struct {
	repo    *repository.Repository
	minio   *store.Store
	loggers *logger.Logger
}

func NewTransactionService(repo *repository.Repository, minio *store.Store,
	loggers *logger.Logger) *TransactionService {
	return &TransactionService{repo: repo, minio: minio, loggers: loggers}
}

func (s *TransactionService) CreateTransaction(in model.CreateTransaction) (out *model.Response, err error) {
	// URL to which the POST request will be sent
	url := "https://partner.atmos.uz/merchant/pay/create"

	// Convert data to JSON
	jsonData, err := json.Marshal(in)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	accessToken := GetToken()

	// Set the Authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken.AccessToken)

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response

	if err = json.NewDecoder(resp.Body).Decode(&out); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	id, err := s.repo.AtmostRepo.CreateTransaction(in)

	if err != nil {
		return out, response.ServiceError(err, codes.Internal)
	}
	out.Result.ID=id
	return out, nil
}
