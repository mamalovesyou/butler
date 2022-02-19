package airbyte

import "github.com/butlerhq/airbyte-client-go/airbyte"

const CheckConnectionStatusFailure = "failed"

// IsFailure return true if a CheckConnectionRead.Status is set to failed
func IsFailure(resp *airbyte.CheckConnectionRead) bool {
	return resp.Status == CheckConnectionStatusFailure
}
