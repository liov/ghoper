#![feature(str_as_mut_ptr)]

static mut HELLO_WORLD: &str = "JING_TAI_BIAN_LIANG__Hello, world!__JING_TAI_BIAN_LIANG";
static mut UBA: u8 = 125u8;

fn main() {
    unsafe {
        /*  静态值，放在代码区，不可修改
                let x = &12;
                let x_addr= x as *const i32 as *mut i32;
                *x_addr = 13;
                println!("x的值：{}",x);*/
        let a1 = "hello word";
        let a2 = "hello word";
        println!("相同字符串字面值是否一个地址：{}", a1.as_ptr() == a2.as_ptr());


        let ptr = HELLO_WORLD.as_ptr() as *const i32;
        println!("HELLO_WORLD的地址：{:p}", ptr);
        HELLO_WORLD = "World, hello!";
        /*let mut_ptr = HELLO_WORLD.as_ptr() as *mut i32;
        *mut_ptr = 1196312906;*/
        println!("HELLO_WORLD的地址：{:p},值：{},首字母：{:?}", HELLO_WORLD.as_ptr(), HELLO_WORLD, *ptr);


        //let mut a = "a var of a";
        //let u8a = a.as_ptr() as *mut u8;
        //栈内str，可修改
        let a_bytes = [b'a', b' ', b'v', b'a', b'r', b' ', b'o', b'f', b' ', b'a'];

        let a = std::str::from_utf8_unchecked(&a_bytes);


        let addr_str_a = a as *const str as *mut u8;
        let addr_a = &a as *const &str;
        let ua = b"c var of c";

        let u = 72u8;
        let u_a = &u as *const u8 as *mut u8;


        println!("u8a:{:?}", addr_str_a);
        //*u8a = [b'c',b' ',b'v',b'a',b'r',b' ',b'o',b'f',b' ',b'c'];
        //*u8a=b'H';
        let uba_ptr = &UBA as *const u8 as*mut u8;
        *uba_ptr += 1;
        *u_a += 1;
        println!("u_a:{:?}", *u_a);
        //*u8a_asu8+=1;
        addr_str_a.copy_from(ua.as_ptr(), 10);
        println!("UBA:{:?},UBA的地址：{:p},a的值：{}", UBA, &UBA as *const u8, a);

        println!("a的地址：{:p},a的底层地址：{:p},a的值：{}", addr_a, addr_str_a, a);


        let mut b = a;
        let addr_b = &b as *const &str;
        println!("b的地址：{:p}", addr_b);
        b.replace("a", "b");
        let addr_b = &b as *const &str;

        println!("b的地址：{:p},b的值：{}", addr_b, *addr_b);


        let c = String::from("c var of c");
        let addr_c = &c as *const String;
        println!("c的地址：{:p}", addr_c);
        let d = c;
        let addr_d = &d as *const String;

        println!("d的地址：{:p},c的值：{}", addr_d, *addr_c);


        let mut e = [1, 1, 1, 1, 1, 1, 1, 1, 1];
        let f = &mut e[3..];
        f[0] = 0;
        println!("e的值：{:?}", e);


        let user = User { name: String::from("jyb"), phone: 196206 };
        let addr_user = &user as *const User;
        println!("user的地址：{:p}", addr_user);
        pass_by_value_ref_user(&user);
        pass_by_value_user(user);
        //奇怪，不同终端表现不一致
        println!("原来的user：{:?}", *addr_user);
    }
}


fn pass_by_value_ref_user(user: &User) {
    let addr_user = &user as *const &User;
    println!("借用函数参数的的地址：{:p}", addr_user);
}

fn pass_by_value_user(mut user: User) {
    let addr_user = &user as *const User;
    user.phone = 2;
    println!("所有权函数参数的的地址：{:p},user的值：{:?}", addr_user, user);
    //std::mem::forget(user);
}

#[derive(Debug)]
struct User {
    name: String,
    phone: i32,
}
