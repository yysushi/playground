package main_test

import (
	"testing"
	"time"

	dtpb "github.com/googleapis/go-type-adapters/adapters"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestDateTime(t *testing.T) {
	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+09:00")
	tmpb, _ := dtpb.TimeToProtoDateTime(tm)
	v, _ := protojson.Marshal(tmpb)
	assert.Equal(t, `{"year":2006,"month":1,"day":2,"hours":15,"minutes":4,"seconds":5,"timeZone":{"id":"Local"}}`, string(v))
	assert.IsType(t, &datetime.DateTime_TimeZone{}, tmpb.TimeOffset)
}

func TestUnmarshal(t *testing.T) {
	b := []byte(`{"year":2006,"month":1,"day":2,"hours":15,"minutes":4,"seconds":5,"timeZone":{"id":"Local"}}`)
	var tmpb datetime.DateTime
	_ = protojson.Unmarshal(b, &tmpb)
	switch x := tmpb.TimeOffset.(type) {
	case *datetime.DateTime_TimeZone:
	default:
		t.Fatalf("unexpected type %T", x)
	}
}
