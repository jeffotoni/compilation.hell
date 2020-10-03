use std::io::prelude::*;
use std::net::TcpListener;
use std::net::TcpStream;

// Minimalist single-threaded http api
// documentation available at https://doc.rust-lang.org/book/ch20-01-single-threaded.html
fn handle_connection(mut stream: TcpStream) {
    let mut buffer = [0; 1024];
    stream.read(&mut buffer).unwrap();
    let get = b"GET / HTTP/1.1\r\n";
    let status_line = if buffer.starts_with(get) {
        "HTTP/1.1 200 OK\r\n\r\n"
    } else {
        "HTTP/1.1 404 NOT FOUND\r\n\r\n"
    };
    let response = format!("{}{}", status_line, "");
    stream.write(response.as_bytes()).unwrap();
    stream.flush().unwrap();
}

fn main() {
    let listener = TcpListener::bind("0.0.0.0:8080").unwrap();
    println!("Starting server at 0.0.0.0:8080");
    for stream in listener.incoming() {
        let stream = stream.unwrap();
        handle_connection(stream);
    }
}