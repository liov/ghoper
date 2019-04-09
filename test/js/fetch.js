let token = "Bearer.eyJhbGciOiJIUzI1NiJ9.eyJ1aWQiOiIxIiwibG9naW5JcCI6IjEwLjQyLjEuMCIsImxvZ2luTmFtZSI6ImFkbWluIiwibG9naW5UaW1lTWlsbGlzIjoxNTU0NzkyMTAyNDEwLCJleHAiOjE1NTQ4Nzg1MDJ9.-2-q2NwRryz4Y9EcMwNJ1QUINwJPUHv2wh76Taei4Nc"

fetch('http://192.168.31.148:30002/api/dubbo/resource/export/res_export_trademark_info', {
    method: 'post', body: JSON.stringify({}), responseType: 'blob', headers: {
        "Content-Type": "application/json",
        "auth-token": token
    }
})
    .then(res => {
        return res.blob();
    }).then(blob => {


    let a = document.createElement('a');
    a.download = "test.xlsx";
    a.style.display = 'none';
    blob.type = "application/excel";
    let url = window.URL.createObjectURL(blob);
    a.href = url;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
})
