fn main(){
    let a = 5;
    let add_a =&a as *const i32;
    println!("{}",unsafe{*add_a});
    let b = Box::new(5);
    let addr_b = &*b as *const i32;
    println!("{:p}:{:p}", *&b,addr_b);
    test1(b); //注释此行，获取正确的值，不注释，错误的值
    println!("{:p}:地址addr_b的值：{}",addr_b, unsafe{*addr_b} )
}

fn test1(a:Box<i32>){
    let addr = a.as_ref();
    println!("{:p}:{}", addr,*addr);
    //此处获取所有权a，并执行完，释放了内存，因此最后获取到错误的值
}

fn test2(a:i32){
    println!("{:p}",&a);
}
