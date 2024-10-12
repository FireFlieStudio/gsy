package TencentCos

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type ProgressBar struct{}

func (p *ProgressBar) ProgressChangedCallback(event *cos.ProgressEvent) {
	switch event.EventType {
	case cos.ProgressStartedEvent:
		fmt.Printf("传输中开始		[已下载/总大小: %s/%s]",
			humanize.Bytes(uint64(event.ConsumedBytes)), humanize.Bytes(uint64(event.TotalBytes)))
	case cos.ProgressDataEvent:
		fmt.Printf("\r传输中		[已下载/总大小: %s/%s, %d%%]",
			humanize.Bytes(uint64(event.ConsumedBytes)), humanize.Bytes(uint64(event.TotalBytes)), event.ConsumedBytes*100/event.TotalBytes)
	case cos.ProgressCompletedEvent:
		fmt.Printf("\r传输中完成		[已下载/总大小: %s/%s]\n",
			humanize.Bytes(uint64(event.ConsumedBytes)), humanize.Bytes(uint64(event.TotalBytes)))
	case cos.ProgressFailedEvent:
		fmt.Printf("\n传输中失败		[已下载/总大小: %s/%s] [Error: %v]\n",
			humanize.Bytes(uint64(event.ConsumedBytes)), humanize.Bytes(uint64(event.TotalBytes)), event.Err)
	default:
		fmt.Printf("Progress Changed Error: unknown progress event type\n")
	}
}
