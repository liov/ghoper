use actix_web::{App, web::{self,ServiceConfig}, http::Cookie, dev::{WebService, Body}};
use super::handler;

use actix_session::CookieSession;
use chrono::Duration;
use std::cell::Cell;
use core::borrow::Borrow;

//这是一件不可思议的事情
//事实是在此之前我已经做过测试
//一个泛型函数，那么这个泛型是需要在调用前确定而不是调用中确定，且需要在参数中确定
//因此一个包含泛型的结构体调用一个泛型函数是可行的
//然而我们并不能够用泛型定义函数，然后在函数内部明确这个泛型，也就是说函数标注的泛型是动态的
//这在理想中应该是可行的，因为函数内部的返回值其实是可推断的，当然我不否认存在不可推断的情况
//然而我认为这仍然是可实现的，动态泛型，泛型本来就代表了动态，让其动态彻底不过是增加runtime
//我认为在降低编程复杂性和性能之间，我们可以多出一个选项来
//不过目前并不能这样做
//以至于我写出了这样的代码
//new(APP.new()),我想我应该将这个new写进方法链式调用
//仍然不行
//no method named `data` found for type `&mut actix_web::app::App<T, B>` in the current scope
//放弃
/*pub(crate) fn new<T,B>(app:&mut App<T,B>) -> &App<T,B> {
    let secret: String =
        std::env::var("SECRET_KEY").unwrap_or_else(|_| "0123".repeat(8));
    let domain: String =
        std::env::var("DOMAIN").unwrap_or_else(|_| "localhost".to_string());

    let state = handler::AppState { counter: Cell::new(0usize) };
    app.data(state)
        .wrap(Logger::default())
        .wrap(CookieSession::signed(&[0; 32])
            .path("/")
            .domain(domain.as_str())
            .max_age_time(Duration::days(1)).
            secure(false))
    .service(
        web::scope("/test")
            .service(web::resource("/index1").route(web::get().to(handler::index1)))
            .service(web::resource("/index2").route(web::get().to(handler::index2)))
            .service(web::resource("/index3").route(web::get().to(handler::index3)))
            .service(web::resource("/index5").route(web::get().to(handler::index5)))
    )
    // serve static files
    //.service(fs::Files::new("/", "./static/").index_file("index.html"))
}*/

pub fn config(cfg: &mut ServiceConfig){
    cfg.service(
        web::scope("/test")
            .service(web::resource("/{id}/{name}").to(handler::index))
            .service(web::resource("/index2").route(web::get().to(handler::index2)))
            .service(web::resource("/index3").route(web::get().to(handler::index3)))
            .service(web::resource("/index5").route(web::get().to(handler::index5)))
    );
}
