const getCookie = function(cname, req) {
  const name = cname + '='
  let decodedCookie
  if (typeof window === 'undefined')
    decodedCookie = decodeURIComponent(req.headers.cookie)
  else decodedCookie = decodeURIComponent(document.cookie)
  const ca = decodedCookie.split(';')
  for (let i = 0; i < ca.length; i++) {
    let c = ca[i]
    while (c.charAt(0) === '') {
      c = c.substring(1)
    }
    if (c.indexOf(name) === 0) {
      return c.substring(name.length, c.length)
    }
  }
  return ''
}

// 设置cookie
function setCookie(c_name, value, expiredays) {
  const exdate = new Date()
  exdate.setDate(exdate.getDate() + expiredays)
  document.cookie =
    c_name +
    '=' +
    escape(value) +
    (expiredays == null ? '' : ';expires=' + exdate.toUTCString())
}

// 获取cookie
function get_Cookie(name) {
  let arr
  const reg = new RegExp('(^| )' + name + '=([^;]*)(;|$)')
  if ((arr = document.cookie.match(reg))) {
    return decodeURIComponent(arr[2])
  } else {
    return null
  }
}

// 删除cookie
function delCookie(name) {
  const exp = new Date()
  exp.setTime(exp.getTime() - 1)
  const cval = getCookie(name)
  if (cval != null) {
    document.cookie = name + '=' + cval + ';expires=' + exp.toUTCString()
  }
}

export default {
  getCookie,
  delCookie
}
