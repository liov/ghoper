#![allow(unused_imports)]

mod handler;
mod models;

mod router;

use actix_web::{App, Either, web, http, HttpRequest, HttpResponse,
                HttpServer,
                middleware::{
                    Logger,
                },
                error::{self, Error}, Responder, dev::*};
use actix_session::{CookieSession, Session};
use actix::prelude::*;
use chrono::Duration;
use std::cell::Cell;
use core::borrow::Borrow;
#[macro_use]
extern crate serde_derive;

fn main() -> std::io::Result<()> {

    std::env::set_var("RUST_LOG", "hoper=info");
    std::env::set_var("RUST_BACKTRACE", "1");
    env_logger::init();


    //let sys = actix_rt::System::new("hoper");
    HttpServer::new( || {
        let secret: String =
            std::env::var("SECRET_KEY").unwrap_or_else(|_| "0123".repeat(8));
        let domain: String =
            std::env::var("DOMAIN").unwrap_or_else(|_| "localhost".to_string());

        App::new().data(handler::AppState { counter: Cell::new(0usize) })
            .wrap(Logger::default())
            .configure(router::config)
        // serve static files
        //.service(fs::Files::new("/", "./static/").index_file("index.html"))
    })
        .bind("127.0.0.1:8000")?
        .run()
        //.shutdown_timeout(2)
        //.start();
    //sys.run()
}

