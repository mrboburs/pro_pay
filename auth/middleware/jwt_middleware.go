package middleware

// var conf = config.Config()

// // JWTProtected handler_func for specify routes group with JWT authentication.
// // See: https://github.com/gofiber/jwt
// handler_func JWTProtected(ctx *gin.Context) handler_func(*gin.Context) {
// 	// Create config for JWT authentication middleware.
// 	config := ctx.Config{
// 		"SigningKey":   []byte(conf.JWTSecretKey),
// 		"ContextKey":   "jwt", // used in private routes
// 		"ErrorHandler": jwtError,
// 	}
// 	gin.New(config)
// 	return ctx
// }

// handler_func jwtError(ctx *gin.Context, err error) {
// 	// Return status 401 and failed authentication error.
// 	if err.Error() == "Missing or malformed JWT" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 		return
// 	}
// 	// Return status 401 and failed authentication error.
// 	ctx.JSON(http.StatusUnauthorized, gin.H{
// 		"error": true,
// 		"msg":   err.Error(),
// 	})
// 	return
// }
