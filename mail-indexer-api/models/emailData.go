package models

type EmailData struct {
	From            string `json:"from"`
	To              string `json:"to"`
	Subject         string `json:"subject"`
	Content         string `json:"content"`
	MessageID       string `json:"message_id"`
	Date            string `json:"date"`
	ContentType     string `json:"content_type"`
	MimeVersion     string `json:"mime_version"`
	ContentEncoding string `json:"content_transfer_encoding"`
	XFrom           string `json:"x_from"`
	XTo             string `json:"x_to"`
	Xcc             string `json:"x_cc"`
	Xbcc            string `json:"x_bcc"`
	XFolder         string `json:"x_folder"`
	XOrigin         string `json:"x_origin"`
	XFilename       string `json:"x_filename"`
}

func NewEmail() EmailData {
	return EmailData{
		Subject:     "(no subject)",
	}
}
