package fileupload

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	"github.com/likecoin/like-migration-backend/pkg/util/fileupload"
)

var uploadCmd = &cobra.Command{
	Use:   "upload file",
	Short: "Convert image format",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			_ = cmd.Usage()
			return
		}

		logger := slog.New(slog.Default().Handler())
		ctx := cmd.Context()
		envCfg := ctx.Value(config.ContextKey).(*config.EnvConfig)

		httpClient := &http.Client{}

		w3, err := fileupload.MakeW3Client(
			logger,
			httpClient,
			envCfg.W3SpaceDID,
			envCfg.W3UcanDIDPrivateKey,
			envCfg.W3UcanProofPath,
		)

		if err != nil {
			panic(err)
		}

		fileBytes, err := os.ReadFile(args[0])
		if err != nil {
			panic(err)
		}

		link, err := w3.Upload(ctx, fileBytes)
		if err != nil {
			panic(err)
		}

		fmt.Println(link)
	},
}

func init() {
	FileUploadCmd.AddCommand(uploadCmd)
}
