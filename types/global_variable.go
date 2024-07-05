package types

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// .env 전역변수
var (
	LotusRpcUrl *string
	AuthToken   *string
)

// 에러를 로깅할 log파일 인스턴스
var LogFile *os.File

// 전역 HTTP 클라이언트 인스턴스
// HTTP connection 인스턴스 재사용
var Client = &http.Client{}

// 전역 변수 초기화
func Init() error {
	var err error

	// HTTP 클라이언트 초기화
	Client = &http.Client{}

	// 로깅 파일을 엽니다
	LogFile, err = os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// dotenv 실행
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// .env에서 RPC URL, LOTUS TOKEN 가져오기
	lotusRpcUrl := os.Getenv("LOTUS_RPC_URL")
	authToken := "Bearer " + os.Getenv("AUTH_TOKEN")
	LotusRpcUrl = &lotusRpcUrl
	AuthToken = &authToken

	if LotusRpcUrl == nil || AuthToken == nil {
		log.Fatalf("Environment variables not set properly")
	}
	return nil
}

// 전역 인스턴스 리소스 정리
func CleanUp() {
	if LogFile != nil {
		LogFile.Close()
	}

	if Client != nil {
		Client.CloseIdleConnections()
	}
}
