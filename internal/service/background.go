package service

import (
	"os"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
		if err != nil {
			return nil, err
		}

		password := os.Getenv("REDIS_PASSWORD")
		if password != "" {
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
		}

		return c, err
	},
}

type Context struct {
	customerID int64
}

var (
	MyEnqueue *work.Enqueuer
)

func RunBackground() {
	// MyEnqueue = work.NewEnqueuer("getcare", redisPool)
	// pool := work.NewWorkerPool(Context{}, 10, "getcare", redisPool)
	// pool.Job("log_write", (*Context).LogWrite)
	// pool.Start()
}

// func (c *Context) LogWrite(job *work.Job) error {
// 	id := job.ArgString("log_id")
// 	clientIP := job.ArgString("client_ip")
// 	method := job.ArgString("method")
// 	fullPath := job.ArgString("full_path")
// 	statusCode := job.ArgString("status_code")
// 	proto := job.ArgString("proto")
// 	userAgent := job.ArgString("user_agent")
// 	param := TrimImportText(job.ArgString("param"))
// 	userID := job.ArgInt64("user_id")
// 	userName := job.ArgString("user_name")
// 	userCode := job.ArgString("user_code")
// 	body := TrimImportText(job.ArgString("body"))
// 	response := TrimImportText(job.ArgString("response"))

// 	today := time.Now()
// 	todayStr := today.Format("2006-01-02 15:04:05")

// 	table := fmt.Sprintf("log_%s", today.Format("20060102"))
// 	columns := "`id`, `user_id`, `user_name`, `user_code`, `client_ip`, `method`, `full_path`, `status_code`, `proto`, `user_agent`, `param`, `body`, `response`, `created_at`, `updated_at`"
// 	values := fmt.Sprintf("'%s', %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'", id, userID, userName, userCode, clientIP, method, fullPath, statusCode, proto, userAgent, param, body, response, todayStr, todayStr)
// 	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s) ON DUPLICATE KEY UPDATE response=VALUES(response), status_code=VALUES(status_code), updated_at=VALUES(updated_at), user_id=VALUES(user_id), user_name=VALUES(user_name), user_code=VALUES(user_code)", table, columns, values)
// 	fmt.Println(sql)
// 	// repository.DbLog.Exec(sql)

// 	return nil
// }
