fn main(){
    let b = Box::new(5);
    let addr =&**&b as *const i32;
    println!("{:p}:{:p}", *&b,addr);
    test1(b);
    println!("{:p}:{}",addr, unsafe { *addr })
}

fn test1(a:Box<i32>){
    let addr = a.as_ref();
    println!("{:p}:{}", addr,*addr);
}

fn test2(a:i32){
    println!("{:p}",&a);
}
