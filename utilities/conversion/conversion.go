package leibniz_conversion

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func StringToTimestamp(datetimeStr string, strFormat string) (*timestamppb.Timestamp, error) {
	if len(datetimeStr) == 0 {
		return nil, fmt.Errorf("empty datetimeStr value")
	}
	if len(strFormat) == 0 {
		return nil, fmt.Errorf("empty strFormat value")
	}

	t, err := time.Parse(strFormat, datetimeStr)

	if err != nil {
		return nil, fmt.Errorf("while converting string to time.time: %w", err)
	}

	return timestamppb.New(t), nil
}
