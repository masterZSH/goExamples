const process = require('process');
var PROTO_PATH = __dirname + '/../pb/helloworld/helloworld.proto';
console.log(PROTO_PATH);
var parseArgs = require('minimist');
var grpc = require('@grpc/grpc-js');
var grpc_xds = require('@grpc/grpc-js-xds');
grpc_xds.register();

// handle exception
process.on("uncaughtExpection", (err) => {
    console.log(err);
})


var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var hello_proto = grpc.loadPackageDefinition(packageDefinition).helloworld;

function main() {
  var argv = parseArgs(process.argv.slice(2), {
    string: 'target'
  });
  // 目标地址 grpc服务端
  var target;
  if (argv.target) {
    target = argv.target;
  } else {
    target = 'localhost:50051';
  }
  var client = new hello_proto.Greeter(target,
                                       grpc.credentials.createInsecure());

  // user 第一个参数传进来 默认值world
  var user;
  if (argv._.length > 0) {
    user = argv._[0]; 
  } else {
    user = 'world';
  }
  client.sayHello({name: user}, function(err, response) {
    if (err) throw err;
    console.log('Greeting:', response.message);
    client.close();
  });
}

main();