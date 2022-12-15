use tokio::sync::mpsc;
use tonic::{transport::Server, Request, Response, Status};
use hello::say_server::{Say, SayServer};
use hello::{SayResponse, SayRequest};
use resize::resize_server::{Resize, ResizeServer};
use resize::{ResizeResponse, ResizeRequest};
use image::io::Reader as ImageReader;
use image::imageops::FilterType;
use sha2::{Sha256, Digest};
use lru::LruCache;
//use std::io::Cursor;

mod hello;
mod resize;

// defining a struct for our service
#[derive(Default)]
pub struct MySay {}

// implementing rpc for service defined in .proto
#[tonic::async_trait]
impl Say for MySay {
  // Specify the output of rpc call
  type SendStreamStream = mpsc::Receiver<Result<SayResponse,Status>>;
  // defining return stream
  type BidirectionalStream = mpsc::Receiver<Result<SayResponse, Status>>;
  
// our rpc impelemented as function
  async fn send(&self,request:Request<SayRequest>)->Result<Response<SayResponse>,Status>{
// returning a response as SayResponse message as defined in .proto
    Ok(Response::new(SayResponse{
// reading data from request which is awrapper around our SayRequest message defined in .proto
      message:format!("hello {}",request.get_ref().name),
    }))
  }

  async fn send_stream(
    &self,
    request: Request<SayRequest>,
  ) -> Result<Response<Self::SendStreamStream>, Status> {
// creating a queue or channel
    let (mut tx, rx) = mpsc::channel(4);
// creating a new task
    tokio::spawn(async move {
// looping and sending our response using stream
      for _ in 0..4{
// sending response to our channel
        tx.send(Ok(SayResponse {
            message: format!("hello"),
        }))
        .await;
      }
    });
// returning our reciever so that tonic can listen on reciever and send the response to client
    Ok(Response::new(rx))
  }

// create a new rpc to receive a stream
  async fn receive_stream(
    &self,
    request: Request<tonic::Streaming<SayRequest>>,
  ) -> Result<Response<SayResponse>, Status> {
// converting request into stream
    let mut stream = request.into_inner();
    let mut message = String::from("");
// listening on stream
    while let Some(req) = stream.message().await? {
      message.push_str(&format!("Hello {}\n", req.name))
    }
// returning response
    Ok(Response::new(SayResponse { message }))
  }

  async fn bidirectional(
    &self,
    request: Request<tonic::Streaming<SayRequest>>,
  ) -> Result<Response<Self::BidirectionalStream>, Status> {
// converting request in stream
    let mut streamer = request.into_inner();
// creating queue
    let (mut tx, rx) = mpsc::channel(4);
    tokio::spawn(async move {
// listening on request stream
      while let Some(req) = streamer.message().await.unwrap(){
// sending data as soon it is available
        tx.send(Ok(SayResponse {
            message: format!("hello {}", req.name),
        }))
        .await;
      }
    });
// returning stream as receiver
    Ok(Response::new(rx))
  }
}

// defining a struct for our service
pub struct ResizeService {
   max_width: u32,
   max_height: u32,
   stream_num: u32,
   cache: LruCache<String, Vec<u8>>,
}

impl Default for ResizeService {
  fn default() -> Self {
    Self {
      max_height: 2000,
      max_width: 2000,
      stream_num: 518400,
      cache: LruCache::new(10),    
    }
  }
}

