package ctrl

import (
	"context"
	pb "github.com/GlennLucy1/learn-ai222/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var OcrAddr = ""

func Detect(ctx context.Context, b64img string) ([]string, error) {
	conn, err := grpc.NewClient(OcrAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := pb.NewOCRClient(conn)

	r, err := c.Detect(ctx, &pb.DetectRequest{B64Img: b64img})
	if err != nil {
		return nil, err
	}

	return r.GetResponse(), nil
}
