#![allow(unused_imports)]

mod handler;
mod auth_handler;
mod auth_routes;
mod email_service;
mod errors;
mod invitation_handler;
mod invitation_routes;
mod models;
mod register_handler;
mod register_routes;
mod schema;
mod utils;
mod router;

use actix_web::{App, Either, web, http, HttpRequest, HttpResponse,
                HttpServer,
                middleware::{
                    identity::{CookieIdentityPolicy, IdentityService},
                    Logger,
                },
                error::{self, Error}, Responder, dev::*};
use actix_session::{CookieSession, Session};
use actix::prelude::*;
use chrono::Duration;
use actix_files as fs;
use dotenv::dotenv;
use std::cell::Cell;
#[macro_use]
extern crate diesel;
#[macro_use]
extern crate serde_derive;

fn main() -> std::io::Result<()> {
    dotenv().ok();
    std::env::set_var("RUST_LOG", "rhoper=info,actix_web=info,actix_server=info");
    std::env::set_var("RUST_BACKTRACE", "1");
    env_logger::init();


    let sys = actix_rt::System::new("hoper");
    HttpServer::new( || {
        let secret: String =
            std::env::var("SECRET_KEY").unwrap_or_else(|_| "0123".repeat(8));
        let domain: String =
            std::env::var("DOMAIN").unwrap_or_else(|_| "localhost".to_string());
        let state = handler::AppState{ counter: Cell::new(0usize) };


    })
        .bind("127.0.0.1:8000")?
        .shutdown_timeout(2)
        .start();
        sys.run()
}
