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
	"log"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type BookSearchInput struct {
	Genre     string `json:"genre" jsonschema_description:"Preferred book type,enum=fiction,enum=sci-fi,enum=mystery,enum=biography,enum=business"`
	MaxPages  int    `json:"max_pages" jsonschema_description:"Maximum page limit (0 means no limit)"`
	MinRating int    `json:"min_rating" jsonschema_description:"Minimum user rating (0-5 scale)"`
}

type BookSearchOutput struct {
	Books []string
}

func NewBookRecommender() tool.InvokableTool {
	bookSearchTool, err := utils.InferTool("search_book", "Search for books based on user preferences",
		func(ctx context.Context, input *BookSearchInput) (output *BookSearchOutput, err error) {
			// Search code
			// ...
			return &BookSearchOutput{Books: []string{"Kono Subarashii Sekai ni Shukufuku o!"}}, nil
		},
	)
	if err != nil {
		log.Fatalf("Failed to create book search tool: %v", err)
	}
	return bookSearchTool
}
