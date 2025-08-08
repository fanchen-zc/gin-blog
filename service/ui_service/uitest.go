package ui_service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gin-test/models"
	"gin-test/pkg/logging"
	"gin-test/pkg/setting"
	"github.com/sashabaranov/go-openai"
	"io/ioutil"
	"strings"
)

type UITest struct {
	ImagePath string
}

func (s *UITest) Analyse() (*models.UIAnalysis, error) {
	// 1. Read the image file
	imageData, err := ioutil.ReadFile(s.ImagePath)
	if err != nil {
		logging.Error(fmt.Sprintf("Failed to read image file: %v", err))
		return nil, err
	}

	// 2. Encode the image to base64
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	// 3. Create OpenAI client
	if setting.OpenAISetting.ApiKey == "YOUR_API_KEY_HERE" || setting.OpenAISetting.ApiKey == "" {
		logging.Warn("OpenAI API key is not set. Returning mock data.")
		return s.getMockAnalysis(), nil // Or return an error
	}
	client := openai.NewClient(setting.OpenAISetting.ApiKey)

	// 4. Construct the prompt
	prompt := `
	Please analyze the user interface in the provided screenshot.
	Identify any potential usability problems or areas for improvement.
	Structure your response as a JSON object with two keys: "problems" and "suggestions".
	- "problems" should be an array of objects, where each object has "element", "description", and "suggestion" fields.
	- "suggestions" should be an array of objects, where each object has "area", "description", and "suggestion" fields.
	Example:
	{
	  "problems": [
	    {
	      "element": "Login Button",
	      "description": "The button's color has low contrast with the background.",
	      "suggestion": "Increase the contrast ratio to meet WCAG AA standards."
	    }
	  ],
	  "suggestions": [
	    {
	      "area": "Overall Layout",
	      "description": "The layout seems cluttered.",
	      "suggestion": "Consider adding more whitespace between sections."
	    }
	  ]
	}
	`

	// 5. Send the request to the OpenAI API
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4VisionPreview,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeText,
						Text: prompt,
					},
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ImageURL{
							URL:    fmt.Sprintf("data:image/png;base64,%s", base64Image),
							Detail: openai.ImageURLDetailAuto,
						},
					},
				},
			},
		},
		MaxTokens: 2000,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		logging.Error(fmt.Sprintf("OpenAI API call failed: %v", err))
		return nil, err
	}

	if len(resp.Choices) == 0 {
		logging.Error("OpenAI API returned no choices.")
		return nil, fmt.Errorf("no response from AI")
	}

	// 6. Unmarshal the JSON response
	var analysisResult models.UIAnalysis
	// The response content might be wrapped in markdown JSON block, let's clean it.
	jsonContent := strings.Trim(resp.Choices[0].Message.Content, "```json\n")
	err = json.Unmarshal([]byte(jsonContent), &analysisResult)
	if err != nil {
		logging.Error(fmt.Sprintf("Failed to unmarshal AI response: %v. Response was: %s", err, resp.Choices[0].Message.Content))
		return nil, err
	}

	// 7. Return the populated struct
	return &analysisResult, nil
}

// getMockAnalysis returns a mock analysis for when the API key is not set.
func (s *UITest) getMockAnalysis() *models.UIAnalysis {
	return &models.UIAnalysis{
		Problems: []models.Problem{
			{
				Element:     "Header",
				Description: "The logo is too small and hard to read.",
				Suggestion:  "Increase the size of the logo by 20%.",
			},
		},
		Suggestions: []models.Suggestion{
			{
				Area:        "Overall Layout",
				Description: "The layout is too cluttered. There is not enough white space between elements.",
				Suggestion:  "Increase the margin and padding between elements to improve readability.",
			},
		},
	}
}
