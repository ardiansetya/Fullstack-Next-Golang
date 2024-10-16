package exception

/*
	SEBENARNYA MAU DI PAKE TAPI BINGGUNG CARA AMBIL SERROR NYA JADI GA UDLU DEH
*/

//
//func ErrorHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			// Tangkap panic atau error
//			if err := recover(); err != nil {
//				// Cek apakah error termasuk NotFoundError
//				if notFoundError(c, err) {
//					return
//				}
//
//				// Jika bukan NotFoundError, kirimkan internal server error
//				internalServerError(c, err)
//			}
//		}()
//
//		// Lanjutkan eksekusi handler berikutnya
//		c.Next()
//	}
//}
//
//func notFoundError(c *gin.Context, err interface{}) bool {
//	// Cek apakah error merupakan NotFoundError
//	exception, ok := err.(NotFoundError)
//	if ok {
//		webResponse := web.WebResponse{
//			Code:   http.StatusNotFound,
//			Status: "NOT FOUND",
//			Data:   exception.Error,
//		}
//		c.JSON(http.StatusNotFound, webResponse)
//		return true
//	}
//	return false
//}
//
//func internalServerError(c *gin.Context, err interface{}) {
//	// Menangani error selain NotFoundError
//	webResponse := web.WebResponse{
//		Code:   http.StatusInternalServerError,
//		Status: "INTERNAL SERVER ERROR",
//		Data:   nil,
//	}
//	c.JSON(http.StatusInternalServerError, webResponse)
//}
//
////
////func ErrorHandler(c *gin.Context, err interface{}) {
////	if notFoundError(c, err) {
////		return
////	}
////
////	internalServerError(c, err)
////}
////
////func notFoundError(c *gin.Context, err interface{}) bool {
////	exception, ok := err.(NotFoundError)
////	if ok {
////		webResponse := web.WebResponse{
////			Code:   http.StatusNotFound,
////			Status: "NOT FOUND",
////			Data:   exception.Error,
////		}
////		c.JSON(http.StatusNotFound, webResponse)
////		return true
////	} else {
////		return false
////	}
////}
////
////func internalServerError(c *gin.Context, err interface{}) {
////	webResponse := web.WebResponse{
////		Code:   http.StatusInternalServerError,
////		Status: "INTERNAL SERVER ERROR",
////		Data:   nil,
////	}
////	c.JSON(http.StatusInternalServerError, webResponse)
////}
//
////func internalServerError() gin.HandlerFunc {
////	return func(c *gin.Context) {
////		webResponse := web.WebResponse{
////			Code:   http.StatusInternalServerError,
////			Status: "INTERNAL SERVER ERROR",
////			Data:   nil,
////		}
////		c.JSON(http.StatusInternalServerError, webResponse)
////	}
////}
