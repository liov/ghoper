//异步函数加了await相当于是同步的，但是如果其他函数掉这个异步函数没有await是异步执行的
const fs = require('fs');

const writeFile = function (fileName,data) {
    return new Promise(function (resolve, reject) {
      fs.writeFile(fileName,data, function(error) {
        if (error) return reject(error);
        resolve('写入成功');
      });
    });
  };

async function foo(){
   let r = await writeFile('../foo.txt','foo')
   console.log(r)
}

async function writeFileAsync(){
    await foo()
    console.log('正在异步写入')
}

writeFileAsync()

const readFile = function (fileName) {
    return new Promise(function (resolve, reject) {
      fs.readFile(fileName, function(error, data) {
        if (error) return reject(error);
        resolve(data);
      });
    });
  };

async function bar(){
    let f = await readFile('../foo.txt')
    console.log(f.toString())
}

function writeFileSync(){
    bar()
    console.log('正在同步写入')
}

writeFileSync()

function fib(n){
  return n<2?1:fib(n-2)+fib(n-1)
}

async function fibAsync(n){
  return new Promise(function (resolve, reject){
    resolve(fib(n)); 
  })
}

async function singleThread(n){
  console.log('计算斐波那契')
  let data = await fibAsync(n)
  console.log(data)
}
//这个函数证明了什么，如我所想，node是单线程的，所以在涉及非io操作时，异步函数是在一个线程里执行的
//本应在计算的同时就输出计算中，但是是计算完成后才输出的
function singleThreadTest(){
  singleThread(45)
  console.log('计算中')
}

singleThreadTest()