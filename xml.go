package goweixin

import (
	"bytes"
	"fmt"
	"io"
)

func MapToXmlString(m map[string]interface{}) string {
	buf := &bytes.Buffer{}
	for k, v := range m {

		if v != nil {
			switch v.(type) {
			case int:
				io.WriteString(buf, fmt.Sprintf("<%s>", k))
				io.WriteString(buf, fmt.Sprintf("%d", v))
				io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
			case int64:
				io.WriteString(buf, fmt.Sprintf("<%s>", k))
				io.WriteString(buf, fmt.Sprintf("%d", v))
				io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
			case string:
				io.WriteString(buf, fmt.Sprintf("<%s>", k))
				io.WriteString(buf, "<![CDATA["+v.(string)+"]]>")
				io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
			case map[string]interface{}:
				io.WriteString(buf, fmt.Sprintf("<%s>", k))
				io.WriteString(buf, MapToXmlString(v.(map[string]interface{})))
				io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
			case []interface{}:
				for _, t := range v.([]interface{}) {
					switch t.(type) {
					case map[string]interface{}:
						io.WriteString(buf, fmt.Sprintf("<%s>", k))
						io.WriteString(buf, MapToXmlString(t.(map[string]interface{})))
						io.WriteString(buf, fmt.Sprintf("</%s>\n", k))
					}
				}
			}
		}

	}
	return buf.String()
}
