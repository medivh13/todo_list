package nats_publisher

import (
	"fmt"
	"log"

	"todo_list/src/infra/broker/nats"
)

// PublisherInterface mendefinisikan kontrak untuk publisher NATS
type PublisherInterface interface {
	Nats(data []byte, subject string) error // Method untuk publish pesan ke NATS
}

// PushWorkerImpl adalah implementasi dari PublisherInterface
type PushWorkerImpl struct {
	nats *nats.Nats // Menyimpan instance koneksi NATS
}

// NewPushWorker membuat instance baru dari PushWorkerImpl
func NewPushWorker(Nats *nats.Nats) PublisherInterface {
	return &PushWorkerImpl{nats: Nats}
}

// Nats mengirim pesan ke subject tertentu di NATS
func (p *PushWorkerImpl) Nats(data []byte, subject string) error {
	// Pastikan koneksi NATS sudah terhubung
	if p.nats == nil || p.nats.Conn == nil || !p.nats.Conn.IsConnected() {
		return fmt.Errorf("NATS connection is not established")
	}

	// Mengirim pesan ke subject yang ditentukan
	if err := p.nats.Conn.Publish(subject, data); err != nil {
		return err
	}

	// Memastikan pesan telah dikirim
	if err := p.nats.Conn.Flush(); err != nil {
		return err
	}

	// Mengecek apakah ada error terakhir yang terjadi pada koneksi NATS
	if err := p.nats.Conn.LastError(); err != nil {
		return err
	}

	log.Printf("Published to [%s]\n", subject) // Logging informasi publikasi

	return nil
}
