package main

import (
	"fmt"

	"github.com/huaweicloud/golangsdk/openstack/obs"
)

func main() {
}
func test() {
	obsClient, _ := obs.New("UNBDJ59E2CSBPBYO44X6", "NMbichpkC2StzfxFU8XRDKaXAKWq34qPa386bZMP", "https://obs.cn-north-4.myhuaweicloud.com")
	/* input := &obs.ListBucketsInput{}
	output, err := obsClient.ListBuckets(input)

	if err != nil {
		fmt.Println(err)
		return
	}
	//列举存储桶
	for _, v := range output.Buckets {
		fmt.Println(v.Name)
	} */
	//列举对象
	{
		//	inputobj := &obs.GetBucketMetadataInput{}
		//inputobj.Bucket = "umb-monitor-store"
		output, err := obsClient.GetBucketStorageInfo("umb-monitor-store")
		if err == nil {
			fmt.Println(float64(output.Size)/float64((1024*1024)), output.ObjectNumber)
			/* 	fmt.Println(output.BaseModel)
			for _, val := range output.ResponseHeaders {
				fmt.Printf("Key:%v\n", val)
			} */
		}
	}

	/* currentTime := time.Now()
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("ShortDay : ", currentTime.Format("Mon,02 Jan 2006 15:04:05")) */
}
