package consumer

// func main() {
// 	conn, err := amqp.Dial(context.TODO(), "amqp://artemis:artemis@localhost:61616", nil)
// 	// conn, err := amqp.Dial(context.TODO(), "broker=\"0.0.0.0\",component=addresses,address=\"gdg\",subcomponent=queues,routing-type=\"anycast\",queue=\"new_orders\"", nil)
// 	if err != nil {
// 		log.Panicln(err.Error())
// 	}

// 	session, err := conn.NewSession(context.TODO(), nil)
// 	if err != nil {
// 		log.Panicln(err.Error())
// 	}

// 	// create a new receiver
// 	receiver, err := session.NewReceiver(context.TODO(), "gdg::new_orders", nil)
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	// receive the next message
// 	msg, err := receiver.Receive(context.TODO(), nil)
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	log.Println(msg)
// }

