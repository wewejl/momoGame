package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "mmoGame/test/pd"
)

func main() {
	list := pb.Person{
		Name: "朱鑫烨",
		Age:  20,
		Phone: &pb.Person_PhoneNumber{
			Number: "1111111111111",
			Type:   1,
		},
		Phones: []*pb.Person_PhoneNumber{
			&pb.Person_PhoneNumber{
				Number: "***222222222222",
				Type:   2,
			},
			&pb.Person_PhoneNumber{
				Number: "***333333333333",
				Type:   3,
			},
			&pb.Person_PhoneNumber{
				Number: "***444444444444",
				Type:   4,
			},
		},
	}

	//编码, 使用protobuf提供库里面的编码方法
	//进行编码
	encodeinfo, err := proto.Marshal(&list)
	if err != nil {
		fmt.Println("proto.Marshal err", err)
		return
	}

	//使用grpc框架传输

	//解码
	lily1 := pb.Person{}
	err = proto.UnmarshalMerge(encodeinfo, &lily1)
	if err != nil {
		fmt.Println("proto.Unmarshal err:", err)
		return
	}

	fmt.Println("解码之后lily1:", lily1.GetName(), lily1.GetAge(), lily1.GetPhone(),lily1.Phones[0].GetNumber())
}
