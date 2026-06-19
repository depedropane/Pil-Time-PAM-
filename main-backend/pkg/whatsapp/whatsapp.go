package whatsapp

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

var waClient *whatsmeow.Client

// InitWA menginisialisasi koneksi ke WhatsApp via whatsmeow
func InitWA(dsn string) {
	// Gunakan logger whatsmeow
	dbLog := waLog.Stdout("Database", "WARN", true)
	
	// Format DSN mungkin perlu disesuaikan jika menggunakan sslmode dll.
	// dsn dari gorm sudah cocok untuk postgres stdlib, tapi lib/pq membutuhkan format yg sesuai
	container, err := sqlstore.New(context.Background(), "postgres", dsn, dbLog)
	if err != nil {
		log.Fatalf("[WhatsApp] Gagal terhubung ke database sesi: %v", err)
	}

	// Ambil device pertama dari database jika ada, jika tidak, akan dibuatkan session baru.
	deviceStore, err := container.GetFirstDevice(context.Background())
	if err != nil {
		log.Fatalf("[WhatsApp] Gagal mendapatkan device: %v", err)
	}

	clientLog := waLog.Stdout("Client", "WARN", true)
	waClient = whatsmeow.NewClient(deviceStore, clientLog)

	if waClient.Store.ID == nil {
		// Belum login, minta generate QR Code
		qrChan, _ := waClient.GetQRChannel(context.Background())
		err = waClient.Connect()
		if err != nil {
			log.Fatalf("[WhatsApp] Gagal connect ke server WA: %v", err)
		}
		
		go func() {
			for evt := range qrChan {
				if evt.Event == "code" {
					fmt.Println("\n\033[1;32m=== SCAN QR CODE INI MENGGUNAKAN WHATSAPP ANDA ===\033[0m")
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					fmt.Println("\033[1;32m==================================================\033[0m")
				} else {
					fmt.Printf("[WhatsApp] Login event: %s\n", evt.Event)
				}
			}
		}()
	} else {
		// Sudah pernah login sebelumnya
		err = waClient.Connect()
		if err != nil {
			log.Fatalf("[WhatsApp] Gagal connect ke server WA: %v", err)
		}
		fmt.Println("[WhatsApp] Berhasil terhubung ke WhatsApp dengan session yang tersimpan.")
	}
}

// SendWarning mengirim pesan peringatan via WhatsMeow
func SendWarning(toPhone, patientName, medicineName, scheduledTime string) {
	if waClient == nil || !waClient.IsConnected() {
		fmt.Println("\033[1;31m[WhatsApp Error] Client belum diinisialisasi atau tidak terhubung.\033[0m")
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05 MST")
	border := "=========================================================================="
	fmt.Println()
	fmt.Println(border)
	fmt.Printf("\033[1;32m[SISTEM OTOMATIS: PENGIRIMAN PERINGATAN WHATSAPP]\033[0m\n")
	fmt.Printf("Waktu Proses : %s\n", timestamp)
	fmt.Printf("Tujuan (Pendamping WA) : \033[1;36m%s\033[0m\n", toPhone)
	fmt.Printf("Nama Pasien            : \033[1m%s\033[0m\n", patientName)
	fmt.Printf("Nama Obat              : \033[1;33m%s\033[0m\n", medicineName)
	fmt.Printf("Jadwal Minum           : %s WIB\n", scheduledTime)
	fmt.Println("--------------------------------------------------------------------------")

	messageContent := fmt.Sprintf(
		"Halo Pendamping,\n\n"+
			"Peringatan! Pasien atas nama *%s* belum menandai jadwal minum obat "+
			"*%s* yang dijadwalkan pada pukul *%s WIB* hari ini.\n\n"+
			"Mohon ingatkan pasien untuk segera meminum obatnya sebelum batas waktu habis.\n\n"+
			"Terima kasih,\n"+
			"Aplikasi Pil-Time",
		patientName, medicineName, scheduledTime,
	)

	// Format nomor telepon (harus format JID: 628xxx@s.whatsapp.net)
	cleanPhone := toPhone
	cleanPhone = strings.ReplaceAll(cleanPhone, " ", "")
	cleanPhone = strings.ReplaceAll(cleanPhone, "-", "")
	cleanPhone = strings.ReplaceAll(cleanPhone, "+", "")
	if strings.HasPrefix(cleanPhone, "0") {
		cleanPhone = "62" + cleanPhone[1:]
	}

	jid := types.NewJID(cleanPhone, types.DefaultUserServer)

	// Struktur pesan di whatsmeow
	msg := &waE2E.Message{
		Conversation: proto.String(messageContent),
	}

	_, err := waClient.SendMessage(context.Background(), jid, msg)
	if err != nil {
		fmt.Printf("\033[1;31m[GATEWAY ERROR] Gagal mengirim pesan ke %s: %v\033[0m\n", toPhone, err)
	} else {
		fmt.Printf("\033[1;32m[GATEWAY SUKSES] Sistem otomatis berhasil mengirimkan pesan WhatsApp ke nomor %s!\033[0m\n", toPhone)
	}
	fmt.Println(border)
	fmt.Println()
}

// Disconnect dipanggil saat server mati
func Disconnect() {
	if waClient != nil {
		waClient.Disconnect()
	}
}
