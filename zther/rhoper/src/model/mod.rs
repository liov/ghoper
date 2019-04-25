/*模块系统变化，新版本引入了一些路径工作方式的变化，简化了模块系统，使其更加清晰：

大部分情况下不再需要 extern crate 。
可以直接使用 use 引入宏，而不再需要使用 #[macro_use] 属性。
绝对路径以 crate 名开头，关键字 crate 指代当前 crate。
foo.rs 和 foo/ 子目录共存，将子模块放在子目录中时不再需要 mod.rs*/
pub mod user;
pub mod tag;
