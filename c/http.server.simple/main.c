/************************
###
### C listen 8080 Web
### 2020-10-18
*/

#include <stdlib.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <stdio.h>

#define PORT 8080

int main() {

  printf("\n\033[0;33mRun Server port:%d\033[0m\n", PORT);
  
  int server_fd = socket(AF_INET, SOCK_STREAM, 0);
  if(server_fd < 0){
        printf("Error in connection.\n");
        exit(1);
    }

  struct sockaddr_in server;
 
  server.sin_family = AF_INET;

  server.sin_port = htons(PORT);

  server.sin_addr.s_addr = htonl(INADDR_ANY);

  bind(server_fd, (struct sockaddr*) &server, sizeof(server)); 

  listen(server_fd, 128);

  while (1) {

    int client_fd = accept(server_fd, NULL, NULL);

    char response[] =
      "HTTP/1.1 200 OK\r\nContent-Length: 21\r\nConnection: close\r\n\r\n<h1>Hello, world!</h1>";

    for (int sent = 0; sent < sizeof(response); sent += send(client_fd, response+sent, sizeof(response)-sent, 0));
    close(client_fd);

  }
  return 0;
}
