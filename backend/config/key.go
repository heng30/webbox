package config

import (
    "os"
    "log"
)

const key string = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDMWtilTXnpCTPq
ZK1+6HMXKtzEphYIZg/aBYPnxz+Ln1PJ7o1NrVqSGY1Tk0Y0+V5P1DY8265SnV41
ZXhT9oyjZ/GYIhVFUDGAwvhosygqDGT9LdSgNEKIP9pS2DesQLU7QA0aGkCpo+sy
YBAs20ygaqYvxxRSVao6G8fMdcTd7Ul0ZGKhnlm5QmyjVLdAGsrE1AvO7caBN0nn
XQuUm9UA4CCbxvJsrXLeaUZfqcCFtZFImfpoP0Q+6pqhMkDZn5OPANB4LUP5pcnl
x/wtIrI6ygGMvc69Tr1DvWxflSb3wSTuBrKB58IRG9NKLTpGzJ8L7I7kiJZYm51C
3AVIgC21AgMBAAECggEAZikIY2MfDgP+wy//mBEm8C8KXPjWzO4RAKJ/NdDjEK7f
GXU7iNigDafegih/EZ6Uqn3he30rYbuwTCtlAzyfRFWDbtWjIpqblM5jlxvxeUvm
9LcDs+9MT4b0YzQaaO+UaoFzy1fSebwrqWxiybn8TbnjSMFShHK0/gXDVZCOkY7X
/mo+cylnYrg9F6/4YfGHzzwKVELAwAB38lWsKjOvC5CbF0Q8+iS4UIPL8MXoApLS
9i7/TWKKxSQdD6qbztcGu05/gPK6RSnynFtWCTKhuF0rSteX3Tp8s1gB9GId29Kd
zGCbFI6ne359L9WeEwjIApVYJLXwXXF8gKhnZsclYQKBgQDR0c4qCharb7a0WbVt
d5bKbtUIK5nxvDhOU5sPd1bNeXlQXpxguFhv2vbxJNX09Hz1tkoWheqEz5qvD13t
FS91iZkOaHqny14xULQu/8g+r0GLDOyedK8ckFqPlwJtUNAJCrsddTO0IMCWL2wH
8oShFoSHGmlkSVmM62TtYZyR3QKBgQD5VSK4WbWO/BNx9wXvpyrfWS1wUviHLnSq
Rvm/SyUfydobblvVJDP8zrwNMt/KeDf2edkERNToIYkdZjdIrxudk7jbm82ezJlY
OfRZKcD5tshfrudBu6lrOaKwrhGMH0jcEtChVNtnZY3fuHW/kH6iARXHH3xdjDE9
yvPKp4IJuQKBgQCVHkgoho0AAGBYIt2XVAPeERDkQqChNgTKXg742fbeB3QZk3QO
JVXtATC5x1UsR0o8EvbqtPgstaMYwZIeeg0FuuoFXGm3sQhbgiDcujqlmka5vVC5
ePIvGcTTAh7edC22D2NWl0JMxW46Dq/3oftyxR8hKbs3ZA354h4Tu2jcHQKBgQDH
Hhh2RR5brDATSYVyGHyClgNVus1Dl/QPsFLvVCG3u0n1sjKgiyMBeFBx8fkLBusa
acDISfgF0hJgaisoXr+dHkX5C/owIZEtBkt/kXrNcg2fbFy/ABGS0hp9IuZckvdq
2ZPH4668ajvdOMDAVbLryhYMSbbWNol0ocoEQ7Xz8QKBgDukqGEXs5YDD6yqsqyW
PM2P3zd9b6kLW7a5ExTKJpEI2dxPQW7D4A0tzWtFghOEGv7LIOjiLj0Q/ROaW+GZ
c1iU0xbj+mekePhSpyM8BZGZEWVj6gNk/aminh32xvMr3pb4TjSmbHnjauxpRwz3
OFyD+hyZJqDafOVuWe3YZ8q4
-----END PRIVATE KEY-----`

func saveDefaultKey() {
	if _, err := os.Stat(KeyFile); err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(KeyFile, []byte(key), 0666); err != nil {
				log.Fatalln(err)
			}
		}
	}
}
