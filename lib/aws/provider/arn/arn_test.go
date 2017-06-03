// Licensed to Pulumi Corporation ("Pulumi") under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// Pulumi licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package arn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParse(t *testing.T) {
	t.Parallel()

	// basic ARN parsing/extraction:
	a1 := New("ec2", "us-west-2", "123456789", "instance:i-0011223344")
	assert.Equal(t, "arn:aws:ec2:us-west-2:123456789:instance:i-0011223344", string(a1))
	a1p, err := a1.Parse()
	assert.Nil(t, err)
	assert.Equal(t, "aws", a1p.Partition)
	assert.Equal(t, "ec2", a1p.Service)
	assert.Equal(t, "us-west-2", a1p.Region)
	assert.Equal(t, "123456789", a1p.AccountID)
	assert.Equal(t, "instance:i-0011223344", a1p.Resource)
	a1prn1, err := a1.ParseResourceName()
	assert.Nil(t, err)
	assert.Equal(t, "i-0011223344", a1prn1)
	assert.Equal(t, "instance", a1p.ResourceType())
	assert.Equal(t, "i-0011223344", a1p.ResourceName())

	// basic ARN parsing/extraction (variant #2: "/" delimiter for resource)
	a2 := New("elasticloadbalancing", "us-west-2", "123456789", "loadbalancer/mylb-123")
	assert.Equal(t, "arn:aws:elasticloadbalancing:us-west-2:123456789:loadbalancer/mylb-123", string(a2))
	a2p, err := a2.Parse()
	assert.Nil(t, err)
	assert.Equal(t, "aws", a2p.Partition)
	assert.Equal(t, "elasticloadbalancing", a2p.Service)
	assert.Equal(t, "us-west-2", a2p.Region)
	assert.Equal(t, "123456789", a2p.AccountID)
	assert.Equal(t, "loadbalancer/mylb-123", a2p.Resource)
	a2prn1, err := a2.ParseResourceName()
	assert.Nil(t, err)
	assert.Equal(t, "mylb-123", a2prn1)
	assert.Equal(t, "loadbalancer", a2p.ResourceType())
	assert.Equal(t, "mylb-123", a2p.ResourceName())

	// basic ARN parsing/extraction (variant #3: "/" delimited resource name pairs)
	a3 := New("elasticbeanstalk", "us-east-2", "123456789", "environment/My App/MyEnvironment")
	assert.Equal(t, "arn:aws:elasticbeanstalk:us-east-2:123456789:environment/My App/MyEnvironment", string(a3))
	a3p, err := a3.Parse()
	assert.Nil(t, err)
	assert.Equal(t, "aws", a3p.Partition)
	assert.Equal(t, "elasticbeanstalk", a3p.Service)
	assert.Equal(t, "us-east-2", a3p.Region)
	assert.Equal(t, "123456789", a3p.AccountID)
	assert.Equal(t, "environment/My App/MyEnvironment", a3p.Resource)
	a3prn1, err := a3.ParseResourceName()
	assert.Nil(t, err)
	assert.Equal(t, "My App/MyEnvironment", a3prn1)
	assert.Equal(t, "environment", a3p.ResourceType())
	assert.Equal(t, "My App/MyEnvironment", a3p.ResourceName())
	a3name1, a3name2, err := a3.ParseResourceNamePair()
	assert.Nil(t, err)
	assert.Equal(t, "My App", a3name1)
	assert.Equal(t, "MyEnvironment", a3name2)
	a3name1b, a3name2b := a3p.ResourceNamePair()
	assert.Equal(t, "My App", a3name1b)
	assert.Equal(t, "MyEnvironment", a3name2b)
}
