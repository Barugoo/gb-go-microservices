package jwt

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		payload, err := Parse(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		req := r.WithContext(context.WithValue(r.Context(), "user_id", payload.UserID))
		next.ServeHTTP(w, req)
	})
}

type Payload struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
	IsPaid bool   `json:"is_paid"`
}

const (
	Secret  = "secret"
	JWTHead = `
	{
		"alg":"HS256",
		"typ":"JWT"
	}`
)

func buildSign(head, payload string) (string, error) {
	// Подготовим хэш функцию для подписи
	// Она будет основана на алгоритме sha256
	// Ключ для простоты будем хранить в константе
	hasher := hmac.New(sha256.New, []byte(Secret))

	// конкатенируем через точку head и payload
	// эти данные мы будем подписывать
	_, err := hasher.Write([]byte(head + "." + payload))
	if err != nil {
		return "", err
	}

	// Вычисляем подпись и кодируем ее в base64
	signBin := hasher.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(signBin), nil
}

func Make(data Payload) (string, error) {

	// Сначала кодируем данные в виде JSON строки
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// По спецификации JWT все части токена должны быть представлены
	// в виде base64url. В стандартную библиотеку go входит такой енкодер
	enc := base64.RawURLEncoding

	// Кодируем переданные данные и заголовок в base64.
	dataB64 := enc.EncodeToString(dataJSON)
	// Head везде будет одинаков, поэтому будем хранить его в константе
	headB64 := enc.EncodeToString([]byte(JWTHead))

	// Формируем подпись токена
	sign, err := buildSign(headB64, dataB64)
	if err != nil {
		return "", fmt.Errorf("Build sign error: %w", err)
	}

	// конкатенируем итоговый результат
	return headB64 + "." + dataB64 + "." + sign, nil
}

func Parse(tok string) (Payload, error) {

	// Распаковываем токен - делим переданную строку на три части
	// Если их не три, то нам передали не правильный токен
	parts := strings.Split(tok, ".")
	if len(parts) != 3 {
		return Payload{}, fmt.Errorf("Token '%s' invalid", tok)
	}

	// Проверяем подпись
	sign, err := buildSign(parts[0], parts[1])
	if err != nil {
		return Payload{}, fmt.Errorf("Build sign error: %w", err)
	}
	if sign != parts[2] {
		return Payload{}, fmt.Errorf("Wrong token signature")
	}

	// Декодируем из base64 payload
	dataJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return Payload{}, fmt.Errorf("Invalid payload: %w", err)
	}

	// Переводим JSON строку в структуру
	payload := Payload{}
	err = json.Unmarshal(dataJSON, &payload)
	if err != nil {
		return Payload{}, fmt.Errorf("Invalid payload JSON: %w", err)
	}

	return payload, nil
}
