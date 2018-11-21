var __extends = (this && this.__extends) || (function () {
    var extendStatics = Object.setPrototypeOf ||
        ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
        function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var Animal = /** @class */ (function () {
    function Animal() {
    }
    return Animal;
}());
var Dog = /** @class */ (function (_super) {
    __extends(Dog, _super);
    function Dog() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    return Dog;
}(Animal));
// 错误：使用数值型的字符串索引，有时会得到完全不同的Animal!
/*
interface NotOkay {
    [x: number]: Animal;
    [x: string]: Dog;
}

let animal1:Animal = {name:"dog1"};
let dog1:Dog ={breed:"labuladuo1"};
let animal2:Animal = {name:"dog2"};
let dog2:Dog ={breed:"labuladuo2"};
let a:NotOkay={};
a[1]=animal1;
a[2]=animal2;
a["1"]=dog1;
a["2"]=dog2
console.log(a[1])
懂了！
*/
