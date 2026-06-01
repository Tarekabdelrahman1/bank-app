package db

import (
	"context"
	"log"
	"os"
	"testing"

	// بنستخدم الدرايفر الحديث عشان يطابق الـ DBTX المتولدة حالياً
	"github.com/jackc/pgx/v5/pgxpool"
)

// اسم المتغير testQueries عشان يطابق ملف التيست النظيف بتاعك
var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbSource := "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

	// فتح الاتصال بالأسلوب الحديث المتوافق
	connPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer connPool.Close()

	// هنا دالة New هتستقبل الـ connPool وهي بتضحك لإن الأنواع متطابقة!
	testQueries = New(connPool)

	os.Exit(m.Run())
}
