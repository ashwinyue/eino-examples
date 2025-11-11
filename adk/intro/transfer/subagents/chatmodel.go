/*
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package subagents

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"

	"github.com/cloudwego/eino-examples/adk/common/model"
)

type GetWeatherInput struct {
	City string `json:"city"`
}

func NewWeatherAgent() adk.Agent {
	weatherTool, err := utils.InferTool(
		"get_weather",
		"Get the current weather for a specified city.", // English description
		func(ctx context.Context, input *GetWeatherInput) (string, error) {
			return fmt.Sprintf(`The temperature in %s is 25°C`, input.City), nil
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	a, err := adk.NewChatModelAgent(context.Background(), &adk.ChatModelAgentConfig{
		Name:        "WeatherAgent",
		Description: "This agent can get the current weather for a given city.",
		Instruction: `Your sole purpose is to get the current weather for a given city using the 'get_weather' tool.
After calling the tool, report the result directly to the user.`,
		Model: model.NewChatModel(),
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []tool.BaseTool{weatherTool},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func NewChatAgent() adk.Agent {
	a, err := adk.NewChatModelAgent(context.Background(), &adk.ChatModelAgentConfig{
		Name:        "ChatAgent",
		Description: "A general-purpose agent for handling conversational chat.", // English description
		Instruction: `You are a friendly conversational assistant.
Your role is to handle general chat and answer questions not related to any specific tool-based tasks.`,
		Model: model.NewChatModel(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func NewRouterAgent() adk.Agent {
	a, err := adk.NewChatModelAgent(context.Background(), &adk.ChatModelAgentConfig{
		Name:        "RouterAgent",
		Description: "A manual router that transfers tasks to other expert agents.",
		Instruction: `You are an intelligent task router.
Your responsibility is to analyze the user's request and delegate it to the most appropriate expert agent.
If no agent can handle the task, just inform the user that it cannot be handled.`,
		Model: model.NewChatModel(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return a
}
