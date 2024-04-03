package config

type PromptSuggestion struct {
	Title   []string `json:"title"`
	Content string   `json:"content"`
}

func GetDefaultPromptSuggestions() []PromptSuggestion {
	return []PromptSuggestion{
		{
			Title:   []string{"Help me study", "vocabulary for a college entrance exam"},
			Content: "Help me study vocabulary: write a sentence for me to fill in the blank, and I'll try to pick the correct option.",
		},
		{
			Title:   []string{"Give me ideas", "for what to do with my kids' art"},
			Content: "What are 5 creative things I could do with my kids' art? I don't want to throw them away, but it's also so much clutter.",
		},
		{
			Title:   []string{"Tell me a fun fact", "about the Roman Empire"},
			Content: "Tell me a random fun fact about the Roman Empire",
		},
		{
			Title:   []string{"Show me a code snippet", "of a website's sticky header"},
			Content: "Show me a code snippet of a website's sticky header in CSS and JavaScript.",
		},
	}

	//str, err := json.Marshal(suggestions)
	//if err != nil {
	//	panic(err)
	//}
	//
	//return string(str)
}
