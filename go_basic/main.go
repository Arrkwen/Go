package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/binary"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Arrkwen/Go/go_basic/myPrometheus"
	_ "github.com/go-sql-driver/mysql"
)

var MinVersion string

// var PublishDate string

func main2() {
	fmt.Printf("当前版本号:%s", MinVersion)
}

func main3() {
	myPrometheus.RealizePrometheus()
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func main() {
	path := "fil"
	fmt.Printf("the file is exist: %v\n", Exists(path))
}

type Girl struct {
	ImageId          string
	Name             string
	IsStranger       uint8
	TotalObjectCount int32
	Label            int64
}

type Boy struct {
	Name    string
	School  string
	Subject string
	Score   float32
}

type Document struct {
	label            int64
	imageId          string
	nameId           string
	isStranger       uint8
	totalObjectCount int32
}

func Uint64ToBytes(x uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, x)
	return b
}

func BytesToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

func Mashall() ([]byte, error) {
	// g := Girl{"sarina", "bb", 0, 1, 105812584719974407}
	// b := Boy{"xiaaming", "ustc", "ai", 95.4}
	// //这个没有MarshalIndent
	// gBytes, err := msgpack.Marshal(g)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// gLen := len(gBytes)
	// fmt.Printf("%d bytes", gLen)
	// lBytes := Uint64ToBytes(uint64(gLen))
	// var docBytes []byte
	// docBytes = append(docBytes, lBytes...)
	// docBytes = append(docBytes, gBytes...)
	// bBytes, err := msgpack.Marshal(b)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// bLen := uint64(len(bBytes))
	// lBytes = Uint64ToBytes(bLen)
	// docBytes = append(docBytes, lBytes...)
	// docBytes = append(docBytes, bBytes...)

	// write label
	doc := Document{
		label:            32,
		imageId:          "zhangzhang",
		nameId:           "xkkx",
		isStranger:       1,
		totalObjectCount: 23,
	}
	docBytes := Uint64ToBytes(uint64(doc.label))

	// write length of doc.imageId
	idLen := len(doc.imageId)
	lenBytes := Uint32ToBytes(uint32(idLen))
	docBytes = append(docBytes, lenBytes...)

	// write doc.imageId
	docBytes = append(docBytes, []byte(doc.imageId)...)

	// write length of doc.nameId
	idLen = len(doc.nameId)
	lenBytes = Uint32ToBytes(uint32(idLen))
	docBytes = append(docBytes, lenBytes...)

	// write doc.nameId
	docBytes = append(docBytes, []byte(doc.nameId)...)

	// write doc.isStranger
	strangerByte := Uint8ToBytes(doc.isStranger)
	docBytes = append(docBytes, strangerByte...)

	// write doc.totalObjectCount
	countBytes := Uint32ToBytes(uint32(doc.totalObjectCount))
	docBytes = append(docBytes, countBytes...)

	return docBytes, nil
}

func BytesToUint8(b []byte) uint8 {
	var x uint8
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &x)
	return x
}

func Uint16ToBytes(x uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, x)
	return b
}

func BytesToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func Uint32ToBytes(x uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, x)
	return b
}

func BytesToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func Unmarshall(byteData []byte) error {
	dataLen := uint64(len(byteData))
	var index uint64
	var nextIndex uint64
	for {

		// read and set label
		index = nextIndex
		nextIndex = index + 8
		if nextIndex > dataLen {
			break
		}
		labelbytes := byteData[index:nextIndex]
		label := int64(BytesToUint64(labelbytes))
		fmt.Println(label)
		// read length of doc.imageId
		index = nextIndex
		nextIndex = index + 4
		if nextIndex > dataLen {
			break
		}
		lenBytes := byteData[index:nextIndex]
		imageIdLen := BytesToUint32(lenBytes)

		// read and set doc.imageId
		index = nextIndex
		nextIndex = index + uint64(imageIdLen)
		if nextIndex > dataLen {
			break
		}
		imageIdBytes := byteData[index:nextIndex]
		fmt.Println(string(imageIdBytes))

		// read length of doc.nameId
		index = nextIndex
		nextIndex = index + 4
		if nextIndex > dataLen {
			break
		}
		lenBytes = byteData[index:nextIndex]
		nameIdLen := BytesToUint32(lenBytes)

		// read and set doc.nameId
		index = nextIndex
		nextIndex = index + uint64(nameIdLen)
		if nextIndex > dataLen {
			break
		}
		nameIdBytes := byteData[index:nextIndex]
		fmt.Println(string(nameIdBytes))

		// read and set doc.isStranger
		index = nextIndex
		nextIndex = index + 1
		if nextIndex > dataLen {
			break
		}
		isStrangerByte := byteData[index:nextIndex]
		fmt.Println(BytesToUint8(isStrangerByte) == 1)

		// read and set doc.totalObjectCount
		index = nextIndex
		nextIndex = index + 4
		if nextIndex > dataLen {
			break
		}
		countBytes := byteData[index:nextIndex]
		fmt.Println(int32(BytesToUint32(countBytes)))
		fmt.Println("END")
	}

	// for {
	// 	if index+8 > dataLen {
	// 		break
	// 	}
	// 	lenBytes := byteData[index : index+8]
	// 	cLen := BytesToUint64(lenBytes)
	// 	index += 8

	// 	// 读取数据girl
	// 	if index+cLen > dataLen {
	// 		break
	// 	}
	// 	gBytes := byteData[index : index+cLen]
	// 	index += cLen

	// 	// 解析数据girl
	// 	var g = Girl{}
	// 	if err := msgpack.Unmarshal(gBytes, &g); err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(g)

	// 	// 读取boy数据长度
	// 	if index+8 > dataLen {
	// 		break
	// 	}
	// 	lenBytes = byteData[index : index+8]
	// 	cLen = BytesToUint64(lenBytes)
	// 	index += 8

	// 	// 读取数据boy
	// 	if index+cLen > dataLen {
	// 		break
	// 	}
	// 	bBytes := byteData[index : index+cLen]
	// 	index += cLen

	// 	// 解析数据Boy
	// 	var b = Boy{}
	// 	if err := msgpack.Unmarshal(bBytes, &b); err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(b)
	// }
	// return nil
	return nil
}

func printLen(data []byte) {
	fmt.Println(string(data))
}

func interfaceT() {
	a := &A{
		name: "xiao",
	}
	// b := &B{
	// 	name: "kun",
	// }
	c := &C{
		persist: a,
	}
	PersistPrint(c)
}

func TimeDuration() {
	stringTime := "2022-03-25 18:32:58"

	loc, _ := time.LoadLocation("Local")

	t0, _ := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	t1 := time.Now()

	duration := t1.Sub(t0)
	fmt.Println(duration)
	if duration < time.Second*time.Duration(60) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}

type Meta struct {
	path string
	data []byte
}

func WriteSnapshotCluster(saveData []byte, seqID int64) error {
	savePath := "test.txt"
	meta := &Meta{
		path: savePath,
		data: saveData,
	}
	fmt.Println(meta)
	return nil
}

func Uint8ToBytes(x uint8) []byte {
	b := make([]byte, 1)
	b[0] = x
	return b
}

type Persist interface {
	Hello()
}

type A struct {
	name string
}

func (a *A) Hello() {
	fmt.Println(a.name)
}

func (a *A) NameA() {
	fmt.Println("A!")
}

type B struct {
	name string
}

func (a *B) Hello1() {
	fmt.Println(a.name)
}

func (a *B) NameB() {
	fmt.Println("B")
}

type C struct {
	persist Persist
}

func PersistPrint(obj *C) {
	obj.persist.Hello()
}

func testContext() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	callCtx1(ctx, &wg)
	// ctx2, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	// defer cancel()
	callCtx2(ctx, &wg)

	time.Sleep(time.Second * 6)
}

func callCtx1(ctx context.Context, wg *sync.WaitGroup) {
	go ctx1(ctx)
}

func callCtx2(ctx context.Context, wg *sync.WaitGroup) {
	go ctx2(ctx)
}

func ctx1(ctx context.Context) {
	fmt.Println("go into ctx1")
	for {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("ctx1", ctx.Err())
				return
			default:
				fmt.Printf("ctx1 deal time is %d\n", i)
			}
		}
	}
	fmt.Println("go out ctx1")

}

func ctx2(ctx context.Context) {
	fmt.Println("go into ctx2")
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("ctx2", ctx.Err())
			return
		default:
			fmt.Printf("ctx2 deal time is %d\n", i)
		}
	}
	fmt.Println("go out ctx2")
}

func GetTaskInfo(taskId string, db *sql.DB) (time.Time, error) {
	querySQL := fmt.Sprintf("SELECT update_time FROM %s WHERE task_id = ?", "task_info_list")

	stmt, _ := db.Prepare(querySQL)

	defer stmt.Close()

	row := stmt.QueryRow(taskId)

	var update_time time.Time

	row.Scan(&update_time)

	return update_time, nil
}
