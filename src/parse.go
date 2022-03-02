package main

import (
	"fmt"
	"github.com/miekg/dns"
)

func main() {
	//dns := DNSServer{}
	var origin DNSServer
	origin.readFromConfig()

	var message dns.Msg

	//域名转化为fqdn
	fqdn := dns.Fqdn(origin.Domain)

	//设置询问消息体,请求解析A记录
	message.SetQuestion(fqdn, dns.TypeA)

	//发送消息体
	result, err := dns.Exchange(&message, origin.DNSAddress)

	if err != nil {
		panic(err)
	}

	if len(result.Answer) < 1 {
		panic("no records")
		return
	}

	for _, answer := range result.Answer {
		a, ok := answer.(*dns.A)
		if ok {
			fmt.Println(a.A)
		}
	}
}
