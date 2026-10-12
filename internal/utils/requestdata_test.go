package honokautils

import (
	"mime"
	"testing"
)

func TestRecoverMultipartRequestDataWithUnquotedSlashBoundary(t *testing.T) {
	const boundary = "R+7uHJNv7+Nl43xw/ecH"
	if _, _, err := mime.ParseMediaType("multipart/form-data; boundary=" + boundary); err == nil {
		t.Fatal("expected Go's standard parser to reject the unquoted slash")
	}
	body := []byte("--" + boundary + "\r\nContent-Disposition: form-data; name=\"request_data\"\r\n\r\n{\"module\":\"api\"}\r\n--" + boundary + "--\r\n")
	got := RecoverMultipartRequestData("multipart/form-data; boundary="+boundary, body)
	if got != `{"module":"api"}` {
		t.Fatalf("got %q", got)
	}
	for name, header := range map[string]string{
		"different boundary": "multipart/form-data; boundary=wrong/boundary",
		"valid boundary":     "multipart/form-data; boundary=valid-boundary",
		"other media type":   "application/json",
	} {
		t.Run(name, func(t *testing.T) {
			if got := RecoverMultipartRequestData(header, body); got != "" {
				t.Fatalf("unexpected recovery: %q", got)
			}
		})
	}
}
