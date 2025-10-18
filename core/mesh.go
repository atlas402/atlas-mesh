package core

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"crypto/rand"
)

type AtlasMesh struct {
	facilitatorURL string
	merchantAddress string
	services       map[string]*ServiceRegistrationParams
	httpClient     *http.Client
}

func New(facilitatorURL, merchantAddress string) *AtlasMesh {
	return &AtlasMesh{
		facilitatorURL: facilitatorURL,
		merchantAddress: merchantAddress,
		services:       make(map[string]*ServiceRegistrationParams),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (m *AtlasMesh) RegisterService(ctx context.Context, params *ServiceRegistrationParams) (*ServiceRegistrationResult, error) {
	serviceID := generateServiceID()
	priceMicro := fmt.Sprintf("%d", int(parseFloat(params.Price)*1000000))

	registrationData := map[string]interface{}{
		"id":             serviceID,
		"name":           params.Name,
		"description":    params.Description,
		"endpoint":       params.Endpoint,
		"category":       params.Category,
		"network":        params.Network,
		"merchantAddress": params.MerchantAddress,
		"accepts": []map[string]interface{}{
			{
				"asset":            getAssetAddress(params.Network),
				"payTo":            params.MerchantAddress,
				"network":          params.Network,
				"maxAmountRequired": priceMicro,
				"scheme":           params.Scheme,
				"mimeType":         "application/json",
			},
		},
		"metadata": params.Metadata,
	}

	if err := m.registerWithFacilitator(ctx, registrationData); err != nil {
		return nil, err
	}

	m.services[serviceID] = params

	return &ServiceRegistrationResult{
		ServiceID:      serviceID,
		FacilitatorURL: fmt.Sprintf("%s/discovery/resources/%s", m.facilitatorURL, serviceID),
	}, nil
}

func (m *AtlasMesh) registerWithFacilitator(ctx context.Context, data map[string]interface{}) error {
	url := fmt.Sprintf("%s/discovery/resources", m.facilitatorURL)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := m.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("registration failed with status %d", resp.StatusCode)
	}

	return nil
}

type ServiceRegistrationParams struct {
	Name          string
	Description   string
	Endpoint      string
	Category      string
	Price         string
	Network       string
	Scheme        string
	MerchantAddress string
	Metadata      map[string]interface{}
}

type ServiceRegistrationResult struct {
	ServiceID      string
	FacilitatorURL string
}

func generateServiceID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("service-%x", b)
}

func getAssetAddress(network string) string {
	if network == "base" {
		return "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"
	}
	return "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
}

func parseFloat(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}



