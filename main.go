package main

import (
	"mime"
	"os"
	"path"
	"path/filepath"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucket(ctx, "website-bucket", &s3.BucketArgs{
			Website: s3.BucketWebsiteArgs{
				IndexDocument: pulumi.String("index.html"),
			},
		})
		if err != nil {
			return err
		}

		websiteDir := "web"

		files, err := os.ReadDir(websiteDir)
		if err != nil {
			return err
		}

		for _, item := range files {
			name := item.Name()
			filePath := filepath.Join(websiteDir, name)
			if _, err := s3.NewBucketObject(ctx, name, &s3.BucketObjectArgs{
				Acl:         pulumi.String("public-read"),
				Bucket:      bucket.ID(),
				Source:      pulumi.NewFileAsset(filePath),
				ContentType: pulumi.String(mime.TypeByExtension(path.Ext(name))),
			}); err != nil {
				return err
			}
		}

		ctx.Export("bucketName", bucket.ID())
		ctx.Export("bucketEndpoint", pulumi.Sprintf("http://%s", bucket.WebsiteEndpoint))
		return nil
	})

}
