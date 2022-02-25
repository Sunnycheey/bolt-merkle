package rpc

import (
	"bytes"
	"context"
	"fmt"
	bolt "go.etcd.io/bbolt"
	pb "go.etcd.io/bbolt/rpc/service/proto"
)

type Server struct {
	db *bolt.DB
	readTx *bolt.Tx
	bucketName string
	path string
}

var db *bolt.DB


func (s *Server) mustEmbedUnimplementedStorageServiceServer() {
	panic("implement me")
}

func (s *Server) Set (ctx context.Context, requestKv *pb.RequestKV) (*pb.Status, error) {
	//s.db, _ = bolt.Open(s.path, 0600, nil)
	writeTx, err := db.Begin(true)
	kv := requestKv.GetKv()
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	//b, _ := writeTx.CreateBucketIfNotExists([]byte("test"))
	b, _ := writeTx.CreateBucketIfNotExists(requestKv.GetBucketName())
	err = b.Put(kv.GetKey(), kv.GetVal())
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	err = writeTx.Commit()
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}

func (s *Server) SetBatch (ctx context.Context, requestKVs *pb.RequestKVs) (*pb.Status, error) {
	//s.db, _ = bolt.Open(s.path, 0600, nil)
	writeTx, err := db.Begin(true)
	kvs := requestKVs.GetKvs()
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	b, _ := writeTx.CreateBucketIfNotExists(requestKVs.GetBucketName())
	for _, kv := range kvs.GetKvs() {
		err := b.Put(kv.GetKey(), kv.GetVal())
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
	if b != nil {
		defer readTx.Rollback()
		v := b.Get(k.GetKey())
		return &pb.ReturnVal{Val: v, S: &pb.Status{Code: 0, Msg: "Success"}}, nil
	}
	return &pb.ReturnVal{Val: nil, S: &pb.Status{Code: 30000, Msg: "Bucket Not Exists"}}, nil
}

func (s *Server) GetBatch(ctx context.Context, ks *pb.RequestKeys) (*pb.ReturnVals, error) {
	//s.db, _ = bolt.Open(s.path, 0600, nil)
	readTx, _ := db.Begin(false)
	defer readTx.Rollback()
	b := readTx.Bucket(ks.GetBucketName())
	if b != nil {
		vals := make([][]byte, 0)
		for _, k := range ks.GetKeys() {
			v := b.Get(k)
			vals = append(vals, v)
		}
		status := &pb.Status{Code: 0, Msg: "Success"}
		return &pb.ReturnVals{Val: vals, S: status}, nil
	}
	return &pb.ReturnVals{Val: nil, S: &pb.Status{Code: 30000, Msg: "Bucket Not Exists"}}, nil
}

func (s *Server) InitDatabase(ctx context.Context, para *pb.InitParam) (*pb.Status, error) {
	d, err := bolt.Open(para.GetPath(), 0600, nil)
	db = d
	bucketName := para.GetBucketName()
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	writeTx, err := db.Begin(true)
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	// create bucket
	_, err = writeTx.CreateBucketIfNotExists(bucketName)
	//defer s.readTx.Rollback()
	if err != nil {
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	if err != nil {
		fmt.Println("commit error!")
		return &pb.Status{Code: uint32(1), Msg: err.Error()}, err
	}
	writeTx.Commit()
	//defer s.db.Close()
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
				if bytes.Compare(k[s:e], rv.GetStart()[s:e]) == -1 || bytes.Compare(k[s:e], rv.GetEnd()[s:e]) == 1{
					flag = true
					break
				}
				s = s + len(rv.GetStart())
			}
			if !flag {
				kvs = append(kvs, &pb.KV{Key: k, Val: v})
			}
		}
		status := &pb.Status{Code: 0, Msg: "Success"}
		return &pb.ReturnKVs{Kvs: kvs, S: status}, nil
	}
	return &pb.ReturnKVs{Kvs: nil, S: &pb.Status{Code: 30000, Msg: "Bucket Not Exists"}}, nil

}

func (s *Server) Del(ctx context.Context, key *pb.RequestKey) (*pb.Status, error) {
	writeTx, _ := db.Begin(true)
	b, _ := writeTx.CreateBucketIfNotExists(key.GetBucketName())
	b.Delete(key.GetKey())
	writeTx.Commit()
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}

func (s *Server) DelBatch(ctx context.Context, keys *pb.RequestKeys) (*pb.Status, error) {
	writeTx, _ := db.Begin(true)
	b, _ := writeTx.CreateBucketIfNotExists(keys.GetBucketName())
	for _, k := range keys.GetKeys() {
		b.Delete(k)
	}
	writeTx.Commit()
	return &pb.Status{Code: 0, Msg: "Success"}, nil
}