package response

type Body struct {
	Version           string                 `json:"version"`
	Response          Response               `json:"response"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
}

type Response struct {
	OutputSpeech     OutputSpeech `json:"outputSpeech,omitempty"`
	Reprompt         Reprompt     `json:"remprompt,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession,omitempty"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

type Reprompt struct {
	OutputSpeech OutputSpeech `json:"outputSpeech,omitempty"`
}

func NewText(text string) Body {
	return Body{
		Version: "0.1",
		Response: Response{
			OutputSpeech: OutputSpeech{
				Type: "PlainText",
				Text: text,
			},
		},
	}
}
