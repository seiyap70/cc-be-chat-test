## A test chat system

### design

   The system was design combined by two part. One for manage online chat server named `server_manager`, 
the other was named `chat_server` which response for client connection.
    A client ask `server_manager` for access url for `chat_server` before
connect, the `server_manager` user the  requested chat room hash to a 
`chat_server` index which response for next client connection. `server_manager`
was designed as restful api server. server may need more efficacious way to
manage `chat_server` online, offline. 
    `chat_server` user websocket wait for client connection. server init register
all command handler which handle request parse, deal client request.