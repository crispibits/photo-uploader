package main

import (
	"fmt"

	"github.com/crispibits/photo-uploader/pkg/client"
	"github.com/crispibits/photo-uploader/pkg/config"
)

func main() {
	config, err := config.ReadOrCreate("")
	if err != nil {
		panic(err)
	}
	client, err := client.NewGCSClient(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", client.Config().GCS.Bucket)

	//checkCardAccess()
	//makeDirectories()
	//copyFiles()
	//verifyFiles()
	//wipeCard()
}
