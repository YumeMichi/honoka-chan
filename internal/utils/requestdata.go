package honokautils

import (
	"bytes"
	"mime"
	"mime/multipart"
	"strings"
)

// RecoverMultipartRequestData accepts the client's unquoted '/' boundary only
// when the multipart body uses exactly the boundary supplied in Content-Type.
func RecoverMultipartRequestData(contentType string, body []byte) string {
	const prefix = "multipart/form-data; boundary="
	if !strings.HasPrefix(contentType, prefix) || len(body) == 0 {
		return ""
	}
	if _, _, err := mime.ParseMediaType(contentType); err == nil {
		return ""
	}
	boundary := strings.TrimSpace(strings.TrimPrefix(contentType, prefix))
	if !strings.Contains(boundary, "/") || strings.ContainsAny(boundary, "\";\r\n") {
		return ""
	}
	if !bytes.HasPrefix(body, []byte("--"+boundary+"\r\n")) {
		return ""
	}
	form, err := multipart.NewReader(bytes.NewReader(body), boundary).ReadForm(1 << 20)
	if err != nil {
		return ""
	}
	defer form.RemoveAll()
	if values := form.Value["request_data"]; len(values) > 0 {
		return values[0]
	}
	return ""
}
