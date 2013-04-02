package goweixin

import (
	"bytes"
	"fmt"
	"io"
)

func MapToXmlString(m map[string]interface{}) string {
	buf := &bytes.Buffer{}
	for k, v := range m {
		io.WriteString(buf, fmt.Sprintf("<%s>", k))
		if v != nil {
			switch v.(type) {
			case int:
				io.WriteString(buf, fmt.Sprintf("%d", v))
			case int64:
				io.WriteString(buf, fmt.Sprintf("%d", v))
			case string:
				io.WriteString(buf, "<![CDATA["+v.(string)+"]]>")
			case map[string]interface{}:
				io.WriteString(buf, MapToXmlString(v.(map[string]interface{})))
			}
		}
		io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
	}
	return buf.String()
}
