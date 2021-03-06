package client // import "github.com/docker/docker/client"

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestDistributionInspectUnsupported(t *testing.T) {
	client := &Client{
		version: "1.29",
		client:  &http.Client{},
	}
	_, err := client.DistributionInspect(context.Background(), "foobar:1.0", "")
	assert.EqualError(t, err, `"distribution inspect" requires API version 1.30, but the Docker daemon API version is 1.29`)
}

func TestDistributionInspectWithEmptyID(t *testing.T) {
	client := &Client{
		client: newMockClient(func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("should not make request")
		}),
	}
	_, err := client.DistributionInspect(context.Background(), "", "")
	if !IsErrNotFound(err) {
		t.Fatalf("Expected NotFoundError, got %v", err)
	}
}
