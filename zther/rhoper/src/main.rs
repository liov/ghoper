use actix_web::{server::{self,HttpServer}, App, HttpRequest};
use std::env;


fn main() {

   let args: Vec<String> = env::args().collect();

    ::std::env::set_var("RUST_LOG", "hoper=info");
    ::std::env::set_var("RUST_BACKTRACE", "1");

    let sys = actix::System::new("hoper");

    server::new( move || router::app().unwrap())
        .bind("127.0.0.1:8000").unwrap()
        .shutdown_timeout(2)
        .start();

    sys.run();
    env_logger::init();
}