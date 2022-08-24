#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>

int main(void)
{
    int sock0;
    struct sockaddr_in addr;
    struct sockaddr_in client;
    socklen_t len;
    int sock1, sock2;

    sock0 = socket(AF_INET, SOCK_STREAM, 0);

    addr.sin_family = AF_INET;

    addr.sin_port = htons(11111);

    addr.sin_addr.s_addr = INADDR_ANY;

    bind(sock0, (struct sockaddr *)&addr, sizeof(addr));

    listen(sock0, 5);

    len = sizeof(client);
    sock1 = accept(sock0, (struct sockaddr *)&client, &len);

    write(sock1, "HELLO", 6);

    close(sock1);

    len = sizeof(client);
    sock2 = accept(sock0, (struct sockaddr *)&client, &len);

    write(sock2, "HOGE", 5);

    close(sock2);

    close(sock0);

    return 0;
}
