package instrumentation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLExtraction(t *testing.T) {
	testCases := map[string]string{
		// Cluster-wide core objects
		"/api/v1/nodes":                    "core/v1/nodes collection",
		"/api/v1/nodes/the-node":           "core/v1/nodes object",
		"/api/v1/nodes/the-node/status":    "core/v1/nodes object status subresource",
		"/api/v1/namespaces":               "core/v1/namespaces collection",
		"/api/v1/namespaces/the-namespace": "core/v1/namespaces object",
		// Namespaced core objects
		"/api/v1/namespaces/the-namespace/pods":                "core/v1/pods collection",
		"/api/v1/namespaces/the-namespace/pods/the-pod":        "core/v1/pods object",
		"/api/v1/namespaces/the-namespace/pods/the-pod/status": "core/v1/pods object status subresource",
		// Cluster-wide non-core objects
		"/apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations":                   "admissionregistration.k8s.io/v1/validatingwebhookconfigurations collection",
		"/apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/the-config":        "admissionregistration.k8s.io/v1/validatingwebhookconfigurations object",
		"/apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/the-config/status": "admissionregistration.k8s.io/v1/validatingwebhookconfigurations object status subresource",
		// Namespaced non-core objects
		"/apis/apps/v1/namespaces/the-namespace/statefulsets":                        "apps/v1/statefulsets collection",
		"/apis/apps/v1/namespaces/the-namespace/statefulsets/the-statefulset":        "apps/v1/statefulsets object",
		"/apis/apps/v1/namespaces/the-namespace/statefulsets/the-statefulset/status": "apps/v1/statefulsets object status subresource",

		// Invalid paths
		// These cases should never happen given we're using the official client library and it always uses the expected format, but we
		// should still do something sensible if these do happen.
		"":                "",
		"/":               "/",
		"/something-else": "/something-else",
		"/api":            "/api",
		"/api/v1":         "/api/v1",
		"/apis":           "/apis",
		"/apis/apps":      "/apis/apps",
		"/apis/apps/v1":   "/apis/apps/v1",
	}

	for input, expectedDescription := range testCases {
		t.Run(input, func(t *testing.T) {
			actualDescription := urlToResourceDescription(input)
			require.Equal(t, expectedDescription, actualDescription)
		})
	}
}
