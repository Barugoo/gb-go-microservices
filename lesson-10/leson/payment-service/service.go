package main

import (
	"context"
	"fmt"

	pb "payment-service/api"
)

type Service struct {
	db Repository
	user
}

func (s *Service) convertAmount(amount int64) string {
	return fmt.Sprintf("%d.%d", amount/100, amount%100)
}

func (s *Service) UpdateTransactionStatus(ctx context.Context, req *pb.UpdateTransactionStatusRequest) (*pb.UpdateTransactionStatusResponse, error) {
	transaction, err := s.db.GetTransaction(ctx, req.TransactionID)
	if err != nil {
		return nil, err
	}
	switch req.NewStatus {
	case pb.TransactionStatusNew:
		transaction.Status = TransactionStatusNew
	case pb.TransactionStatusDone:
		transaction.Status = TransactionStatusDone
	case pb.TransactionStatusFailure:
		transaction.Status = TransactionStatusFailure
	}
	transaction, err := s.db.UpdateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTransactionResponse{true}, nil
}

func (s *Service) BuyMovie(ctx context.Context, req *pb.BuyMovieRequest) (*pb.BuyMovieResponse, error) {
	resp, er := s.movieService.GetMovie(ctx, req.MovieID) // достать фильм из муви сервиса чтобы получить его стоимость
	if err != nil {
		return nil, err
	}
	moviePrice := resp.Price

	balanceAmount, err := s.db.GetUserBalance(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if moviePrice > balanceAmount { // если баланс меньше чем стоимость фильма - вернуть ошибку что нет деняш
		return status.Err // ERR NOT ENOUGH FUNDS
	}

	// должно быть выплнено одной транзакцией в проде!!
	//отсюда
	movieOwnership := &MovieOwnership{
		UserID:  req.UserId,
		MovieID: req.MovieId,
	}
	movieOwnership, err := s.db.CreateMovieOwnership(ctx, movieOwnership)
	if err != nil {
		return nil, err
	}

	transaction := &Transaction{
		UserID: req.UserId,
		Amount: -moviePrice,
		Status: TransactionStatusDone,
	}
	transaction, err := s.db.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	// до сюда
	return &pb.BuyMovieResponse{true}, nil
}

func (s *Service) GenerateDepositLink(ctx context.Context, req *pb.GenerateDepositLinkRequest) (*pb.GenerateDepositLinkResponse, error) {
	transaction := &Transaction{
		UserID: req.UserId,
		Amount: req.Amount,
		Status: TransactionStatusNew,
	}
	transaction, err := s.db.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}
	payload := fmt.Sprintf(`
	{
		"Amount": "%d",
		"OrderId": "d",
		"Description": "Пополнение счета аккаунта на %s рублей",
		"DATA": {
			"Phone": "%s",
			"Email": "%s"
		},
		"Receipt": {
			"Email": "a@test.ru",
			"Phone": "+79031234567",
			"EmailCompany": "b@test.ru",
			"Taxation": "osn"
		}
		"success_url": "http://localhost:8080/accept-payment?order_id=%s",
		"failure_url": "http://localhost:8080/decline-payment?order_id=%s",
	}
	`, req.Amount, orderID, convertAmount(req.Amount), user.Phone, user.Email, orderID, orderID,
	)

	return &pb.GenerateDepositLinkResponse{Payload: payload}, nil
}
