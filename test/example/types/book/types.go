// Code generated by goctl. DO NOT EDIT.
package book

type BookRequest struct {
	Name string `json:"name"` // 姓名
	Age  int    `json:"age"`  // 年龄
}

type BookResponse struct {
	Code int    `json:"code"` // 业务码
	Msg  string `json:"msg"`  // 业务消息
}
