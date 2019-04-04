const jsonp = function(url) {
  // 创建script标签，设置其属性
  const script = document.createElement('script')
  script.setAttribute('src', url)
  // 把script标签加入head，此时调用开始
  document.getElementsByTagName('head')[0].appendChild(script)
}
export default { jsonp }
