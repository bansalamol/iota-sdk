package functions

import (
	"encoding/json"

	"github.com/iota-uz/iota-sdk/modules/bichat/domain/entities/llm"

	"gorm.io/gorm"
)

type CompletionFunc func(args map[string]interface{}) (string, error)

func New() *ChatTools {
	return &ChatTools{}
}

func Default(db *gorm.DB) *ChatTools {
	return &ChatTools{
		Definitions: []ChatFunctionDefinition{
			NewGetSchema(db),
		},
	}
}

type ChatTools struct {
	Definitions []ChatFunctionDefinition
}

func (c *ChatTools) Add(def ChatFunctionDefinition) {
	c.Definitions = append(c.Definitions, def)
}

func (c *ChatTools) OpenAiTools() []llm.Tool {
	tools := make([]llm.Tool, 0, len(c.Definitions))
	for _, def := range c.Definitions {
		tools = append(tools, llm.Tool{
			Type: "function",
			Function: &llm.FunctionDefinition{
				Name:        def.Name(),
				Description: def.Description(),
				Parameters:  def.Arguments(),
			},
		})
	}
	return tools
}

func (c *ChatTools) Funcs() map[string]CompletionFunc {
	funcs := map[string]CompletionFunc{}
	for _, def := range c.Definitions {
		funcs[def.Name()] = def.Execute
	}
	return funcs
}

func (c *ChatTools) Call(name string, args string) (string, error) {
	if fn, ok := c.Funcs()[name]; ok {
		parsedArgs := map[string]interface{}{}
		if err := json.Unmarshal([]byte(args), &parsedArgs); err != nil {
			return "", err
		}
		return fn(parsedArgs)
	}
	return "", nil
}
