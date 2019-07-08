use std::sync::{mpsc, Arc};
use std::{thread, fmt};
use std::rc::Rc;
use std::fmt::Display;
use std::time::Duration;

pub struct Foo<T>{
    id:Rc<T>
}

impl<T:Display> Display for Foo<T> {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "foo {}", self.id)
    }
}

unsafe impl<T> Send for Foo<T> {}

// 线程数量
const THREAD_COUNT :i32 = 10;

fn main(){
    let var : Arc<i32> = Arc::new(5);
    let share_var = var.clone();

    // 创建一个通道
    let (tx, rx)= mpsc::sync_channel(0);
    let new_thread = thread::spawn(move || {
        // 创建线程用于发送消息
        for id in 0..THREAD_COUNT {
            // 注意Sender是可以clone的，这样就可以支持多个发送者
            let thread_tx = tx.clone();
            thread::spawn(move || {
                // 发送一个消息，此处是数字id
                thread_tx.send(id + 1).unwrap();
                println!("send {}", id + 1);
            });
            //顺序打印
            thread::sleep(Duration::from_millis(100));
        }
    });

    // 在主线程中接收子线程发送的消息并输出
    for _ in 0..THREAD_COUNT {
        println!("receive {}",rx.recv().unwrap());
    }

    new_thread.join().unwrap();

    let new_thread = thread::spawn(move || {

    });
}
