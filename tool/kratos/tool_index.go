package main

import "time"

var toolIndexs = []*Tool{
	{
		Name:      "kratos",
		Alias:     "kratos",
		BuildTime: time.Date(2019, 6, 21, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-kratos/kratos/tool/kratos",
		Summary:   "Kratos工具集本体",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "kratos",
		Hidden:    true,
	},
	{
		Name:      "protoc",
		Alias:     "kratos-protoc",
		BuildTime: time.Date(2019, 10, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-kratos/kratos/tool/kratos-protoc",
		Summary:   "快速方便生成pb.go的protoc封装，windows、Linux请先安装protoc工具",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "kratos",
	},
	{
		Name:      "swagger",
		Alias:     "swagger",
		BuildTime: time.Date(2019, 5, 5, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-swagger/go-swagger/cmd/swagger",
		Summary:   "swagger api文档",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "goswagger.io",
	},
	{
		Name:      "genbts",
		Alias:     "kratos-gen-bts",
		BuildTime: time.Date(2019, 10, 31, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-kratos/kratos/tool/kratos-gen-bts",
		Summary:   "缓存回源逻辑代码生成器",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "kratos",
	},
	{
		Name:      "genmc",
		Alias:     "kratos-gen-mc",
		BuildTime: time.Date(2019, 7, 23, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/go-kratos/kratos/tool/kratos-gen-mc",
		Summary:   "mc缓存代码生成",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "kratos",
	},
	{
		Name:      "wire",
		Alias:     "wire",
		BuildTime: time.Date(2019, 8, 20, 0, 0, 0, 0, time.Local),
		Install:   "go get -u github.com/google/wire/cmd/wire",
		Platform:  []string{"darwin", "linux", "windows"},
		Hidden:    true,
	},
	{
		Name:         "genproject",
		Alias:        "kratos-gen-project",
		Install:      "go get -u github.com/go-kratos/kratos/tool/kratos-gen-project",
		BuildTime:    time.Date(2019, 12, 22, 0, 0, 0, 0, time.Local),
		Platform:     []string{"darwin", "linux", "windows"},
		Hidden:       true,
		Requirements: []string{"wire"},
	},
	{
		Name:      "testgen",
		Alias:     "testgen",
		BuildTime: time.Date(2020, 1, 20, 22, 33, 0, 0, time.Local),
		Install:   "go get -u github.com/go-kratos/kratos/tool/testgen",
		Summary:   "测试代码生成",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "kratos",
	},
	{
		Name:      "testcli",
		Alias:     "testcli",
		BuildTime: time.Date(2020, 1, 20, 22, 33, 0, 0, time.Local),
		Install:   "go get -u github.com/go-kratos/kratos/tool/testcli",
		Summary:   "测试代码运行",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "kratos",
	},
}
