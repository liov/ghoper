///两数之和
use std::collections::HashMap;
use core::fmt::Debug;


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

    for i in  0..num_vec1.len(){
        for j in 0..num_vec2.len(){
            if i==0 {result[cap-i-j-1] = 0} else if j == 0 { result[cap-i-num_vec2.len()] = 0 }
            product =(num_vec1[num_vec1.len()-i-1]-con)*(num_vec2[num_vec2.len()-j-1]-con)+decade;
            result[cap-i-j-1] = product%10 + result[cap-i-j-1];
            if result[cap-i-j-1] >=10 {result[cap-i-j-1]-10;}
            decade = product/10+result[cap-i-j-1]/10;
            println!("{}：{}",cap-i-j-1,product%10);
        }
        result[cap-i-num_vec2.len()-1] = decade;
        decade = 0;
        println!("{}：{}",cap-i-num_vec2.len()-1,result[cap-i-num_vec2.len()-1]);
        println!();
    }

    if decade > 0 { result[0] = decade +con} else {result.remove(0); }
    for i in &result{
        *i+*i+con;
    }
    String::from_utf8(result).unwrap()
}
