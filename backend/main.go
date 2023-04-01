package main

import (
    "flag"
    "fmt"
    "local/config"
    "local/db"
    "local/svr"
    "log"
    "os"
    "path/filepath"
    "strings"
    "crypto/md5"
)

var (
    queryUser = flag.Bool("query", false, "query all users")
    addUser   = flag.String("add", "", "add user")
    delUser   = flag.String("del", "", "delete user")
)

func main() {
    parseArgs()
    initEnv()
    db.Init(config.DBPath + "/" + config.DB_BAME)
    defer db.Uninit()

    if flag.NFlag() > 0 {
        runAsCln()
    } else {
        runAsSvr()
    }
}

func parseArgs() {
    flag.Usage = func() {
        fmt.Printf("Usage: %s [-qurey=true] | [-add <username,password>] | [del <username>]\n", filepath.Base(os.Args[0]))
        flag.PrintDefaults()
        fmt.Println()
        fmt.Println("Version:", VERSION)
        fmt.Println("Configure: ~/.config/webbox/config.json")
        fmt.Println("Database: ~/.local/share/webbox/webbox.db")
        fmt.Println("cert.pem: ~/.config/webbox/cert.pem")
        fmt.Println("key.pem: ~/.config/webbox/key.pem")
    }
    flag.Parse()
}

func runAsCln() {
    if *queryUser {
        if tokens, err := db.QueryUsers(); err != nil {
            log.Fatalln(err)
        } else {
            fmt.Println(tokens)
        }
    } else if *addUser != "" {
        up := strings.Split(*addUser, ",")
        if len(up) != 2 {
            log.Fatalf("input format invalid: %s", *addUser)
        }

        up[0] = strings.TrimSpace(up[0])
        up[1] = strings.TrimSpace(up[1])

        if (up[0] == "" || up[1] == "") {
            log.Fatalf("input format invalid: %s", *addUser)
        }

        up[1] = fmt.Sprintf("%x", md5.Sum([]byte(up[1])))

        if err := db.AddUser(up[0], up[1]); err != nil {
            log.Fatalln(err)
        }
    } else if *delUser != "" {
        if err := db.DelUser(*delUser); err != nil {
            log.Fatalln(err)
        }
    }
}

func runAsSvr() {
    svr.Start()
}

func initEnv() {
    log.SetFlags(log.Lshortfile | log.LstdFlags)
    config.Init()
}
