package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/generationoutput"
	"github.com/stablecog/sc-go/database/repository"
	"github.com/stablecog/sc-go/server/requests"
	"github.com/stablecog/sc-go/server/responses"
	"github.com/stablecog/sc-go/shared"
	"github.com/stretchr/testify/assert"
)

func TestGenerateUnauthorizedIfUserIdMissingInContext(t *testing.T) {
	reqBody := map[string]interface{}{
		"generate": "generate",
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	MockController.HandleCreateGeneration(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 401, resp.StatusCode)
	var respJson map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, "Unauthorized", respJson["error"])
}

func TestGenerateUnauthorizedIfUserIdNotUuid(t *testing.T) {
	reqBody := map[string]interface{}{
		"generate": "generate",
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx := context.WithValue(req.Context(), "user_id", "not-uuid")

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 401, resp.StatusCode)
	var respJson map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, "Unauthorized", respJson["error"])
}

func TestGenerateFailsWithInvalidStreamID(t *testing.T) {
	reqBody := requests.GenerateRequestBody{
		StreamID: "invalid",
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx := context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)
	var respJson map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, "invalid_stream_id", respJson["error"])
}

func TestGenerateEnforcesNumOutputsChange(t *testing.T) {
	reqBody := requests.GenerateRequestBody{
		StreamID: MockSSEId,
		Height:   shared.MAX_GENERATE_HEIGHT,
		Width:    shared.MAX_GENERATE_WIDTH,
		// Minimum is not enforced since it should default to 1
		NumOutputs: -1,
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx := context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)
	var respJson map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, "invalid_model_id", respJson["error"])

	// ! Max
	reqBody = requests.GenerateRequestBody{
		StreamID:   MockSSEId,
		Height:     shared.MAX_GENERATE_HEIGHT,
		Width:      shared.MAX_GENERATE_WIDTH,
		NumOutputs: shared.MAX_GENERATE_NUM_OUTPUTS + 1,
	}
	body, _ = json.Marshal(reqBody)
	w = httptest.NewRecorder()
	// Build request
	req = httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx = context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp = w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)
	respBody, _ = io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, fmt.Sprintf("Number of outputs can't be more than %d", shared.MAX_GENERATE_NUM_OUTPUTS), respJson["error"])
}

func TestGenerateEnforcesMaxWidthMaxHeight(t *testing.T) {
	reqBody := requests.GenerateRequestBody{
		StreamID: MockSSEId,
		Height:   shared.MAX_GENERATE_HEIGHT + 1,
		Width:    shared.MAX_GENERATE_WIDTH,
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx := context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)
	var respJson map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, fmt.Sprintf("Height is too large, max is: %d", shared.MAX_GENERATE_HEIGHT), respJson["error"])

	// ! Width
	reqBody = requests.GenerateRequestBody{
		StreamID: MockSSEId,
		Height:   shared.MAX_GENERATE_HEIGHT,
		Width:    shared.MAX_GENERATE_WIDTH + 1,
	}
	body, _ = json.Marshal(reqBody)
	w = httptest.NewRecorder()
	// Build request
	req = httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx = context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp = w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)
	respBody, _ = io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, fmt.Sprintf("Width is too large, max is: %d", shared.MAX_GENERATE_WIDTH), respJson["error"])
}

func TestGenerateRejectsInvalidModelOrScheduler(t *testing.T) {
	// ! invalid_scheduler_id
	reqBody := requests.GenerateRequestBody{
		StreamID:    MockSSEId,
		Height:      shared.MAX_GENERATE_HEIGHT,
		Width:       shared.MAX_GENERATE_WIDTH,
		SchedulerId: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		ModelId:     uuid.MustParse(repository.MOCK_GENERATION_MODEL_ID),
		NumOutputs:  1,
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx := context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)
	var respJson map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, "invalid_scheduler_id", respJson["error"])

	// ! invalid_model_id
	reqBody = requests.GenerateRequestBody{
		StreamID:    MockSSEId,
		Height:      shared.MAX_GENERATE_HEIGHT,
		Width:       shared.MAX_GENERATE_WIDTH,
		SchedulerId: uuid.MustParse(repository.MOCK_SCHEDULER_ID),
		ModelId:     uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		NumOutputs:  1,
	}
	body, _ = json.Marshal(reqBody)
	w = httptest.NewRecorder()
	// Build request
	req = httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx = context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp = w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)
	respBody, _ = io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &respJson)

	assert.Equal(t, "invalid_model_id", respJson["error"])
}

func TestGenerateNoCredits(t *testing.T) {
	// ! Perfectly valid request
	reqBody := requests.GenerateRequestBody{
		StreamID:       MockSSEId,
		Height:         shared.MAX_GENERATE_HEIGHT,
		Width:          shared.MAX_GENERATE_WIDTH,
		SchedulerId:    uuid.MustParse(repository.MOCK_SCHEDULER_ID),
		ModelId:        uuid.MustParse(repository.MOCK_GENERATION_MODEL_ID),
		NumOutputs:     1,
		InferenceSteps: shared.MAX_GENERATE_INTERFERENCE_STEPS_FREE + 1,
		Prompt:         "A portrait of a cat by Van Gogh",
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx := context.WithValue(req.Context(), "user_id", repository.MOCK_NO_CREDITS_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 400, resp.StatusCode)
	var errResp map[string]interface{}
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &errResp)
	assert.Equal(t, "insufficient_credits", errResp["error"])
}

