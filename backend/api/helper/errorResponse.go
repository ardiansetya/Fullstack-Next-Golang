package helper

import (
	"jamal/api/api/models/web"
)

func HandleErrorResponse(response *web.WebResponse, err error) {
	if err != nil {
		// RESPONE KTIKA INTERNAL SERVER RROR
		if response.Code != 404 {
			*response = web.WebResponse{
				Code:   500,
				Status: "Internal Server Error",
				Data:   nil,
			}
		}
	}
}
