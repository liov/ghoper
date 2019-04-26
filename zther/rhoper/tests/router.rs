use actix_web::{App, Either, web, http, HttpRequest, HttpResponse, Result,
                middleware::{
                    identity::{CookieIdentityPolicy, IdentityService},
                    Logger,
                },
                error::{self, Error}, Responder, dev::*};
use actix::dev::Handler;
use bytes::Bytes;
use futures::stream::once;
use futures::future::{Future, result};
/*use std::sync::Arc;
use std::sync::atomic::{AtomicPtr, Ordering};*/
use std::cell::Cell;
use chrono::Duration;
use crate::{auth_routes, invitation_routes, register_routes};
use actix_files as fs;
use actix::prelude::Addr;
use crate::models::DbExecutor;
use actix::prelude::*;

pub fn app(address: Addr<DbExecutor>) -> std::result::Result<App<(), ()>, ()> {
    /*    let inc = Arc::new(AtomicPtr::new(0));
        let cloned = inc.clone();*/
    let secret: String =
        std::env::var("SECRET_KEY").unwrap_or_else(|_| "0123".repeat(8));
    let domain: String =
        std::env::var("DOMAIN").unwrap_or_else(|_| "localhost".to_string());
    let app = App::new()
        .data(address)
        .data(AppState { counter: Cell::new(0) })
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
            .service(web::resource("/index2").to(index2))
            .service(web::resource("/index3").to(index3))
            .service(web::resource("/index5").to(index5))
            //.resource("/index6", move |r| r.h(MyHandler(cloned)))
            .resource("/index7", |r| r.method(http::Method::GET).f(index7))
            .resource("/index12", |r| r.route().a(index12))
            .resource("/index13", |r| r.route().a(index13))
            .resource("/index14", |r| r.f(index14))
            .resource("/index15", |r| r.f(index15))
            .resource("/index16", |r| r.f(index16))
    ).service(
        web::scope("/api")
            .service(
                web::resource("/auth")
                    .route(web::post().to_async(auth_routes::login))
                    .route(web::delete().to(auth_routes::logout))
                    .route(web::get().to_async(auth_routes::get_me)),
            )
            // routes to invitation
            .service(
                web::resource("/invitation").route(
                    web::post().to_async(invitation_routes::register_email),
                ),
            )
            // routes to register as a user after the
            .service(
                web::resource("/register/{invitation_id}")
                    .route(web::post().to_async(register_routes::register_user)),
            ),
    )
        // serve static files
        .service(fs::Files::new("/", "./static/").index_file("index.html"))

        .finish();
    Ok(app)
}


// 静态文本
fn index1(req: &HttpRequest) -> String {
    "Hello world!".to_owned()
}

fn index2(_req: &HttpRequest) -> &'static str {
    "Hello world!"
}

fn index3(_req: &HttpRequest) -> impl Responder {
    Bytes::from_static("Hello world!".as_ref())
}


// 有状态
struct AppState {
    counter: Cell<usize>,
}

fn index5(state: web::Data<AppState>) -> String {
    let count = state.counter.get() + 1; // <- get count
    state.counter.set(count); // <- store new count in state
    format!("Request number: {}", count) // <- response with count
}

/*
struct MyHandler(Arc<AtomicPtr<usize>>);

impl<S> Handler<S> for MyHandler {
    type Result = HttpResponse;

    /// Handle request
    fn handle(&mut self, msg:S, _ :&mut  &mut Context<Self>) -> Self::Result {
        self.0.fetch_add(1, Ordering::Relaxed);
        HttpResponse::Ok().into()
    }
}
*/

//自定义响应类型
#[derive(Serialize)]
struct MyObj {
    name: &'static str,
}

/// Responder
impl Responder for MyObj {
    type Error = Error;
    type Future = Box<dyn Future<Item=HttpResponse, Error=Error>>;

    fn respond_to(self, _: &HttpRequest) -> Result<HttpResponse, Error> {
        let body = serde_json::to_string(&self)?;

        // Create response and set content type
        Ok(HttpResponse::Ok()
            .content_type("application/json")
            .body(body))
    }
}

fn index7(req: &HttpRequest) -> impl Responder {
    MyObj { name: "user" }
}

// 异步内容
fn index12(req: &HttpRequest) -> Box<Future<Item=HttpResponse, Error=Error>> {
    result(Ok(HttpResponse::Ok()
        .content_type("text/html")
        .body(format!("Hello!"))))
        .responder()
}

fn index13(req: &HttpRequest) -> Box<Future<Item=&'static str, Error=Error>> {
    result(Ok("Welcome!"))
        .responder()
}

fn index14(req: &HttpRequest) -> HttpResponse {
    let body = once(Ok(Bytes::from_static(b"test")));

    HttpResponse::Ok()
        .content_type("application/json")
        .body(Body::Streaming(Box::new(body)))
}

fn is_error() -> bool {
    false
}

fn index15(req: &HttpRequest) -> Result<Box<Future<Item=HttpResponse, Error=Error>>, Error> {
    if is_error() {
        Err(error::ErrorBadRequest("bad request"))
    } else {
        Ok(Box::new(
            result(Ok(HttpResponse::Ok()
                .content_type("text/html")
                .body(format!("Hello!"))))))
    }
}


fn index16(req: &HttpRequest) -> impl Responder {
    let body = once(Bytes::from_static(b"Hello Word!"));

    HttpResponse::Ok()
        .content_type("application/json")
        .body(BodyStream::Streaming(Box::new(body)))
}


type RegisterResult = Either<HttpResponse, Box<Future<Item=HttpResponse, Error=Error>>>;

fn is_a_variant() -> bool {
    true
}

fn index17(req: &HttpRequest) -> impl Responder {
    if is_a_variant() { // <- choose variant A
        Either::A(
            HttpResponse::BadRequest().body("Bad data"))
    } else {
        Either::B(
            // <- variant B
            result(Ok(HttpResponse::Ok()
                .content_type("text/html")
                .body(format!("Hello!")))))
    }
}

fn index18() -> Result<HttpResponse> {
    Ok(HttpResponse::Ok()
        .content_type("text/html; charset=utf-8")
        .body(include_str!("../static/index.html")))
}