// implementing rpc for service defined in .proto
#[tonic::async_trait]
impl Resize for ResizeService {
  // Specify the output of rpc call
  // type SendStreamStream = mpsc::Receiver<Result<ResizeResponse,Status>>;
  // defining return stream
  // type BidirectionalStream = mpsc::Receiver<Result<ResizeResponse, Status>>;
  
// our rpc impelemented as function
  async fn send(&self,request:Request<ResizeRequest>)->Result<Response<ResizeResponse>,Status>{
// returning a response as ResizeResponse message as defined in .proto
    let req = &request.get_ref();

    if req.chunk.len() as u32 > self.stream_num {
      if req.width * req.height > self.stream_num {
        return Err(Status::invalid_argument(format!("total image size {} > {} and resize dimension is {}x{} > {}, please use endpoint resize.Resize/Bidirectional", req.chunk.len(), self.stream_num, req.width, req.height, self.stream_num)));
      }

      return Err(Status::invalid_argument(format!("total image size {} > {}, please use endpoint resize.Resize/SendStream or resize.Resize/Bidirectional", req.chunk.len(), self.stream_num)));
    }

    match req.width * req.height {
      num if num <= self.stream_num => {
        let chunk = &req.chunk;
        let mut hasher = Sha256::new();
        hasher.update(chunk);
        let hash = hasher.finalize();
        println!("sha: {:x}", hash);

        let img = image::load_from_memory(chunk).unwrap();
        println!("resizing {} to {}x{} using {:?}", req.name, req.width, req.height, FilterType::Gaussian);
        // let scaled = img.resize(req.width, req.height, FilterType::Nearest);
        let scaled = img.resize_exact(req.width, req.height, FilterType::Gaussian);
        let name = format!("{}x{}-{}", req.width, req.height, req.name);

        //scaled.save(name.clone()).unwrap();
        let mut resized_chunk = Vec::new();
        scaled.write_to(&mut resized_chunk, image::ImageOutputFormat::from(image::ImageFormat::Jpeg)).unwrap();

        Ok(Response::new(ResizeResponse{
        // reading data from request which is awrapper around our ResizeRequest message defined in .proto
          name: name.clone(),
          original: chunk.clone(),
          chunk: resized_chunk,
          width: req.width,
          height: req.height,
        }))
      }
      _ => Err(Status::invalid_argument(format!("resize dimension is {}x{} > {}, please use stream endpoint resize.Resize/ReceiveStream", req.width, req.height, self.stream_num))),
    }
  }

/*

  async fn send_stream(
    &self,
    request: Request<ResizeRequest>,
  ) -> Result<Response<Self::SendStreamStream>, Status> {
// creating a queue or channel
    let (mut tx, rx) = mpsc::channel(4);
// creating a new task
    tokio::spawn(async move {
// looping and sending our response using stream
      for _ in 0..4{
// sending response to our channel
        tx.send(Ok(ResizeResponse {
            chunk: format!("hello"),
        }))
        .await;
      }
    });
// returning our reciever so that tonic can listen on reciever and send the response to client
    Ok(Response::new(rx))
  }

// create a new rpc to receive a stream
  async fn receive_stream(
    &self,
    request: Request<tonic::Streaming<ResizeRequest>>,
  ) -> Result<Response<ResizeResponse>, Status> {
// converting request into stream
    let mut stream = request.into_inner();
    let mut message = String::from("");
// listening on stream
    while let Some(req) = stream.message().await? {
      message.push_str(&format!("Hello {}\n", req.name))
    }
// returning response
    Ok(Response::new(ResizeResponse { message }))
  }

  async fn bidirectional(
    &self,
    request: Request<tonic::Streaming<ResizeRequest>>,
  ) -> Result<Response<Self::BidirectionalStream>, Status> {
// converting request in stream
    let mut streamer = request.into_inner();
// creating queue
    let (mut tx, rx) = mpsc::channel(4);
    tokio::spawn(async move {
// listening on request stream
      while let Some(req) = streamer.message().await.unwrap(){
// sending data as soon it is available
        tx.send(Ok(ResizeResponse {
            message: format!("hello {}", req.name),
        }))
        .await;
      }
    });
// returning stream as receiver
    Ok(Response::new(rx))
  }
*/
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
// defining address for our service
  let addr = "[::1]:50051".parse().unwrap();
// creating a service
  let say = MySay::default();
  let image_resize = ResizeService::default(); 
  println!("Server listening on {}", addr);
// adding our service to our server.
  Server::builder()
    .add_service(SayServer::new(say))
    .add_service(ResizeServer::new(image_resize))
    .serve(addr)
    .await?;
  Ok(())
}