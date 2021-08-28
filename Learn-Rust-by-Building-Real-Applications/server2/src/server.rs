use crate::http::{Request, Response, StatusCode, ParseError};
use std::convert::TryFrom;
use std::io::{Read, Write};
use std::net::TcpListener;

pub trait Handler {
  fn handle_request(&mut self, request: &Request) -> Response;

  fn handle_bad_request(&mut self, e: &ParseError) -> Response {
    println!("Failed to parse a request: {}", e);
    Response::new(StatusCode::BadRequest, None)
  }
}

pub struct Server {
  addr: String,   
}

//fn arr(a: &[u8]) {}

impl Server {
  pub fn new(addr: String) -> Self {
    Self { addr }
  }

  pub fn run(self, mut handler: impl Handler) {
    println!("Listening on {}", self.addr);

    let listener = TcpListener::bind(&self.addr).unwrap();

    loop {
      match listener.accept() {
        Ok((mut stream, _addr)) => {
          //let a = [1, 3, 3, 3, 2, 3];
          //arr(&a[1..3]);
          let mut buffer = [0; 1024];
          match stream.read(&mut buffer) {
            Ok(_) => {
              //println!("Received a request: {:?}", String::from_utf8_lossy(&buffer));

              let response = match Request::try_from(&buffer[..]) {
                Ok(request) => handler.handle_request(&request),
                  /*
                  dbg!(request);
                  Response::new(StatusCode::Ok, Some("<h1>it works</h1>".to_string()))
                  */
                Err(e) => handler.handle_bad_request(&e),
                  /*
                  println!("Failed to parse a request: {}", e);
                  Response::new(StatusCode::BadRequest, None)
                  */
              };

              if let Err(e) = response.send(&mut stream) {
                println!("Failed to send response: {}", e);
              }
            },
            Err(e) => println!("Failed ot read from a connection: {}", e),
          }
        }
        Err(e) => println!("Failed to establish a connection: {}", e),
      }
    }
  }
}