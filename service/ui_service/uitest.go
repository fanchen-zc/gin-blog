package ui_service

import "gin-test/models"

type UITest struct {
	ImagePath string
}

func (s *UITest) Analyse() (*models.UIAnalysis, error) {
	// In a real application, you would implement the logic to call the GPT-5 API
	// and perform the UI analysis here.
	// For this example, we'll just return a mock analysis.

	// The image path is available in s.ImagePath if you need to use it.

	return &models.UIAnalysis{
		Problems: []models.Problem{
			{
				Element:     "Header",
				Description: "The logo is too small and hard to read.",
				Suggestion:  "Increase the size of the logo by 20%.",
			},
			{
				Element:     "Navigation Bar",
				Description: "The contrast between the text and the background is too low, making it difficult to read.",
				Suggestion:  "Use a darker color for the text or a lighter color for the background.",
			},
		},
		Suggestions: []models.Suggestion{
			{
				Area:        "Overall Layout",
				Description: "The layout is too cluttered. There is not enough white space between elements.",
				Suggestion:  "Increase the margin and padding between elements to improve readability.",
			},
		},
	}, nil
}
