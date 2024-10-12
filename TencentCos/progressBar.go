package TencentCos

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gsync/utils"
)

type ProgressBar struct{}

func (p *ProgressBar) ProgressChangedCallback(event *cos.ProgressEvent) {
	switch event.EventType {
	case cos.ProgressStartedEvent:
		fmt.Printf("传输开始		[已传输/总大小: %s/%s]",
			humanize.Bytes(uint64(event.ConsumedBytes)), humanize.Bytes(uint64(event.TotalBytes)))
	case cos.ProgressDataEvent:
		utils.Bar(event.ConsumedBytes, event.TotalBytes)
	case cos.ProgressCompletedEvent:
		utils.CloseBar()
	case cos.ProgressFailedEvent:
		fmt.Printf("\n传输中失败		[已传输/总大小: %s/%s] [Error: %v]\n",
			humanize.Bytes(uint64(event.ConsumedBytes)), humanize.Bytes(uint64(event.TotalBytes)), event.Err)
	default:
		fmt.Printf("Progress Changed Error: unknown progress event type\n")
	}
}
