package likenft

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/util/fileresolver"
	"github.com/likecoin/like-migration-backend/pkg/util/fileupload"
	"github.com/likecoin/like-migration-backend/pkg/util/image"
	imageutil "github.com/likecoin/like-migration-backend/pkg/util/image"
)

// go run ./cmd/cli likenft convert-and-upload-assets ./test-images/1.webp
var convertAndUploadAssetsCmd = &cobra.Command{
	Use:   "convert-and-upload-assets srcpath",
	Short: "Get ISCN Record",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}
		srcPath := args[0]

		ctx := cmd.Context()
		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)
		logger := slog.New(slog.Default().Handler())

		httpClient := &http.Client{}
		fileResolver := fileresolver.MakeFileResolver(
			httpClient,
			envCfg.IpfsHTTPBaseURL,
			envCfg.ArweaveHTTPBaseURL,
		)
		uploader, err := fileupload.MakeW3Client(
			logger,
			httpClient,
			envCfg.W3SpaceDID,
			envCfg.W3UcanDIDPrivateKey,
			envCfg.W3UcanProofPath,
		)
		if err != nil {
			panic(err)
		}

		srcData, err := fileResolver.Resolve(ctx, srcPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("len(srcData): %d\n", len(srcData))
		srcDataUploadedLink, err := uploader.Upload(ctx, srcData)
		if err != nil {
			panic(err)
		}
		fmt.Println(srcDataUploadedLink)

		tempFile, err := os.CreateTemp("", "")
		if err != nil {
			panic(err)
		}
		defer os.Remove(tempFile.Name())
		logger = logger.With("tempFile", tempFile.Name())
		_, err = tempFile.Write(srcData)
		if err != nil {
			panic(err)
		}
		tempFile.Close()

		targetPath, err := imageutil.DefaultMimeTypeConversion.GetDefaultConvertedPath(tempFile.Name())
		if err != nil {
			panic(err)
		}
		logger = logger.With("targetPath", targetPath)
		defer os.Remove(targetPath)

		err = image.Convert(tempFile.Name(), targetPath)
		if err != nil {
			panic(err)
		}

		dataToBeUploaded, err := fileResolver.Resolve(ctx, targetPath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("len(dataToBeUploaded): %d\n", len(dataToBeUploaded))

		logger.Info("Uploading to w3...")
		uploadedLink, err := uploader.Upload(ctx, dataToBeUploaded)
		if err != nil {
			panic(err)
		}

		fmt.Println(uploadedLink)
	},
}

func init() {
	LikeNFTCmd.AddCommand(convertAndUploadAssetsCmd)
}
