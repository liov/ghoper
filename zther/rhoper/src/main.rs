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

use actix_web::{App, Either, web, http, HttpRequest, HttpResponse,
                HttpServer,
                middleware::{
                    identity::{CookieIdentityPolicy, IdentityService},
                    Logger,
                },
                error::{self, Error}, Responder, dev::*};
use actix::prelude::*;
use chrono::Duration;
use actix_files as fs;
use dotenv::dotenv;
#[macro_use]
extern crate diesel;
#[macro_use]
extern crate serde_derive;

fn main() -> std::io::Result<()> {
    dotenv().ok();
    std::env::set_var("RUST_LOG", "rhoper=info,actix_web=info,actix_server=info");
    std::env::set_var("RUST_BACKTRACE", "1");
    env_logger::init();


    let sys = actix::System::new("hoper");
    HttpServer::new( || {
        let secret: String =
            std::env::var("SECRET_KEY").unwrap_or_else(|_| "0123".repeat(8));
        let domain: String =
            std::env::var("DOMAIN").unwrap_or_else(|_| "localhost".to_string());

        App::new()
            .wrap(Logger::default())
            .wrap(IdentityService::new(
                CookieIdentityPolicy::new(secret.as_bytes())
                    .name("auth")
                    .path("/")
                    .domain(domain.as_str())
                    .max_age_time(Duration::days(1))
                    .secure(false), // this can only be true if you have https
            )).service(
            web::scope("/test")
                .service(web::resource("/index1").route(web::get().to(handler::index1)))
                .service(web::resource("/index2").route(web::get().to(handler::index2)))
                .service(web::resource("/index3").route(web::get().to(handler::index3)))
        )
            // serve static files
            //.service(fs::Files::new("/", "./static/").index_file("index.html"))

    })
        .bind("127.0.0.1:8000")?
        .shutdown_timeout(2)
        .start();
        sys.run()
}