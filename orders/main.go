package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Azure/go-amqp"
)

type Order struct {
	ID int
	Quantidade int
	Equipe string
}

type Producer struct {
	sender *amqp.Sender
}

func NewProducer (session *amqp.Session)  *Producer{
	sender, err := session.NewSender(context.TODO(), "gdg::new_orders", nil)
	if err != nil {
		log.Println(err.Error())
	}

	return &Producer{sender}
}

func main() {
	http.HandleFunc("/criar-pedido", createOrder)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Println(err)
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	conn, err := amqp.Dial(context.TODO(), "amqp://artemis:artemis@localhost:61616", nil)
	if err != nil {
		log.Println(err.Error())
	}

	session, err := conn.NewSession(context.TODO(), nil)
	if err != nil {
		log.Println(err.Error())
	}

	producer := NewProducer(session)

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	quantidade, _ := strconv.Atoi(r.URL.Query().Get("quantidade"))
	
	// send a message
	order := &Order{
		ID: id,
		Quantidade: quantidade,
		Equipe: "isis-joao-vitor-adam-alyson-marcelo",
	}
	orderJson, err := json.Marshal(order)
	if err != nil {
		log.Panicln("Erro ao gerar o JSON da Order")
	}
	
	err = producer.sender.Send(context.TODO(), amqp.NewMessage(orderJson), nil)
	if err != nil {
		log.Panicln(err.Error())
	}

	io.WriteString(w, "Deu certo!\n")
}