func TestGenerateValidRequest(t *testing.T) {
	// ! Perfectly valid request
	reqBody := requests.GenerateRequestBody{
		StreamID:       MockSSEId,
		Height:         shared.MAX_GENERATE_HEIGHT,
		Width:          shared.MAX_GENERATE_WIDTH,
		SchedulerId:    uuid.MustParse(repository.MOCK_SCHEDULER_ID),
		ModelId:        uuid.MustParse(repository.MOCK_GENERATION_MODEL_ID),
		NumOutputs:     1,
		InferenceSteps: shared.MAX_GENERATE_INTERFERENCE_STEPS_FREE + 1,
		Prompt:         "A portrait of a cat by Van Gogh",
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx := context.WithValue(req.Context(), "user_id", repository.MOCK_NORMAL_UUID)

	MockController.HandleCreateGeneration(w, req.WithContext(ctx))
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
	var generateResp responses.QueuedResponse
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &generateResp)

	// Make sure valid uuid
	_, err := uuid.Parse(generateResp.ID)
	assert.Nil(t, err)

	// make sure we have this ID on our map
	streamid, _ := MockController.Redis.GetCogRequestStreamID(ctx, "first:"+generateResp.ID)
	assert.Equal(t, MockSSEId, streamid)
}

func TestSubmitGenerationToGallery(t *testing.T) {
	// ! Generation that doesnt exist
	reqBody := requests.GenerateSubmitToGalleryRequestBody{
		GenerationOutputIDs: []uuid.UUID{uuid.New()},
	}
	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	// Build request
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx := context.WithValue(req.Context(), "user_id", repository.MOCK_NORMAL_UUID)

	MockController.HandleSubmitGenerationToGallery(w, req.WithContext(ctx))
	resp := w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
	var submitResp responses.GenerateSubmitToGalleryResponse
	respBody, _ := io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &submitResp)

	assert.Equal(t, 0, submitResp.Submitted)

	// ! Generation that does exist
	// Retrieve generation
	meta, err := MockController.Repo.GetUserGenerations(uuid.MustParse(repository.MOCK_ADMIN_UUID), 50, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, meta.Generations)
	var idx int
	var outputIdx int
	// Find generation that is not submitted
	for i, generation := range meta.Generations {
		if len(generation.Outputs) > 0 {
			for oidx, output := range generation.Outputs {
				if output.GalleryStatus == generationoutput.GalleryStatusNotSubmitted {
					idx = i
					outputIdx = oidx
					break
				}
			}
		}
	}
	assert.Equal(t, generationoutput.GalleryStatusNotSubmitted, meta.Generations[idx].Outputs[outputIdx].GalleryStatus)

	reqBody = requests.GenerateSubmitToGalleryRequestBody{
		GenerationOutputIDs: []uuid.UUID{meta.Generations[idx].Outputs[outputIdx].ID},
	}
	body, _ = json.Marshal(reqBody)
	w = httptest.NewRecorder()
	// Build request
	req = httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx = context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleSubmitGenerationToGallery(w, req.WithContext(ctx))
	resp = w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
	respBody, _ = io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &submitResp)
	assert.Equal(t, 1, submitResp.Submitted)

	// Make sure updated in DB
	meta, err = MockController.Repo.GetUserGenerations(uuid.MustParse(repository.MOCK_ADMIN_UUID), 50, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, meta.Generations)
	assert.Equal(t, generationoutput.GalleryStatusSubmitted, meta.Generations[idx].Outputs[0].GalleryStatus)

	// ! Generation that is already submitted
	// Retrieve generation
	meta, err = MockController.Repo.GetUserGenerations(uuid.MustParse(repository.MOCK_ADMIN_UUID), 50, nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, meta.Generations)
	assert.Equal(t, generationoutput.GalleryStatusSubmitted, meta.Generations[idx].Outputs[outputIdx].GalleryStatus)

	reqBody = requests.GenerateSubmitToGalleryRequestBody{
		GenerationOutputIDs: []uuid.UUID{meta.Generations[idx].Outputs[outputIdx].ID},
	}
	body, _ = json.Marshal(reqBody)
	w = httptest.NewRecorder()
	// Build request
	req = httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Setup context
	ctx = context.WithValue(req.Context(), "user_id", repository.MOCK_ADMIN_UUID)

	MockController.HandleSubmitGenerationToGallery(w, req.WithContext(ctx))
	resp = w.Result()
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
	respBody, _ = io.ReadAll(resp.Body)
	json.Unmarshal(respBody, &submitResp)

	assert.Equal(t, 0, submitResp.Submitted)
}
