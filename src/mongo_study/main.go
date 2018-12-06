package main

import (
	"fmt"
	"time"
)
import  "gopkg.in/mgo.v2/bson"
import  "gopkg.in/mgo.v2"


type RootUser struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	UID       uint32        `json:"uid" bson:"uid"`               // 用户ID
	Alias     string        `json:"alias" bson:"alias"`           // 别名
	CreatedAt time.Time     `json:"created_at" bson:"created_at"` // 创建时间
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"` // 更新时间
	Enabled   bool          `json:"enabled" bson:"enabled"`       // 是否启用
}

type IamUserGroup struct {
	Id             bson.ObjectId   `json:"id" bson:"_id"`
	RootProjectId  uint32          `json:"root_project_id" bson:"root_project_id"`   // 根project id
	OwnerProjectId uint32          `json:"owner_project_id" bson:"owner_project_id"` // 所属的project id
	Alias          string          `json:"alias" bson:"alias"`                       // 组别名
	Description    string          `json:"description" bson:"description"`           // 描述
	CreatedAt      time.Time       `json:"created_at" bson:"created_at"`             // 创建时间
	UpdatedAt      time.Time       `json:"updated_at" bson:"updated_at"`             // 更新时间
	Enabled        bool            `json:"enabled" bson:"enabled"`                   // 是否启用
	PolicyIds      []bson.ObjectId `json:"-" bson:"policy_ids"`                      // 策略IDs
}

type QRootUser struct {
	UID       *uint32        `json:"uid" bson:"uid"`               // 用户ID
}

var rus []RootUser



func main() {

	//update()
	//find()
	//updateall()


	//list := []uint32{1,2,3,4}
	//fmt.Printf("%v",list)

	//add_Policy()
    //pl, _ := add_Policy2("read_ob_test")
    //fmt.Println(pl)
	//if err != nil {
		bind_policy_for_user(1517381413,[]bson.ObjectId{bson.ObjectIdHex("5c076c0186b4e2207f000006")})
	//}


	//find_action_or()
	//add_url_bind_action()
	//group_by()
}

func updateall()  {

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mongo_test").C("table1")


	type user struct {
		UID       uint32        `json:"uid" bson:"uid"`               // 用户ID
		Alias     string        `json:"alias" bson:"alias"`           // 别名
	}

	c.DropCollection()
	c.EnsureIndex(mgo.Index{
		Key:[]string{"uid"},
		Unique:true,
	})

	u1 := user{1,"1"}
	u2 := user{2,"2"}
	u3 := user{3,"3"}
	u4 := user{3,"3"}

	c.Insert(u1)
	c.Insert(u2)
	c.Insert(u3)
	c.Insert(u4)


	type update struct {
		Alias *string `bson:"alias"`
	}

	s := "100"
	up := update{Alias:&s}

	err = c.Update(bson.M{"uid":bson.M{"$in":[]uint32{10}}},bson.M{"$set":up} )

	//fmt.Println(info)
	fmt.Println(err)
	//c.Find(bson.M{"uid":2}).Apply(mgo.Change{
	//	Upsert:    false,
	//	ReturnNew: false,
	//	Update:bson.M{"$set":bson.M{}},
	//},
	//nil)

	var all []user
	c.Find(bson.M{}).All(&all)
	fmt.Println(all)
}

func find()  {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("iam").C("action_url_action_relation")

	var urlAcQrn interface{}

	n,_ := c.Find(bson.M{}).Count()
	fmt.Println(n)
	err = c.Find(bson.M{
		"urls":bson.M{"$in":[]string{"haha"}},
	}).One(urlAcQrn)

	fmt.Println(err)
	fmt.Println(urlAcQrn)

	return
	query := bson.M{
		"alias":    "testgorup3",
		"owner_project_id": uint32(1380432009),
	}
	count,_:= c.Find(query).Count()
	fmt.Println(count)

	//ug := &IamUserGroup{}

	fmt.Println(err)


	var new_user RootUser

	c.Find(bson.M{"uid": 11}).Apply(mgo.Change{
		Upsert:    true,
		ReturnNew: true,
		Update: bson.M{
			"$setOnInsert": bson.M{
				"created_at": time.Now(),
			},
			"$currentDate": bson.M{
				"updated_at": true,
			},
		},
	}, &new_user)

	fmt.Println(new_user)
}

