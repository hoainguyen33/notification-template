package domain

import (
	"fmt"
	"getcare-notification/internal/model"
	"getcare-notification/internal/repository"
	"getcare-notification/utils"
	"strings"
	"time"
)

type RequestOtpDomain interface {
	RequestOtpAdd(requestOtpAdd model.RequestOtpAdd) error
	RequestOTP(requestOtpAdd model.RequestOtpAdd, code string) error
	CheckRequestOTP(phone string) error
	VerifyOTP(param model.VerifyOTPParam) error
}

type requestOtpDomain struct {
	RequestOtpRepository repository.RequestOtpRepository
	VerifyOtpRepository  repository.VerifyOtpRepository
	//GetcarePhoneDomain  GetcarePhoneDomain
}

func NewRequestOtpDomain(requestOtpRepository repository.RequestOtpRepository, verifyOtpRepository repository.VerifyOtpRepository,

//getcarePhoneDomain GetcarePhoneDomain
) RequestOtpDomain {
	return &requestOtpDomain{
		RequestOtpRepository: requestOtpRepository,
		VerifyOtpRepository:  verifyOtpRepository,
		//GetcarePhoneDomain:  getcarePhoneDomain,
	}
}

// add request otp to database
func (ro *requestOtpDomain) RequestOtpAdd(requestOtpAdd model.RequestOtpAdd) error {
	code := utils.GeneratorNumber(6)
	// create request otp with condition
	if err := ro.RequestOTP(requestOtpAdd, code); err != nil {
		return err
	}

	// message := fmt.Sprintf("Phahub: Ma OTP de xac thuc SDT cua quy khach la %s. Xin vui long xac nhan trong 150 giây", code)
	// messageCode := fmt.Sprintf("%s-%s-%s-%s", strings.ToUpper(GeneratorString(4)), strings.ToUpper(GeneratorString(4)), strings.ToUpper(GeneratorString(4)), strings.ToUpper(GeneratorString(4)))
	// phoneSMS := &socket.GetcarePhoneSMS{
	// 	Phone:       requestOtpAdd.Phone,
	// 	Message:     message,
	// 	MessageCode: messageCode,
	// }
	// push message sms with notification and socket messager
	// err := ro.GetcarePhoneDomain.NewMessageWithoutUser(phoneSMS)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (ro *requestOtpDomain) RequestOTP(requestOtpAdd model.RequestOtpAdd, code string) error {
	phone := FormatPhone(requestOtpAdd.Phone)

	if err := ro.CheckRequestOTP(phone); err != nil {
		return err
	}

	today := time.Now()
	requestOTP := &model.RequestOtp{
		Phone:     phone,
		Code:      code,
		ExpiredAt: today.Local().Add(time.Second * time.Duration(150)),
	}

	if _, err := ro.RequestOtpRepository.Create(requestOTP); err != nil {
		return err
	}
	return nil
}

func (ro *requestOtpDomain) CheckRequestOTP(phone string) error {
	today := time.Now()
	where := fmt.Sprintf("phone=%s", phone)
	records, err := ro.RequestOtpRepository.ListSort("created_at DESC", where)
	if err != nil {
		return nil
	}
	if len(records) == 0 {
		return nil
	}
	// if > 3 request otp in 1h then return ErrOTP1h
	if len(records) >= 3 {
		check3Interval := records[0].CreatedAt.Add(time.Hour * time.Duration(1))
		if today.Before(check3Interval) {
			return utils.ErrOTP1h
		}
	}

	//check60s := records[0].CreatedAt.Time.Add(time.Second * time.Duration(60))
	//if today.Before(check60s) {
	//	return ErrOTP60s
	//}

	return nil
}

// can chinh sua lai
func (ro *requestOtpDomain) VerifyOTP(param model.VerifyOTPParam) error {
	code := param.Code
	phone := FormatPhone(param.Phone)
	today := time.Now()
	where := fmt.Sprintf("phone=%s", phone)
	records, err := ro.RequestOtpRepository.ListSort("created_at DESC", where)
	if err != nil {
		return nil
	}

	if len(records) == 0 {
		return utils.ErrOTPIncorrect
	}

	// verifyOtp := &model.VerifyOtp{
	// 	RequestOtpID: records[0].ID,
	// 	Code:         code,
	// }
	// hoi lai
	// if _, _, err := ro.VerifyOtpRepository.Create(verifyOtp); err != nil {
	// 	return err
	// }

	if code != records[0].Code {
		return utils.ErrOTPIncorrect
	}

	expiredAt := records[0].ExpiredAt
	if today.After(expiredAt) {
		return utils.ErrOTPExpired
	}
	where = fmt.Sprintf("request_otp_id=%d", records[0].ID)
	verifyOtpItems, err := ro.VerifyOtpRepository.ListSort("", where)
	if err == nil {
		if len(verifyOtpItems) > 3 {
			return utils.ErrOTP3Times
		}
	}

	for _, record := range records {
		where = fmt.Sprintf("request_otp_id=%d", record.ID)
		if err := ro.VerifyOtpRepository.DeleteWhere(where); err != nil {
			return err
		}
	}
	where = fmt.Sprintf("phone='%s'", phone)
	if err := ro.RequestOtpRepository.DeleteWhere(where); err != nil {
		return err
	}

	return nil
}

func FormatPhone(phone string) string {
	if strings.HasPrefix(phone, "+") {
		return phone
	}

	//todo: xử lý thêm mã vùng của các quốc gia khác

	return phone
}
