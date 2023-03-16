
struct Root {
  1: Wifi wifi
  2: MQTT mqtt
}

struct Wifi {
  1: string ssid
  2: string passphrase
  3: i32 timeoutSeconds = 10
}

struct MQTT {
  1: string host
  2: i32 port = 1883
  3: string topicPrefix
  4: bool auth
  5: string username
  6: string password
  7: bool ssl
  8: string fingerprint
}
