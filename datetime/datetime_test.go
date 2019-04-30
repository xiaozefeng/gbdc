package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestCurrentDateTimeStr(t *testing.T) {
	fmt.Println(CurrentDateTimeStr())
	fmt.Println(CurrentDateStr())
	fmt.Println(CurrentTimeStr())
	fmt.Println(time.Now().Format(fmt.Sprintf("%s-%s-%s %s:%s:%s", Year,
		Month,
		Day,
		Hour,
		Minute,
		Second,
	)))
}
