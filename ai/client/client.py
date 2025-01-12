import grpc

from pb.ocr_pb2 import DetectRequest
from pb.ocr_pb2_grpc import OCRStub


def run():
    with grpc.insecure_channel("localhost:50000") as channel:
        stub = OCRStub(channel)
        response = stub.Detect(DetectRequest(b64_img=""))
        print(response.response)

if __name__ == '__main__':
    run()