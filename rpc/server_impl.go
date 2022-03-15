package rpc

import (
	"bytes"
	"context"
	bolt "go.etcd.io/bbolt"
	pb "go.etcd.io/bbolt/rpc/service/proto"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

type Server struct {
	db         *bolt.DB
	readTx     *bolt.Tx
	bucketName string
	path       string
}

var db *bolt.DB

func (s *Server) mustEmbedUnimplementedStorageServiceServer() {
	panic("implement me")
}

func (s *Server) Set(ctx context.Context, requestKv *pb.RequestKV) (*pb.Status, error) {
	writeTx, err := db.Begin(true)
	kv := requestKv.GetKv()
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	//b, _ := writeTx.CreateBucketIfNotExists([]byte("test"))
	log.Printf("Handle <Set>, key: %s, value: %s", kv.GetKey(), kv.GetVal().String())
	b, _ := writeTx.CreateBucketIfNotExists(requestKv.GetBucketName())
	v, err := proto.Marshal(kv.GetVal())
	if err != nil {
		panic("Error at marshal")
	}
	err = b.Put(kv.GetKey(), v)
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	err = writeTx.Commit()
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}

func (s *Server) SetBatch(ctx context.Context, requestKVs *pb.RequestKVs) (*pb.Status, error) {
	//s.db, _ = bolt.Open(s.path, 0600, nil)
	writeTx, err := db.Begin(true)
	kvs := requestKVs.GetKvs()
	//log.Printf("Handle <SetBatch>, key value sequences: %s", kvs.GetKvs())
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	b, _ := writeTx.CreateBucketIfNotExists(requestKVs.GetBucketName())
	for _, kv := range kvs.GetKvs() {
		v, err := proto.Marshal(kv.GetVal())
		if err != nil {
			panic("Error at marshal!")
		}
		err = b.Put(kv.GetKey(), v)
		if err != nil {
			return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
		}
	}
	writeTx.Commit()
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}

func (s *Server) Get(ctx context.Context, k *pb.RequestKey) (*pb.ReturnVal, error) {
	readTx, _ := db.Begin(false)
	b := readTx.Bucket(k.GetBucketName())
	log.Printf("Handle <Get>, key: %s", k.GetKey())
	if b != nil {
		defer readTx.Rollback()
		v := b.Get(k.GetKey())
		vval := &pb.VValue{}
		err := proto.Unmarshal(v, vval)
		if err != nil {
			panic("Unmarshal failed!")
		}
		return &pb.ReturnVal{Val: vval, S: &pb.Status{Code: 0, Msg: "Success"}}, nil
	}
	return &pb.ReturnVal{Val: nil, S: &pb.Status{Code: 30000, Msg: "Bucket Not Exists"}}, nil
}

func (s *Server) GetBatch(ctx context.Context, ks *pb.RequestKeys) (*pb.ReturnVals, error) {
	//s.db, _ = bolt.Open(s.path, 0600, nil)
	readTx, _ := db.Begin(false)
	defer readTx.Rollback()
	b := readTx.Bucket(ks.GetBucketName())
	//log.Printf("Handle <GetBatch>, key sequence: %s", ks.GetKeys())
	if b != nil {
		vals := make([]*pb.VValue, 0)
		for _, k := range ks.GetKeys() {
			v := b.Get(k)
			vval := &pb.VValue{}
			err := proto.Unmarshal(v, vval)
			if err != nil {
				panic("Unmarshal failed!")
			}
			vals = append(vals, vval)
		}
		status := &pb.Status{Code: 0, Msg: "Success"}
		return &pb.ReturnVals{Val: vals, S: status}, nil
	}
	return &pb.ReturnVals{Val: nil, S: &pb.Status{Code: 30000, Msg: "Bucket Not Exists"}}, nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (s *Server) InitDatabase(ctx context.Context, para *pb.InitParam) (*pb.Status, error) {
	log.SetOutput(os.Stdout)
	log.Printf("Handle <init>, database path: %s", para.GetPath())
	exists, err := FileExists(para.GetPath())
	if err != nil {
		panic(err)
	}
	if exists && para.GetDeleteIfExists() {
		os.Remove(para.GetPath())
	}
	d, err := bolt.Open(para.GetPath(), 0600, nil)
	db = d
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	s.path = para.GetPath()
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}

func (s *Server) GetRootHash(ctx context.Context, bucket *pb.Bucket) (*pb.RootHash, error) {
	readTx, _ := db.Begin(false)
	b := readTx.Bucket(bucket.GetBucketName())
	if b != nil {
		defer readTx.Rollback()
		hash := b.GetRootHash()
		status := &pb.Status{Code: 0, Msg: "Success"}
		return &pb.RootHash{S: status, Rh: hash}, nil
	}
	return &pb.RootHash{Rh: nil, S: &pb.Status{Code: 30000, Msg: "Bucket Not Exists"}}, nil
}

func (s *Server) CloseDB(ctx context.Context, empty *pb.Empty) (*pb.Status, error) {
	log.Printf("Handle <CloseDB>")
	err := db.Close()
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}

func (s *Server) RangeQuery(ctx context.Context, arrayRangeKey *pb.ArrayRangeKey) (*pb.ReturnKVs, error) {
	readTx, err := db.Begin(false)
	defer readTx.Rollback()
	if err != nil {
		panic(err)
	}
	b := readTx.Bucket(arrayRangeKey.GetBucketName())
	log.Printf("Handle <RangeQuery>, range sequence: %s", arrayRangeKey.GetRangeKey())
	if b != nil {
		c := b.Cursor()
		kvs := make([]*pb.KV, 0)
		start, end := make([]byte, 0), make([]byte, 0)
		for _, k := range arrayRangeKey.GetRangeKey() {
			start = append(start, k.GetStart()...)
			end = append(end, k.GetEnd()...)
		}
		for k, v := c.Seek(start); k != nil && bytes.Compare(k, end) <= 0; k, v = c.Next() {
			flag := false
			s, e := 0, 0
			for _, rv := range arrayRangeKey.GetRangeKey() {
				e = e + len(rv.GetStart())
				// k[s:e] < start || k[s:e] > end
				if bytes.Compare(k[s:e], rv.GetStart()[s:e]) == -1 || bytes.Compare(k[s:e], rv.GetEnd()[s:e]) == 1 {
					flag = true
					break
				}
				s = s + len(rv.GetStart())
			}
			if !flag {
				vval := &pb.VValue{}
				err := proto.Unmarshal(v, vval)
				if err != nil {
					panic("Unmarshal failed!")
				}
				kvs = append(kvs, &pb.KV{Key: k, Val: vval})
			}
		}
		status := &pb.Status{Code: 0, Msg: "Success"}
		return &pb.ReturnKVs{Kvs: kvs, S: status}, nil
	}
	return &pb.ReturnKVs{Kvs: nil, S: &pb.Status{Code: 30000, Msg: "Bucket Not Exists"}}, nil

}

func (s *Server) Del(ctx context.Context, key *pb.RequestKey) (*pb.Status, error) {
	log.Printf("Handle <Del>, key: %s", key.GetKey())
	writeTx, _ := db.Begin(true)
	b, _ := writeTx.CreateBucketIfNotExists(key.GetBucketName())
	b.Delete(key.GetKey())
	writeTx.Commit()
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}

func (s *Server) DelBatch(ctx context.Context, keys *pb.RequestKeys) (*pb.Status, error) {
	log.Printf("Handle <DelBatch>, key sequence: %s", keys.GetKeys())
	writeTx, _ := db.Begin(true)
	b, _ := writeTx.CreateBucketIfNotExists(keys.GetBucketName())
	for _, k := range keys.GetKeys() {
		b.Delete(k)
	}
	writeTx.Commit()
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}
