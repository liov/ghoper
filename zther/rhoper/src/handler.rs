use actix_web::{App, Either, web, http, HttpRequest, HttpResponse, Result,
                middleware::{
                    identity::{CookieIdentityPolicy, IdentityService},
                    Logger,
                },
                error::{self, Error}, Responder, dev::*};
use bytes::Bytes;
use chrono::Duration;



pub fn index1(req: HttpRequest) -> String {
    "Hello world!".to_owned()
}

pub fn index2(_req: HttpRequest) -> &'static str {
    "Hello world!"
}

pub fn index3(_req: HttpRequest) -> impl Responder {
    Bytes::from_static("Hello world!".as_ref())
}
