const copy = function (obj) {
    let newobj = {};
    for (let attr in obj) {
        newobj[attr] = obj[attr];
    }
    return newobj;
};

function writeObj(obj){
    var description = "";
    for(var i in obj){
        var property=obj[i];
        description+=i+" = "+property+"\n";
    }
    console.log("obj:"+description);
}
export {copy,writeObj}
