package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	EnvMongoURI    = "MONGO_URI"     // ชื่อ Mongo connection URI
	EnvMongoDBName = "MONGO_DB_NAME" // ชื่อ database

	defaultConnectTimeout = 10 * time.Second // timeout เชื่อมต่อ Mongo
	defaultPingTimeout    = 5 * time.Second  // timeout ping เพื่อตรวจว่า DB พร้อมใช้งาน
)

type IMongo interface {
	Close(ctx context.Context) error
	Collection(name string) *mongo.Collection
}

type Mongo struct {
	Client   *mongo.Client   // client สำหรับเชื่อมต่อ MongoDB
	Database *mongo.Database // database ที่เลือกไว้พร้อมใช้งาน
}

func NewMongoFromEnv(ctx context.Context) (*Mongo, error) {
	//1. อ่าน URI server
	uri := os.Getenv(EnvMongoURI)
	if uri == "" {
		return nil, fmt.Errorf("%s is required", EnvMongoURI)
	}

	//2. อ่านชื่อ database
	dbName := os.Getenv(EnvMongoDBName)
	if dbName == "" {
		return nil, fmt.Errorf("%s is required", EnvMongoDBName)
	}
	//3.call connect function
	return NewMongo(ctx, uri, dbName)
}

func NewMongo(ctx context.Context, uri, dbName string) (*Mongo, error) {
	if uri == "" {
		return nil, fmt.Errorf("mongo uri is required")
	}
	if dbName == "" {
		return nil, fmt.Errorf("mongo db name is required")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	//1. สร้าง context พร้อม timeout สำหรับขั้นตอน connect
	connectCtx, cancelConnect := context.WithTimeout(ctx, defaultConnectTimeout)
	defer cancelConnect()

	//2. เชื่อมต่อ Mongo ด้วย URI
	client, err := mongo.Connect(connectCtx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("connect mongo: %w", err)
	}

	//3. สร้าง context แยกสำหรับ ping ตรวจความพร้อมของ connection
	pingCtx, cancelPing := context.WithTimeout(ctx, defaultPingTimeout)
	defer cancelPing()

	//4. ตรวจว่าเชื่อมต่อไป
	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("ping mongo: %w", err)
	}

	//4. คืน object Mongo ที่มีทั้ง client และ database พร้อมใช้งาน
	return &Mongo{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (m *Mongo) Close(ctx context.Context) error {
	if m == nil || m.Client == nil {
		return nil
	}

	if ctx == nil {
		ctx = context.Background()
	}

	// ปิดการเชื่อมต่อ MongoDB
	if err := m.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("disconnect mongo: %w", err)
	}

	return nil
}

func (m *Mongo) Collection(name string) *mongo.Collection {
	return m.Database.Collection(name)
}
