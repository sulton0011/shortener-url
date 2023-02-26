package errors

import (
	"errors"
	"shortener-url/config"
	"shortener-url/pkg/logger"
)

type ErrorService struct {
	log      logger.LoggerI
	svcsPort string
	svcsName string
}

func NewErrorService(log logger.LoggerI, svcsName string, svcsPort string) *ErrorService {
	return &ErrorService{
		log:      log,
		svcsName: svcsName,
		svcsPort: svcsPort,
	}
}

type ErrorStorage struct {
	log      logger.LoggerI
	strgName string
}

func NewErrorStorage(log logger.LoggerI, strgName string) *ErrorStorage {
	return &ErrorStorage{
		log:      log,
		strgName: strgName,
	}
}

func (e *ErrorService) Wrap(err *error, funcName string, req interface{}) {
	// fmt.Print(err)
	if *err == nil {
		return
	}
	*err = Wrap(*err, funcName)

	e.log.Error(msges(config.ErrorModel, e.svcsName),
		logger.Error(*err),
		logger.Any("Service Port", e.svcsPort),
		logger.Any("request:", req),
	)

	*err = Wrap(Wrap(*err, e.svcsName), config.ErrorModel)

}

func (e *ErrorService) Info(funcName string, req interface{}) {
	e.log.Error(msges(config.InfoModel, e.svcsName),
		// logger.Any("Service Port", e.svcsPort),
		logger.Any("request:", req),
	)
}

func (e *ErrorStorage) Wrap(err *error, funcName string, req interface{}) {
	if *err == nil {
		return
	}
	*err = Wrap(*err, funcName)

	e.log.Error(msges(config.ErrorModel, e.strgName),
		logger.Error(*err),

		logger.Any("request:", req),
	)

	*err = Wrap(*err, e.strgName)

}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return errors.New(msg + config.ErrorStyle + err.Error())
}

func msges(msg1, msg2 string) string {
	return msg1 + config.ErrorStyle + msg2
}

func New(msg string) error {
	return errors.New(msg)
}

func WrapCheck(err *error, msg string) {
	if *err == nil {
		return
	}
	er := *err

	*err = errors.New(msg + config.ErrorStyle + er.Error())
}
