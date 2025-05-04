package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {

	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "What image formats are supported?",
			Answer:   "We support JPG, PNG, and GIF formats. Please make sure your image is in one of these formats before uploading.",
		},
		{
			Question: "What is the maximum file size allowed?",
			Answer:   "The maximum image size allowed is 5MB. Larger files will not be accepted.",
		},
		{
			Question: "How long does it take to upload an image?",
			Answer:   "Uploads usually take just a few seconds depending on your internet connection.",
		},
		{
			Question: "Is my image secure?",
			Answer:   "Yes, all images are uploaded over a secure connection and stored privately.",
		},
		{
			Question: "Can I preview my image before uploading?",
			Answer:   "Yes. You can see a preview before confirming the upload.",
		},
		{
			Question: "How do I delete an uploaded image?",
			Answer:   `You can delete images through your dashboard under "My Uploads".`,
		},
		{
			Question: "Still have questions?",
			Answer:   "Please contact our support at support@example.com.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}

}
