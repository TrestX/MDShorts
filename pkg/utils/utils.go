package utils

import (
	"net/url"
	"strings"

	"github.com/aekam27/trestCommon"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	trestCommon.LoadConfig()
}

func EmailLoginOTP(email, name, verificationCode, typ string) (string, error) {
	subject := viper.GetString("emaillogin.loginsubject")
	htmlBody := viper.GetString("emaillogin.initial") + viper.GetString("emaillogin.logintop") + name + viper.GetString("emaillogin.mid") + viper.GetString("emaillogin.verifcode") + verificationCode + "</h1>" + viper.GetString("emaillogin.end")
	if typ == "Signup" {
		subject = viper.GetString("emaillogin.signupsubject")
		htmlBody = viper.GetString("emaillogin.initial") + viper.GetString("emaillogin.greettop") + name + viper.GetString("emaillogin.mid") + viper.GetString("emaillogin.verifcode") + verificationCode + "</h1>" + viper.GetString("emaillogin.greetmsg") + viper.GetString("emaillogin.end")
	}

	textBody := viper.GetString("emaillogin.initial") + "\n\n" + "Hi " + email + ",\n\n" + viper.GetString("emaillogin.mid") + verificationCode + ""
	return sendEmail(email, subject, htmlBody, textBody)
}

func SendVerificationCode(email, name, verificationCode string) (string, error) {
	url := createUrl(verificationCode, "verifyemail")
	subject := viper.GetString("email.subject")
	htmlBody := viper.GetString("email.initial") + name + viper.GetString("email.mid") + " href=" + url + ">Verify Email Now</a>" + viper.GetString("email.end")
	textBody := viper.GetString("email.initial") + name + viper.GetString("email.mid") + " href=" + url + ">Verify Email Now</a>" + viper.GetString("email.end")
	return sendEmail(email, subject, htmlBody, textBody)
}

func sendEmail(email, subject, htmlBody, textBody string) (string, error) {
	svc, err := createSeSSession()
	if err != nil {
		trestCommon.ECLog3("send email verification failed", err, logrus.Fields{"email": email, "htmlBody": htmlBody})
		return "", err
	}
	from := viper.GetString("email.from")
	to := email
	input := &ses.SendEmailInput{
		Source: &from,
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(htmlBody),
				},
				Text: &ses.Content{
					Data: aws.String(textBody),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Destination: &ses.Destination{
			ToAddresses: []*string{&to},
		},
	}
	_, err = svc.SendEmail(input)
	if err != nil {
		trestCommon.ECLog3("send email verification failed", err, logrus.Fields{"email": email, "htmlBody": htmlBody})
		return "", err
	}
	return "Sent Successfully", nil
}

func createUrl(verificationcode, path string) string {
	cart := viper.GetString("website.url")
	website := cart
	if strings.Contains(cart, "https") {
		cartSplit := strings.Split(cart, "/")
		website = cartSplit[2]
	}
	u := &url.URL{
		Scheme: "https",
		Host:   website,
		Path:   path + "/" + verificationcode,
	}
	return u.String()
}

func createSeSSession() (*ses.SES, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(viper.GetString("aws.region")),
		Credentials: credentials.NewStaticCredentials(viper.GetString("aws.aws_access_key_id"),
			viper.GetString("aws.aws_secret_access_key"), "")},
	)
	if err != nil {
		trestCommon.ECLog2("creating ses session", err)
		return nil, err
	}
	svc := ses.New(sess)
	return svc, nil
}
