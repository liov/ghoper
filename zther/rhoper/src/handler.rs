use actix_web::{App, Either, web, http, HttpRequest, HttpResponse, Result,
                middleware::{
                    Logger,
                },
                error::{self, Error}, Responder, dev::*};
use bytes::Bytes;
use chrono::Duration;
use std::cell::Cell;
use futures::stream::once;
use futures::future::{Future, result};


pub fn index(info: web::Path<(u32, String)>) -> impl Responder {
    format!("Hello {}! id:{}", info.1, info.0)
}

pub fn index2(_req: HttpRequest) -> &'static str {
    "Hello world!"
}

pub fn index3(_req: HttpRequest) -> impl Responder {
    Bytes::from_static("Hello world!".as_ref())
}
// 有状态
pub struct AppState {
    pub counter: Cell<usize>,
}

pub fn index5(state: web::Data<AppState>) -> String {
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

/*//自定义响应类型
#[derive(Serialize)]
struct MyObj {
    name: &'static str,
}

/// Responder
impl Responder for MyObj {
    type Error = Error;
    type Future = Box<dyn Future<Item=HttpResponse, Error=Error>>;

    fn respond_to(self, _: &HttpRequest) ->  Box<Future<Item=HttpResponse, Error=Error>> {
        let body = serde_json::to_string(&self).unwrap();

        // Create response and set content type
        Box( Ok(HttpResponse::Ok()
            .content_type("application/json")
            .body(body)))
    }
}

pub fn index7(req: &HttpRequest) -> impl Responder {
    MyObj { name: "user" }
}*/

/*// 异步内容
pub fn index12(req: &HttpRequest) -> Box<Future<Item=HttpResponse, Error=Error>> {
    result(Ok(HttpResponse::Ok()
        .content_type("text/html")
        .body(format!("Hello!"))))
}*/

/*
fn index13(req: &HttpRequest) -> Box<Future<Item=&'static str, Error=Error>> {
    result(Ok("Welcome!"))

}
*/
/*
fn index16(req: &HttpRequest) -> HttpResponse {
    let body = once(Ok(Bytes::from_static(b"test")));

    HttpResponse::Ok()
        .content_type("application/json")
        .streaming(Box::new(body))
}*/

fn is_error() -> bool {
    false
}

fn index15(_req: &HttpRequest) -> Result<Box<Future<Item=HttpResponse, Error=Error>>, Error> {
    if is_error() {
        Err(error::ErrorBadRequest("bad request"))
    } else {
        Ok(Box::new(
            result(Ok(HttpResponse::Ok()
                .content_type("text/html")
                .body(format!("Hello!"))))))
    }
}



/*
type RegisterResult = Either<HttpResponse, Box<Future<Item=HttpResponse, Error=Error>>>;

fn is_a_variant() -> bool {
    true
}

fn index17(req: &HttpRequest) -> RegisterResult {
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
*/

fn index18() -> Result<HttpResponse> {
    Ok(HttpResponse::Ok()
        .content_type("text/html; charset=utf-8")
        .body(include_str!("../../../static/template/rust/index.html")))
}
