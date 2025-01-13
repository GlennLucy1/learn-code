from concurrent import futures

import grpc

import core
from pb import ocr_pb2_grpc
from pb.ocr_pb2 import DetectResponse
from pb.ocr_pb2_grpc import OCRServicer
from datetime import datetime

ocr = core.CoreOcr()

class OCRServer(OCRServicer):
    def Detect(self, request, context):
        current_time = datetime.now()
        formatted_time = current_time.strftime("%Y-%m-%d %H:%M:%S")
        b64_img = request.b64_img
        result = ocr.classification(b64_img)
        print("Detect at:", formatted_time, result, request.b64_img)
        print()
        return DetectResponse(response=list(result))

VALID_KEY = "123e4567-e89b-12d3-a456-426614174000"
class AuthInterceptor(grpc.ServerInterceptor):
    def intercept_service(self, continuation, handler):
        def new_handler(request, context):
            uuid_header = dict(context.invocation_metadata()).get('key')
            if uuid_header != VALID_KEY:
                context.abort(grpc.StatusCode.UNAUTHENTICATED, "Invalid Key")
            return continuation(request, context)
        return new_handler

def serve():
    port = "50000"
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=10),
    )

    ocr_pb2_grpc.add_OCRServicer_to_server(OCRServer(), server)
    server.add_insecure_port('[::]:' + port)

    server.start()
    print("Server started, listening on " + port)
    print()
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
