package main

func tasks(d *Dispatcher) {
	// 接口列表
	d.HandleFunc("send_welcome_email", sendWelcomeEmail)
	d.HandleFunc("send_reminder_email", sendReminderEmail)
}
