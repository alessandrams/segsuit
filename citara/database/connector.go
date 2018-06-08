package main

import (
    "log"
    "github.com/gocql/gocql"
    "fmt"
)

type Device struct{
    Serial uint64
    FirmwareV int
    DevType int
}

type Token struct {
    DevToken uint64
    DevSerial uint64
    Status int
}

type User struct {
    Email string
    Name string
    Password string
}

func databaseConn() (session *gocql.Session){
    cluster := gocql.NewCluster("172.16.1.228")
    cluster.Consistency = gocql.Quorum
    session, err := cluster.CreateSession()

    if err != nil {
        log.Printf("[ERROR] Failed to establish connection to host.\n")
        panic(err)
    }

    log.Printf("[INFO] Connected to database\n")

    return session
}

func getDeviceInfo(session *gocql.Session){
    var device Device

    qry := session.Query("SELECT device_serial, firmware_version, device_type FROM identity.devices")
    itr := qry.Iter()
    for itr.Scan(&device.Serial, &device.FirmwareV, &device.DevType){
      fmt.Println("Query answer:", device)
    }
}

func getTokenInfo(session *gocql.Session){
    var token Token

    qry := session.Query("SELECT device_token, device_serial, token_status FROM identity.tokens")
    itr := qry.Iter()
    for itr.Scan(&token.DevToken, &token.DevSerial, &token.Status){
      fmt.Println("Query answer:", token)
    }
}

func getUserInfo(session *gocql.Session){
    var user User
    fmt.Println("Query answer:")
    qry := session.Query("SELECT user_email, name, password FROM users.users")
    itr := qry.Iter()
    for itr.Scan(&user.Email, &user.Name, &user.Password){
      fmt.Println("Query answer:", user)
    }
}

func main() {
    session := databaseConn()
    getUserInfo(session)
    defer session.Close()
}
