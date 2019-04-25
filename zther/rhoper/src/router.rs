use actix_web::{App,Either,Error,Responder,HttpRequest,HttpResponse,Body,fs,http::{Method,header},middleware::{self,cors::Cors},dev::Handler};
use bytes::Bytes;
use futures::stream::once;
use futures::future::{Future, result};
use std::sync::Arc;
use std::sync::atomic::{AtomicUsize, Ordering};

pub fn app() -> Result<App<()>,()> {

    let app = App::new()
        .resource("/", |r| r.f(index1));
    Ok(app)
}




// 静态文本
fn index2(_req: HttpRequest) -> &'static str {
    "Hello world!"
}

fn index3(_req: &HttpRequest) -> impl Responder {
    Bytes::from_static("Hello world!")
}

fn index(req: HttpRequest) -> String {
    "Hello world!".to_owned()
}

// 有状态
fn index5(req: &HttpRequest<AppState>) -> String {
    let count = req.state().counter.get() + 1; // <- get count
    req.state().counter.set(count); // <- store new count in state
    format!("Request number: {}", count) // <- response with count
}
struct MyHandler(Arc<AtomicUsize>);

impl<S> Handler<S> for MyHandler {
    type Result = HttpResponse;

    /// Handle request
    fn handle(&mut self, req: HttpRequest<S>) -> Self::Result {
        self.0.fetch_add(1, Ordering::Relaxed);
        HttpResponse::Ok().into()
    }
}

#[derive(Serialize)]
struct MyObj {
    name: &'static str,
}

/// Responder
impl Responder for MyObj {
    type Item = HttpResponse;
    type Error = Error;

    fn respond_to<S>(self, req: &HttpRequest<S>) -> Result<HttpResponse, Error> {
        let body = serde_json::to_string(&self)?;

        // Create response and set content type
        Ok(HttpResponse::Ok()
            .content_type("application/json")
            .body(body))
    }
}

fn index6(req: HttpRequest) -> impl Responder {
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

type RegisterResult = Either<HttpResponse, Box<Future<Item=HttpResponse, Error=Error>>>;

fn index16(req: &HttpRequest) -> impl Responder {
    let body = once(Ok(Bytes::from_static(b"Hello Word!")));

    HttpResponse::Ok()
        .content_type("application/json")
        .body(Body::Streaming(Box::new(body)))
}
