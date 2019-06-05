var fs=require("fs");
var path = require("path");
//递归创建目录 同步方法
function mkdirsSync(dir) {
  if (fs.existsSync(dir)) {
    return true;
  } else {
    if (mkdirsSync(path.dirname(dir))) {
      fs.mkdirSync(dir);
      return true;
    }
  }
}
function copy(src, dir,fileName) {
  const dst = dir+"/"+fileName;
  if (fs.existsSync(dst)) return
  mkdirsSync(dir)

  fs.open(dst, 'w+',function(err, fd) {

    if (err) {
      return console.error(err);
    }
    console.log("文件打开成功！");
  })
  fs.writeFileSync(dst, fs.readFileSync(src));
}

exports.init =function() {
  copy("./node_modules/tinymce/skins/ui/oxide/skin.min.css","../../static/css/tinymce/skins/ui/oxide","skin.min.css");
  copy("./node_modules/tinymce/skins/ui/oxide/content.min.css","../../static/css/tinymce/skins/ui/oxide","content.min.css");

}


