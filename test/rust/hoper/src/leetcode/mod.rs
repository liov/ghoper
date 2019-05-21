///两数之和
use std::collections::HashMap;



//暴力法52ms，2MB
//一遍hash版本，0ms，2.7MB
pub fn two_sum1(nums: Vec<i32>, target: i32) -> Vec<i32> {
    for i in 0..nums.len(){
        for j in i+1..nums.len(){
            if nums[j] == target - nums[i]{
                return vec![i as i32,j as i32]
            }
        }
    }
    panic!("不存在")
}

pub fn two_sum2(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut map: HashMap<i32, usize> = HashMap::new();
    let mut index = 0;
    while index < nums.len() {
        if let Some(j) = map.get(&(target - nums[index])) {
            return vec![*j as i32, index as i32];
        }
        map.insert(nums[index], index);
        index = index + 1
    }
    panic!("不存在")
}


///两数相加
//执行用时 : 4 ms, 在Add Two Numbers的Rust提交中击败了100.00% 的用户

//内存消耗 : 2 MB, 在Add Two Numbers的Rust提交中击败了100.00% 的用户
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }

    pub fn push(&mut self, val: i32) {
        match self.next {
            Some(ref mut next) =>
                next.push(val),
            None =>self.next = Some(Box::new(ListNode::new(val)))
        }
    }
}

pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut next1 = l1;
    let mut next2 = l2;
    let mut result = Box::new(ListNode::new(0));
    let mut first =true;
    let mut carry= 0;
    let mut sum = 0;
    loop {
        let lx = next1.unwrap_or(Box::new(ListNode::new(0)));
        let ly = next2.unwrap_or(Box::new(ListNode::new(0)));

        sum = lx.val+ly.val;

        if first {
            result.val = sum%10;
            first = false;
            if sum >= 10 { carry =1}else { carry = 0 }
        }else {
            if sum%10+carry == 10{
                result.push(0);
                carry = 1;
            }else {
                result.push(sum%10+carry);
                if sum >= 10 { carry =1}else { carry = 0 }
            }
        }

        next1 = lx.next;
        next2 = ly.next;
        if next1==None&&next2==None&&carry==0{
            break;
        }
    }
    Some(result)
}

///字符串相乘

//执行用时 : 0 ms, 在Multiply Strings的Rust提交中击败了100.00% 的用户
//内存消耗 : 2 MB, 在Multiply Strings的Rust提交中击败了100.00% 的用户
/*m 位数和n位数相乘，结果位数为m+n-1或m + n，因此存进位数据数组大小申请为m+n位
对应位相乘后的结果与进位数据数组对应位置相加，十位数存入进位数组下一位，个位数留在该位
存进位数据数组在转为字符串返回的时候，要把前导零去掉*/
pub fn multiply(num1: String, num2: String) -> String {
    let num_vec1 = num1.as_bytes();
    let num_vec2 = num2.as_bytes();
    let cap =num_vec1.len()+num_vec2.len();
    let mut result =Vec::with_capacity(cap);
    unsafe  {result.set_len(cap);}
    let con = 48;
    let mut product = 0; //乘积
    let mut decade = 0; //十位

    if (num_vec1.len()==1&&num_vec1[0]==con)||(num_vec2.len()==1&&num_vec2[0]==con){
        return String::from("0")
    }

    for i in  0..num_vec1.len(){
        for j in 0..num_vec2.len(){
            if i==0 {result[cap-i-j-1] = 0} else if j == 0 { result[cap-i-num_vec2.len()-1] = 0 }
            product =(num_vec1[num_vec1.len()-i-1]-con)*(num_vec2[num_vec2.len()-j-1]-con)+decade;
            decade = product/10+(product%10 + result[cap-i-j-1])/10;
            result[cap-i-j-1] = (product%10 + result[cap-i-j-1])%10;

            if i==num_vec1.len()-1 {
                result[cap-i-j-1] = result[cap-i-j-1] +con
            } else if j == 0 {
                result[cap-i-j-1] = result[cap-i-j-1] +con
            }
        }
        result[cap-i-num_vec2.len()-1] = decade;
        decade = 0;

    }

    if result[0] != 0 { result[0] = result[0] +con} else {result.remove(0); }
    String::from_utf8(result).unwrap()
}

///电话号码的字母组合

//Y 组合子是 lambda 演算中的一个概念，是任意函数的不动点，在函数式编程中主要作用是 提供一种匿名函数的递归方式。
//"λ" 字符可以看作 function 声明，"."字符前为参数列表，"."字符后为函数体。
//不动点是函数的一个特征：对于函数 f(x)，如果有变量  a 使得  f(a)=a 成立，则称 a 是函数 f 上的一个不动点。

//注意结尾没有分号的那一行 x+1，与你见过的大部分代码行不同。表达式的结尾没有分号。如果在表达式的结尾加上分号，
// 它就变成了语句，而语句不会返回值。在接下来探索具有返回值的函数和表达式时要谨记这一点。

pub fn letter_combinations(digits: String) -> Vec<String> {
    let combination = vec![
        vec![b'a',b'b',b'c'],
        vec![b'd',b'e',b'f'],
        vec![b'g',b'h',b'i'],
        vec![b'j',b'k',b'l'],
        vec![b'm',b'n',b'o'],
        vec![b'p',b'q',b'r',b's'],
        vec![b't',b'u',b'v'],
        vec![b'w',b'x',b'y',b'z'],
    ];

    let mut result =Vec::new();

    let cap =digits.len();
    let mut middle:Vec<u8> =Vec::with_capacity(cap);
    unsafe  {middle.set_len(cap);}

    let digits_vec = digits.as_bytes();
    //闭包递归实在不会，改成了函数递归，消耗几何未知，有个clone()难受啊
    fn get(n:usize,  res:&mut Vec<String>,  mid: &mut Vec<u8>, com:&Vec<Vec<u8>>, dig:&[u8]){
        if n < mid.len() {
            for i in &(com[dig[n] as usize -50]){
                mid[n] = *i;
                get(n+1,res,mid,com,dig);
                if n == mid.len()-1 {
                    res.push(String::from_utf8(mid.clone()).unwrap());
                }
            }
        }
    }


    //let Y = |y|(|x|y(x(x)))(|x|y(|n|x(x)(n)));

    get(0,&mut result,&mut middle,&combination,&digits_vec);
    for i in &result{
        println!("{}",i)
    }
    vec![String::from("a")]
}

///两数相除

pub fn divide(dividend: i32, divisor: i32) -> i32 {
    if dividend ==-0x80000000 && divisor == -1{
        return  0x7fffffff
    }
    if dividend == 0 {
        return 0
    }
    if divisor == 1 {return dividend}
    if divisor == -1 {return -dividend}
    let mut result = 0;
    let mut dividend_copy = dividend;
    let mut divisor_copy = divisor;
    if dividend_copy > 0 {dividend_copy = -dividend_copy}
    if divisor_copy > 0 {divisor_copy = -divisor_copy}
    while dividend_copy <= divisor_copy{
        dividend_copy = dividend_copy - divisor_copy;
        result+=1;
    }
    if (dividend>0&&divisor<0)||(dividend<0&&divisor>0){
        return  -result
    }
    result
}
