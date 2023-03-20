package handler

import (
	"honoka-chan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TosCheckHandler(ctx *gin.Context) {
	ctx.Header("user_id", "3241988")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1679236701&version=1.1&token=cHPoOHP5dAs2dh30EkOW8FndO07xlpKHrDRdVOtT7Whlo1opiEMXSwk1JJdAFd4cSeKQvGVRwH2Z7sFh1gnz3gd&nonce=6&requestTimeStamp=1679236698")
	ctx.Header("X-Message-Sign", "3OYeXseR08OvVJfG9cEU6CbEXwbjAhL93vTEL6G4i3FqCY5wpELp0XR8FVZeHo7wsO9UI3+5JJZylnlWvaPgaXej2oefsk5cWHO2rKvrPxaqWRfz5YeGZBvXQejY81KgRRZBWZaQBlHEacH+aILl608xwQGQ98wGtyyMYfOf4Ss=")
	ctx.String(http.StatusOK, utils.ReadAllText("assets/toscheck.json"))
}
