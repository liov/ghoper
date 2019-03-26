#![feature(async_await,await_macro, futures_api)]

use hoper::utils::tree::MyTree;
use std::ops::Deref;


fn main() {
/*   println!("猜？？!");

   println!("Please input your guess.");

   let secret_number = rand::thread_rng().gen_range(1, 101);
   loop {
       let mut guess = String::new();

       io::stdin().read_line(&mut guess).expect("Failed to read line");

       let guess: u32 = match guess.trim().parse() {
           Ok(num) => num,
           Err(_) => continue,
       };

       match guess.cmp(&secret_number) {
           Ordering::Less => println!("Too small!"),
           Ordering::Greater => {println!("Too big!");break;},
           Ordering::Equal => {
               println!("You guessed: {}", guess);
               println!("You win!");
               break;
           },
       }
   }

    let heart_eyecat = '😻';

   println!("{}", heart_eyed_cat);

   let mut user: User = User{
       id: 0,
       sex: Sex::Boy,
       name: String::new(),
       password: String::new(),
       email: String::new(),
       phone: String::new(),
   };

   println!("{:?}",user);

   user.set_name(String::from("哈哈s"));

   println!("{:?}",user);

   println!("{}",user.name.len());

   let mut message:Message=Message::Write(String::from("sss"));
   println!("{:?}",message);


   let a: [i32; 3];


   let args: Vec<String> = env::args().collect();*/

/*    ::std::env::set_var("RUST_LOG", "hoper=info");
    ::std::env::set_var("RUST_BACKTRACE", "1");

    let sys = System::new("hoper");

    server::new( move || router::app().unwrap())
        .bind("127.0.0.1:8000").unwrap()
        .shutdown_timeout(2)
        .start();

    sys.run();*/
    let x = std :: f64 :: consts :: PI;
    let mut t = MyTree::new();
    t.insert(x);
    println!("{:?}",t);
    t.insert(3f64);
    println!("{:?}",t);
    t.insert(5f64);
    println!("{:?}",t);
    t.insert(9f64);
    println!("{:?}",t);
    t.insert(66f64);
    println!("{:?}",t);
    t.insert(18f64);
    println!("{:?}",t);
    t.insert(12f64);
    println!("{:?}",t);
    t.insert(111f64);
    println!("{:?}",t);
    t.insert(12f64);
    println!("{:?}",t);
    t.insert(2f64);
    println!("{:?}",t);
    t.peek();
    let m = MyBox::new("Rust");
    hello(&m);
    env_logger::init();

}



struct MyBox<T>(T);
impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}
impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &T {
        &self.0
    }
}
impl<T> Drop for MyBox<T> {
    fn drop(&mut self) {
        println!("清理Mybox")
    }
}
fn hello(name: &str) {
    println!("Hello, {}!", name);
}

#[derive(Debug)]
enum Message {
   Quit,
   Move { x: i32, y: i32 },
   Write(String),
   ChangeColor(i32, i32, i32),
}

impl Message {
    async  fn call(&self) {
       // 在这里定义方法体
       await!(async {
            // 省略业务代码
            "set state".to_owned()
        });
   }
}

trait Conn {
   fn connect(&self) ->i32;
}

struct Cacher<T>
   where T: Fn(u32) -> u32
{
   calculation: T,
   value: Option<u32>,
}

impl<T> Cacher<T> where T: Fn(u32) -> u32 {
   fn new(calculation: T) -> Cacher<T> {
       Cacher {
           calculation,
           value: None,
       }
   }

   fn value(&mut self, arg: u32) -> u32 {
       match self.value {
           Some(v) => v,
           None => {
               let v = (self.calculation)(arg);
               self.value = Some(v);
               v
           },
       }
   }
}
