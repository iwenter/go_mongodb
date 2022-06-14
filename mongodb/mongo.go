package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type student struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 指定获取要操作的数据集
	collection := client.Database("xiaobo").Collection("student")
	//find, err := collection.Find(context.TODO(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(find)

	// 插入一条数据
	//var s1 = student{Name: "小红", Age: 18}
	//insertOne, err := collection.InsertOne(context.TODO(), s1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("InsertOne:", insertOne.InsertedID)

	// 插入多条数据
	//s2 := student{Name: "小兰", Age: 19}
	//s3 := student{Name: "小黑", Age: 22}
	//
	//students := []interface{}{s2, s3}
	//insertMany, err := collection.InsertMany(context.TODO(), students)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("InsertMany:", insertMany.InsertedIDs)

	// 更新数据
	//filter := bson.D{{"name", "小红"}}
	//update := bson.D{{"$inc", bson.D{{"age", 13}}}}
	//
	//updateOne, err := collection.UpdateOne(context.TODO(), filter, update)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Matched %d documents and updated %v documents.\n", updateOne.MatchedCount, updateOne.ModifiedCount)

	// 查找一条数据
	// 创建一个Student变量用来接收查询的结果
	//var result student
	//err = collection.FindOne(context.TODO(), bson.D{{"name", "小红"}}).Decode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Found a single document: %+v\n", result)

	// 查找多条数据
	//findOptions := options.Find()
	//findOptions.SetLimit(2)
	//// 创建一个切片用来接收查询的结果
	//var results []*student
	//
	//// 把bson.D{{}} 作为一个filter来匹配所有文档
	//cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//// 查找多个文档返回一个光标
	//// 遍历游标允许我们一次解码一个文档
	//for cur.Next(context.TODO()) {
	//	// 创建一个空的Student变量用来接收查询的结果
	//	var elem student
	//	// 将查询的结果解码到Student变量中
	//	err := cur.Decode(&elem)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	results = append(results, &elem)
	//}
	//if err := cur.Err(); err != nil {
	//	log.Fatal(err)
	//}
	//// 关闭游标
	//cur.Close(context.TODO())
	//fmt.Printf("Found multiple documents (limit 2): %#v\n", results)

	// 删除数据
	// 删除名字是小黄的哪个
	deleteOne, err := collection.DeleteOne(context.TODO(), bson.D{{"name", "小黄"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteOne.DeletedCount)

	//// 删除所有数据
	//deleteMany, err := collection.DeleteMany(context.TODO(), bson.D{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Deleted %v documents in the trainers collection\n", deleteMany.DeletedCount)

	// 关闭连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}

// 连接池模式
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {

	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	applyURI := options.Client().ApplyURI(uri)
	applyURI.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, applyURI)
	if err != nil {
		return nil, err
	}

	return client.Database(name), nil

}
