import axios from 'axios'
const upload = async function(classify, $file) {
  // 第一步.将图片上传到服务器.
  const formdata = new FormData()
  formdata.append('file', $file)
  const res = await axios({
    url: '/api/upload/' + classify,
    method: 'post',
    data: formdata,
    headers: { 'Content-Type': 'multipart/form-data' }
  })
  return res.data.data.url
}

const getBase64 = function(img, callback) {
  const reader = new FileReader()
  reader.addEventListener('load', () => callback(reader.result))
  reader.readAsDataURL(img)
}

export { upload, getBase64 }
