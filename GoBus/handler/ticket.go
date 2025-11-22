package handler

import (
	"GoBus/dto"
	"GoBus/model"
	"GoBus/service"
)

func ProcessTicket(request dto.Request) (dto.Response, error) {
	destinations, err := service.LoadDestinations()
	if err != nil {
		return dto.Response{}, err
	}

	price, exists := destinations[request.Destination]
	if !exists {
		return dto.Response{
			Name:        request.Name,
			Destination: request.Destination,
			Price:       0,
			Message:     "Destination not found",
		}, nil
	}

	ticket := model.Ticket{
		Name:        request.Name,
		Destination: request.Destination,
		Price:       price,
	}

	return dto.Response{
		Name:        ticket.Name,
		Destination: ticket.Destination,
		Price:       ticket.Price,
		Message:     "Ticket processed successfully",
	}, nil
}
