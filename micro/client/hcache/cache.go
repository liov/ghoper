package hcache

import (
	"fmt"
	"github.com/golang/groupcache"
	pb "github.com/golang/groupcache/groupcachepb"
	"micro/protobuf"
	"net/http"
)

const (
	BasePath   string = "/basepath/"
	groupName         = "helloworld"
	cacheBytes        = 1024 * 1024 * 1024 * 16
)

func Start() {

	port := ":8086"
	me := "http://localhost:8086"

	opts := groupcache.HTTPPoolOptions{BasePath: BasePath}
	peers := groupcache.NewHTTPPoolOpts(me, &opts)

	peers.Set("http://localhost:8333", "http://localhost:8222")
	groupcache.NewGroup(groupName, cacheBytes, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			dest.SetProto(&protobuf.User{Name: "cache"})
			return nil
		}))

	http.HandleFunc(BasePath, peers.ServeHTTP)
	http.ListenAndServe(port, nil)

}

func GetFromPeer(groupName, key string, peers *groupcache.HTTPPool) (value []byte, err error) {
	req := &pb.GetRequest{Group: &groupName, Key: &key}
	res := &pb.GetResponse{}

	peer, ok := peers.PickPeer(key)
	if ok == false {
		fmt.Println("peers PickPeer failed: ", key)
		return
	}

	err = peer.Get(nil, req, res)
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}
