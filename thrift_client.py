import time

import thriftpy2
from thriftpy2.rpc import make_client

test_thrift = thriftpy2.load("idl/test.thrift", module_name="test_thrift")

client = make_client(test_thrift.test, '127.0.0.1', 7000)

while 1:
    print(client.test("ping"))
    time.sleep(1)
