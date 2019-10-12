// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package httpstate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleURL(t *testing.T) {
	t.Parallel()
	assert.Equal(t,
		"https://app.pulumi.com/pulumi-bot/my-stack",
		cloudConsoleURL("https://api.pulumi.com", "pulumi-bot", "my-stack"))

	assert.Equal(t,
		"http://app.pulumi.example.com/pulumi-bot/my-stack",
		cloudConsoleURL("http://api.pulumi.example.com", "pulumi-bot", "my-stack"))

	assert.Equal(t,
		"http://localhost:3000/pulumi-bot/my-stack",
		cloudConsoleURL("http://localhost:8080", "pulumi-bot", "my-stack"))

	assert.Equal(t, "", cloudConsoleURL("https://example.com", "pulumi-bot", "my-stack"))

	assert.Equal(t, "", cloudConsoleURL("not-even-a-rea-url", "pulumi-bot", "my-stack"))
}
