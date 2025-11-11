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

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"

	"github.com/cloudwego/eino-examples/adk/common/model"
	"github.com/cloudwego/eino-examples/adk/common/prints"
)

func main() {
	ctx := context.Background()

	toolA, err := utils.InferTool("tool_a", "设置用户名", toolAFn)
	if err != nil {
		log.Fatalf("InferTool failed, error: %v", err)
	}

	toolB, err := utils.InferTool("tool_b", "设置用户年龄", toolBFn)
	if err != nil {
		log.Fatalf("InferTool failed, error: %v", err)
	}

	a, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "ChatModelAgent",
		Description: "Chat model agent",
		Instruction: "You are a chat model agent, first call tool_a, then call tool_b",
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []tool.BaseTool{
					toolA,
					toolB,
				},
			},
		},
		Model: model.NewChatModel(),
	})
	if err != nil {
		log.Fatalf("NewChatModelAgent failed, error: %v", err)
	}

	r := adk.NewRunner(ctx, adk.RunnerConfig{
		Agent: a,
	})

	iter := r.Query(ctx, "My name is Alice and I am 18 years old")
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}

		prints.Event(event)
	}
}

type ToolAInput struct {
	Name string `json:"input" jsonschema_description:"User name"`
}

func toolAFn(ctx context.Context, in *ToolAInput) (string, error) {

	adk.AddSessionValue(ctx, "user-name", in.Name)
	return in.Name, nil
}

type ToolBInput struct {
	Age int `json:"input" jsonschema_description:"User age"`
}

func toolBFn(ctx context.Context, in *ToolBInput) (string, error) {
	adk.AddSessionValue(ctx, "user-age", in.Age)
	userName, _ := adk.GetSessionValue(ctx, "user-name")
	return fmt.Sprintf("User name: %v, User age: %v", userName, in.Age), nil
}