func update()  {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("mongo_test").C("table1")

	type user struct {
		UID       uint32        `json:"uid" bson:"uid"`               // 用户ID
		Alias     string        `json:"alias" bson:"alias"`           // 别名
	}

	//u1 := user{1,"1"}
	//u2 := user{2,"2"}
	//u3 := user{3,"3"}
	//
	//c.Insert(u1)
	//c.Insert(u2)
	//c.Insert(u3)


	var res user

	c.Find(bson.M{"alias":"2"}).Count()

	fmt.Println(res)

	c.Find(bson.M{}).
		Apply(mgo.Change{
		Upsert:    false,
		ReturnNew: true,
		Update: bson.M{"$set": bson.M{
			"enabled":    false,
		}},
	}, &res)

	fmt.Println(res)

	c.UpdateAll(bson.M{"$or":[]bson.M{ bson.M{"uid":bson.M{"$in":[]uint32{1}}},bson.M{"alias":"2"} }},bson.M{"$set": bson.M{"enabled":true,}})
}

var (
	policy_id1 = bson.NewObjectId()
	policy_id2 = bson.NewObjectId()
	policy_id3 = bson.NewObjectId()
	policy_id4 = bson.NewObjectId()
	policy_id5 = bson.NewObjectId()

	id1 = bson.ObjectIdHex("5a9fb77e43c8ce15db19fb21")
	id2 = bson.ObjectIdHex("5a9fb77e43c8ce15db19fb22")

)
func add_Policy(){

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("iam").C("policies")


	info,err := c.RemoveAll(
		bson.M{
			"root_uid":1810657124,
			"alias":bson.M{"$in":[]string{
				"test_read_kodo","test_read_kodo2","test_read_kodo3",
			}},

		})

	fmt.Println(info,err)

	err = c.Insert(
		bson.M{
			"_id":policy_id1,
			"root_uid":1810657124,
			"alias":"test_read_kodo",
			"description":"desc",
			"edit_type":2,
			"create_at":time.Now(),
			"updated_at":time.Now(),
			"enabled":true,
			"statement":[]bson.M{
				{
					"action":[]string{"kodo/ReadObject"},
					"resource":[]string{"qrn:kodo:z0:1810657124:b2"},
					"effect":"Allow",
				},
			},
		},


		bson.M{
			"_id":policy_id2,
			"root_uid":1810657124,
			"alias":"test_read_kodo2",
			"description":"desc",
			"edit_type":2,
			"create_at":time.Now(),
			"updated_at":time.Now(),
			"enabled":true,
			"statement":[]bson.M{
				{
					"action":[]string{"kodo/ReadObject"},
					"resource":[]string{"qrn:kodo:z0:1810657124:b1"},
					"effect":"Allow",
				},
			},
		},


		bson.M{
			"_id":policy_id3,
			"root_uid":1810657124,
			"alias":"test_read_kodo3",
			"description":"desc",
			"edit_type":2,
			"create_at":time.Now(),
			"updated_at":time.Now(),
			"enabled":true,
			"statement":[]bson.M{
				{
					"action":[]string{"kodo/ReadObject"},
					"resource":[]string{"*"},
					"effect":"Allow",
				},
			},
		},


		bson.M{
			"_id":policy_id4,
			"root_uid":1810657124,
			"alias":"test_create_kodo1",
			"description":"desc",
			"edit_type":2,
			"create_at":time.Now(),
			"updated_at":time.Now(),
			"enabled":true,
			"statement":[]bson.M{
				{
					"action":[]string{"kodo/CreateObject"},
					"resource":[]string{"*"},
					"effect":"Allow",
				},
			},
		},

		bson.M{
			"_id":policy_id5,
			"root_uid":1810657124,
			"alias":"test_read_kodo2",
			"description":"desc",
			"edit_type":2,
			"create_at":time.Now(),
			"updated_at":time.Now(),
			"enabled":true,
			"statement":[]bson.M{
				{
					"action":[]string{"kodo/CreateObject"},
					"resource":[]string{"qrn:kodo:z0:1810657124:b2"},
					"effect":"Allow",
				},
			},
		},
		)
	fmt.Println(err)

}



func add_Policy2(alias string ) (id bson.ObjectId, err error){

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("iam").C("policies")

	id = bson.NewObjectId()
	_,err = c.Upsert(

		bson.M{
			"alias":alias,
		},

		bson.M{
			"_id":id,
			"root_uid":1810657124,
			"alias":alias,
			"description":"desc",
			"edit_type":2,
			"create_at":time.Now(),
			"updated_at":time.Now(),
			"enabled":true,
			"statement":[]bson.M{
				{
					"action":[]string{"kodo/ReadObject"},
					"resource":[]string{"qrn:kodo:*:1810657124:b2"},
					"effect":"Allow",
				},
			},
		},
	)

	return id,err
}


func bind_policy_for_user(iuid uint32, policyids []bson.ObjectId)  {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("iam").C("users")

	for _,id:= range policyids  {

		err = c.Update(
			bson.M{
				"iuid":iuid,
			},

			bson.M{
				"$addToSet":bson.M{
					"policy_ids":id,
				},
			},
		)

		fmt.Println(err,id)

	}


}

