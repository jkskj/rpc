import thriftpy2
from thriftpy2.rpc import make_server

test_thrift = thriftpy2.load("idl/test.thrift", module_name="test_thrift")


class Test(object):

    def test(self, req):
        print(req)
        return "pong"


server = make_server(test_thrift.test, Test(), '127.0.0.1', 6000)
server.serve()
