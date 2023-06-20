package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		testName    string
		token       string
		expectError bool
	}{
		{
			testName:    "authorized",
			token:       "Bearer test_token",
			expectError: false,
		},
		{
			testName:    "no token provided",
			token:       "",
			expectError: true,
		},
		{
			testName:    "no bearer token",
			token:       "no-bearer token",
			expectError: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			req := events.APIGatewayCustomAuthorizerRequest{
				AuthorizationToken: tt.token,
				MethodArn:          "testARN",
			}

			_, err := handleRequest(context.TODO(), req)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
