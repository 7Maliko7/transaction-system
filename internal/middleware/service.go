package middleware

import "github.com/7Maliko7/transaction-system/internal/service"

type Middleware func(service service.Servicer) service.Servicer
