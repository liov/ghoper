import axios from "axios";

const cookieparser = process.server ? require('cookieparser') : undefined;

export default function ({app: {router}, store, error, req, redirect}) {
    if (!store.state.user) {
        if (process.server) {
            if (req.headers.cookie) {
                const parsed = cookieparser.parse(req.headers.cookie);
                if (parsed.token) {
                    axios.defaults.headers.common['Cookie'] = req.headers.cookie;
                    axios.get(`/api/user/get`).then((res) => {
                        if (res.data.msg === 'ok') {
                            let user = res.data.data;
                            //console.log(user);
                            store.commit('SET_USER', user)
                        }
                    });
                }
            } else {
                /*            error({
                                message: 'You are not connected',
                                statusCode: 403
                            });*/
                redirect({path: '/user/signin'});
            }
        } else {
            /*            error({
                            message: 'You are not connected',
                            statusCode: 403
                        });*/
            redirect({path: '/user/signin'});
        }

    }
}
