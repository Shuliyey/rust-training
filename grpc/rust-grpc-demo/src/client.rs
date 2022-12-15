use futures::stream::iter;
use hello::say_client::SayClient;
use hello::SayRequest;

mod hello;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
// creating a channel ie connection to server
  let channel = tonic::transport::Channel::from_static("http://[::1]:50051")
  .connect()
  .await?;
  println!("calling hello.Say/Send");
// creating gRPC client from channel
  let mut client = SayClient::new(channel);
// creating a new Request
  let request = tonic::Request::new(
      SayRequest {
          name:String::from("anshul")
      },
  );
// sending request and waiting for response
  let response = client.send(request).await?.into_inner();
  println!("RESPONSE={:?}", response);

  println!("\ncalling hello.Say/SendStream");
  let request2 = tonic::Request::new(
    SayRequest {
        name:String::from("anshul")
    },
  );
  // now the response is stream
  let mut response2 = client.send_stream(request2).await?.into_inner();
  // listening to stream
  while let Some(res) = response2.message().await? {
      println!("NOTE = {:?}", res);
  }

  println!("\ncalling hello.Say/ReceiveStream");
  // creating a stream
  let request3 = tonic::Request::new(iter(vec![
    SayRequest {
        name: String::from("anshul"),
    },
    SayRequest {
        name: String::from("rahul"),
    },
    SayRequest {
        name: String::from("vijay"),
    },
  ]));
  // sending stream
  let response3 = client.receive_stream(request3).await?.into_inner();
  println!("RESPONSE=\n{}", response3.message);

  println!("\ncalling hello.Say/Bidirectional");
// creating a stream
  let request4 = tonic::Request::new(iter(vec![
    SayRequest {
        name: String::from("anshul2"),
    },
    SayRequest {
        name: String::from("rahul2"),
    },
    SayRequest {
        name: String::from("vijay2"),
    },
  ]));
  
  // calling rpc
  let mut response4 = client.bidirectional(request4).await?.into_inner();
// listening on the response stream
  while let Some(res) = response4.message().await? {
      println!("NOTE = {:?}", res);
  }

  Ok(())
}