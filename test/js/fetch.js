fetch('http://localhost:8030/api/dubbo/resource/export/res_export_trademark_info', {
    method: 'post', body: JSON.stringify({}), responseType: 'blob', headers: {
        "Content-Type": "application/json",
        "auth-token": "Bearer.eyJhbGciOiJIUzI1NiJ9.eyJ1aWQiOiIxIiwibG9naW5JcCI6IjEwLjQyLjEuMCIsImxvZ2luTmFtZSI6ImFkbWluIiwibG9naW5UaW1lTWlsbGlzIjoxNTUzMjE4MjQ4NzA0LCJleHAiOjE1NTMzMDQ2NDh9.y86KHYvMOUV9-dm8xVWTqqj86o5PAr0J4IdyTRYtKbw"
    }
})
    .then(res => {
        return res.blob();
    }).then(blob => {


    var a = document.createElement('a');
    a.download = "test.xlsx";
    a.style.display = 'none';
    blob.type = "application/excel";
    var url = window.URL.createObjectURL(blob);
    a.href = url;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
})
