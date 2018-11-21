

function add(a,b,c,d,e,f){

    var params = {
        'tag_id': a,
        'title': b,
        'desc': c,
        'content': d,
        'created_by': e,
        'state': f,

    };

    return params
}


// We need this to build our post string
var querystring = require('querystring');
var http = require('http');


function PostCode() {
    // Build the post string from an object
    var post_data = querystring.stringify({
        'tag_id': 1,
        'title': 'test1',
        'desc': 'test-desc',
        'content': 'test-content',
        'created_by': 'test-created',
        'state': 1

    });

    // An object of options to indicate where to post to
    var post_options = {
        host: '127.0.0.1',
        port: '8000',
        path: '/api/article',
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Content-Lenght': post_data.length
        }
    };

    // Set up the request
    var post_req = http.request(post_options, function(res) {
        res.setEncoding('utf8');
        res.on('data', function (chunk) {
            console.log('Response: ' + chunk);
        });
    });

    // post the data
    post_req.write(post_data);
    post_req.end;

}

PostCode();
