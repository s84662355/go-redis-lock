package main

import(
	"goredislock"
	"github.com/go-redis/redis"
	"fmt"
	//"strconv"
	)

var client = redis.NewClient(&redis.Options{Addr:"127.0.0.1:6379",Password: "", DB:0})

 

 

func main(){
	key:="11 5gh76"
	lock := goredislock.NewWLock( client , "a776hg6fsh")
	if( lock.Lock(key, 20 ,1000) ){
		fmt.Println("bbbbbbbbb")

	    fmt.Println(lock.UnLock(key))
		return
	}
	fmt.Println("ffffffff")
}