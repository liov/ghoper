///两数之和
///

pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    for i in 0..nums.len(){
        for j in i+1..nums.len(){
            if nums[j] == target - nums[i]{
                return vec![i as i32,j as i32]
            }
        }
    }
    vec![0,0]
}
