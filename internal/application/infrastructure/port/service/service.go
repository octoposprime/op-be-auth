package application

// ServicePort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the other services.
type ServicePort interface {
    LoggingServicePort 
    UserServicePort
}