func find_action_or()  {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("iam").C("url_bind_action")


	count,err := c.Find(
		bson.M{
			"url":bson.M{"$in":[]string{"POST:/upload","POST:/upload/"}},
		},
	).Count()

	fmt.Println(err)
	fmt.Println(count)
}

func add_url_bind_action()  {

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("iam").C("url_bind_action")

	if err = c.EnsureIndex(mgo.Index{
		Key:    []string{"url"},
		Unique: true,
	}); err != nil {
		return
	}

	err = c.DropCollection()
	if err!=nil{
		fmt.Println(err)
		return
	}
	err = c.Insert(

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"GET:",
			"action":"kodo/ReadObject",
			"desc":"下载文件",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"GET:/buckets",
			"action":"kodo/ListBuckets",
			"desc":"列举一个账号的所有空间",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/mkbucketv2",
			"action":"kodo/CreateBucket",
			"desc":"创建存储空间，同时绑定一个七牛二级域名，用于访问资源",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"GET:/v6/domain/list",
			"action":"kodo/ReadBucket",
			"desc":"获取一个空间绑定的域名列表",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/drop",
			"action":"kodo/ReadBucket",
			"desc":"删除指定存储空间",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/upload",
			"action":"kodo/CreateObject",
			"desc":"用于在一次 HTTP 会话中上传单一的一个文件",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},


		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/mkblk",
			"action":"kodo/CreateObject",
			"desc":"为后续分片上传创建一个新的块，同时上传第一片数据",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/bput",
			"action":"kodo/CreateObject",
			"desc":"上传指定块的一片数据，具体数据量可根据现场环境调整，同一块的每片数据必须串行上传",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},


		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/mkfile",
			"action":"kodo/CreateObject",
			"desc":"将上传好的所有数据块按指定顺序合并成一个资源文件",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"GET:/stat",
			"action":"kodo/ReadObject",
			"desc":"仅获取资源的 Metadata 信息，不返回资源内容",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/chgm",
			"action":"kodo/ModifyObject",
			"desc":"修改文件的 MIME 类型信息",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/move",
			"action":"kodo/ReadObject and DeleteObject for src, CreateObject for dest",
			"desc":"将源空间的指定资源移动到目标空间，或在同一空间内对资源重命名",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/copy",
			"action":"kodo/ReadObject for src, CreateObject for dest",
			"desc":"将指定资源复制为新命名资源",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/delete",
			"action":"kodo/DeleteBucket for bucket, DeleteObject for object",
			"desc":"删除指定资源",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"GET:/list",
			"action":"kodo/ListObjects",
			"desc":"用于列举指定空间里的所有文件条目	",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/fetch",
			"action":"kodo/CreateObject",
			"desc":"从指定 URL 抓取资源，并将该资源存储到指定空间中。每次只抓取一个文件，抓取时可以指定保存空间名和最终资源名",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/batch",
			"action":"kodo/Various",
			"desc":"指在单一请求中执行多次获取元信息、移动、复制、删除操作，极大提高资源管理效率",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/prefetch",
			"action":"kodo/CreateObject",
			"desc":"对于设置了镜像存储的空间，从镜像源站抓取指定名称的资源并存储到该空间中",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},


		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/image",
			"action":"kodo/CreateObject",
			"desc":"为存储空间指定一个镜像回源网址，用于取回资源",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},


		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/deleteAfterDays",
			"action":"kodo/ModifyObject",
			"desc":"设置指定资源的生命周期，即设置一个文件多少天后删除",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},


		bson.M{
			"_id":bson.NewObjectId(),
			"url":"POST:/chtype",
			"action":"kodo/ModifyObject",
			"desc":"修改文件的存储类型信息，即低频存储和标准存储的互相转换",
			"service":"kodo",
			"create_at":time.Now(),
			"update_at":time.Now(),
		},

	)


	fmt.Println(err)

}

func group_by()  {

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("iam").C("actions2")

	type Action struct {
		Id        bson.ObjectId     `json:"id" bson:"_id"`
		Url       string        	`json:"url" bson:"url"`               // url格式   ex:POST:/upload
		Name      string            `json:"name" bson:"name"`             // 动作名
		Alias     string            `json:"alias" bson:"alias"`           // 动作别名
		Service   string            `json:"service" bson:"service"`       // 服务（产品）
		CreatedAt time.Time         `json:"created_at" bson:"created_at"` // 创建时间
		UpdatedAt time.Time         `json:"updated_at" bson:"updated_at"` // 更新时间
		Enabled   bool              `json:"enabled" bson:"enabled"`       // 是否启用
	}
	var s []Action

	c.Find(bson.M{}).All(&s)

	fmt.Println(len(s))
}